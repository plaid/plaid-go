"""Release policy. All cryptographic verification is delegated to GitHub CLI.

There is deliberately no production option to accept preverified JSON or skip
verification. Tests replace subprocess boundaries, not the production policy.
"""

import hashlib
import json
import os
from pathlib import Path, PurePosixPath
import re
import subprocess
import tempfile


SDK_REPO = "plaid/plaid-go"
CONFIG_REPO = "plaid/plaid-openapi"
SPEC_REPO = "plaid/plaid-openapi"
GEN_WORKFLOW = ".github/workflows/generate-go-sdk.yml"
RELEASE_WORKFLOW = ".github/workflows/release.yml"
GEN_REF = "refs/heads/master"
RELEASE_REF = "refs/heads/master"
GEN_PREDICATE = "https://plaid.com/attestations/sdk-generation/v1"
RELEASE_PREDICATE = "https://plaid.com/attestations/sdk-release/v1"
MANIFEST = ".sdk-release/manifest.json"
GEN_BUNDLE = ".sdk-release/generation.bundle.json"
PUBLIC_REMOTE = "https://github.com/plaid/plaid-go.git"


class ReleaseError(Exception):
    pass


def require(condition, message):
    if not condition:
        raise ReleaseError(message)


def sha256(data):
    return hashlib.sha256(data).hexdigest()


def canonical(value):
    return json.dumps(value, sort_keys=True, separators=(",", ":"), ensure_ascii=True).encode()


def parse_json(data):
    def pairs(values):
        result = {}
        for key, value in values:
            require(key not in result, "Duplicate JSON key")
            result[key] = value
        return result

    def constant(_):
        raise ReleaseError("Non-JSON numeric constant")

    try:
        return json.loads(data, object_pairs_hook=pairs, parse_constant=constant)
    except (ValueError, UnicodeError) as error:
        raise ReleaseError("Invalid JSON") from error


def run(*args, cwd=None, env=None, input=None):
    result = subprocess.run([str(arg) for arg in args], cwd=cwd, env=env, input=input,
                            stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    require(result.returncode == 0, f"{args[0]} failed; operation stopped before publication")
    return result.stdout


def git(repo, *args):
    return run("git", "-C", repo, *args)


def hex_value(value, length, label):
    require(isinstance(value, str) and re.fullmatch(rf"[0-9a-f]{{{length}}}", value), f"Invalid {label}")


def decimal(value, label):
    require(isinstance(value, str) and re.fullmatch(r"[1-9][0-9]*", value), f"Invalid {label}")


def version(value):
    numeric = r"(?:0|[1-9][0-9]*)"
    require(isinstance(value, str) and re.fullmatch(rf"{numeric}\.{numeric}\.{numeric}(?:-[0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*)?", value), "Invalid SDK version")
    require(int(value.split(".")[0]) >= 49, "Automated releases start at v49")
    for identifier in value.partition("-")[2].split("."):
        require(not (identifier.isdigit() and len(identifier) > 1 and identifier.startswith("0")), "Invalid prerelease version")


def keys(value, expected, label):
    require(isinstance(value, dict) and set(value) == set(expected.split()), f"Invalid {label} fields")


def validate_manifest(value):
    keys(value, "schema_version sdk source config generator legacy content_sha256 generation", "manifest")
    require(type(value["schema_version"]) is int and value["schema_version"] == 1, "Unsupported manifest schema")
    keys(value["sdk"], "repository version module", "sdk")
    sdk = value["sdk"]
    require(sdk["repository"] == SDK_REPO, "Untrusted SDK repository")
    version(sdk["version"])
    require(sdk["module"] == f"github.com/{SDK_REPO}/v{sdk['version'].split('.')[0]}", "Incorrect module path")
    keys(value["source"], "repository commit path sha256", "source")
    require(value["source"]["repository"] == SPEC_REPO and value["source"]["path"] == "2020-09-14.yml", "Untrusted public spec location")
    keys(value["config"], "repository commit sha256", "config")
    require(value["config"]["repository"] == CONFIG_REPO, "Untrusted public config repository")
    for name in ("source", "config"):
        hex_value(value[name]["commit"], 40, f"{name} commit")
        hex_value(value[name]["sha256"], 64, f"{name} digest")
    require(value["config"]["commit"] == value["source"]["commit"], "Spec and config must come from the same public commit")
    keys(value["generator"], "cli_version cli_sha256 image_digest", "generator")
    require(isinstance(value["generator"]["cli_version"], str) and re.fullmatch(r"[0-9]+\.[0-9]+\.[0-9]+", value["generator"]["cli_version"]), "Invalid Fern CLI version")
    hex_value(value["generator"]["cli_sha256"], 64, "Fern CLI digest")
    keys(value["legacy"], "source_commit spec_sha256 tree_sha256 generator_digest", "legacy")
    hex_value(value["legacy"]["source_commit"], 40, "legacy generator source commit")
    for name in ("spec_sha256", "tree_sha256"):
        hex_value(value["legacy"][name], 64, f"legacy {name}")
    require(value["legacy"]["spec_sha256"] == value["source"]["sha256"], "Both SDKs must use the same public spec")
    for digest in (value["generator"]["image_digest"], value["legacy"]["generator_digest"]):
        require(isinstance(digest, str) and digest.startswith("sha256:"), "Generator must be digest-pinned")
        hex_value(digest[7:], 64, "generator image digest")
    hex_value(value["content_sha256"], 64, "content digest")
    keys(value["generation"], "repository workflow run_id", "generation")
    require(value["generation"]["repository"] == CONFIG_REPO and value["generation"]["workflow"] == GEN_WORKFLOW, "Untrusted generation workflow")
    decimal(value["generation"]["run_id"], "generation run ID")
    return value


def tree_files(repo, commit):
    hex_value(commit, 40, "SDK commit")
    require(git(repo, "cat-file", "-t", commit).strip() == b"commit", "SDK SHA must identify a commit")
    entries = []
    for record in git(repo, "ls-tree", "-rz", commit).split(b"\0"):
        if not record:
            continue
        header, raw_path = record.split(b"\t", 1)
        mode, kind, oid = header.decode().split()
        path = raw_path.decode("utf-8")
        require(mode in ("100644", "100755") and kind == "blob", "Symlinks/submodules are forbidden")
        require(not path.startswith("/") and "\\" not in path and all(part not in ("", ".", "..") and part.lower() != ".git" for part in path.split("/")), "Unsafe tracked path")
        require(not any(ord(char) < 32 or ord(char) == 127 for char in path), "Control character in tracked path")
        require(not path.startswith(".sdk-release/") or path in (MANIFEST, GEN_BUNDLE), "Unexpected .sdk-release file")
        entries.append((path, mode, oid))
    objects = list(dict.fromkeys(oid for _, _, oid in entries))
    data = run("git", "-C", repo, "cat-file", "--batch", input=("\n".join(objects) + "\n").encode())
    blobs = {}
    offset = 0
    for oid in objects:
        end = data.index(b"\n", offset)
        actual_oid, kind, raw_size = data[offset:end].split()
        size = int(raw_size)
        require(actual_oid.decode() == oid and kind == b"blob" and size <= 100 * 1024 * 1024, "Invalid Git blob")
        offset = end + 1
        blobs[oid] = data[offset:offset + size]
        require(len(blobs[oid]) == size and data[offset + size:offset + size + 1] == b"\n", "Truncated Git blob")
        offset += size + 1
    return {path: (mode, blobs[oid]) for path, mode, oid in entries}


def content_bytes(files, legacy_only=False):
    entries = [[path, mode, sha256(data)] for path, (mode, data) in sorted(files.items())
               if path not in (MANIFEST, GEN_BUNDLE) and (not legacy_only or path.startswith("legacy/"))]
    return canonical(entries)


def validate_tree(repo, commit):
    files = tree_files(repo, commit)
    require(MANIFEST in files and GEN_BUNDLE in files, "SDK has no generation manifest and signed bundle")
    manifest = validate_manifest(parse_json(files[MANIFEST][1]))
    require(sha256(content_bytes(files)) == manifest["content_sha256"], "SDK content was changed after generation")
    require(sha256(content_bytes(files, legacy_only=True)) == manifest["legacy"]["tree_sha256"], "Legacy tree digest mismatch")
    require("go.mod" in files and "go.sum" in files, "Missing module files")
    module_lines = re.findall(rb"^module\s+(\S+)\s*$", files["go.mod"][1], re.MULTILINE)
    require(module_lines == [manifest["sdk"]["module"].encode()], "Manifest module differs from go.mod")
    require(".fern/metadata.json" in files, "Missing Fern metadata")
    metadata = parse_json(files[".fern/metadata.json"][1])
    keys(metadata, "cliVersion generatorName generatorVersion sdkVersion", "public Fern metadata")
    require(metadata.get("sdkVersion") == "v" + manifest["sdk"]["version"], "Fern SDK version mismatch")
    require(metadata.get("cliVersion") == manifest["generator"]["cli_version"], "Fern CLI version mismatch")
    require(metadata["generatorName"] == "fernapi/fern-go-sdk" and isinstance(metadata["generatorVersion"], str)
            and re.fullmatch(r"[0-9]+\.[0-9]+\.[0-9]+", metadata["generatorVersion"]), "Unexpected public generator metadata")
    require(not any(path.endswith("/go.mod") and path.startswith("legacy/") for path in files), "Legacy must share the root module")
    if manifest["sdk"]["version"].split(".")[0] == "49":
        require("legacy/client.go" in files and "legacy/configuration.go" in files, "v49 must bundle legacy")
    return manifest, files


def verified_attestation(artifact, bundle, *, repository, workflow, ref, commit, predicate_type, predicate, subjects, run_id, private):
    identity = f"https://github.com/{repository}/{workflow}@{ref}"
    arguments = ["gh", "attestation", "verify", str(artifact), "--bundle", str(bundle),
                 "--hostname", "github.com", "--repo", repository, "--cert-identity", identity,
                 "--cert-oidc-issuer", "https://token.actions.githubusercontent.com", "--source-ref", ref,
                 "--source-digest", commit, "--signer-digest", commit, "--deny-self-hosted-runners",
                 "--predicate-type", predicate_type, "--format", "json"]
    if private:
        arguments.append("--no-public-good")
    results = parse_json(run(*arguments))
    require(isinstance(results, list) and results, "No verified attestation returned")
    for result in results:
        verification = result.get("verificationResult", {})
        certificate = verification.get("signature", {}).get("certificate", {})
        statement = verification.get("statement", {})
        expected_certificate = {
            "subjectAlternativeName": identity,
            "issuer": "https://token.actions.githubusercontent.com",
            "buildSignerURI": identity,
            "buildSignerDigest": commit,
            "sourceRepositoryURI": f"https://github.com/{repository}",
            "sourceRepositoryDigest": commit,
            "sourceRepositoryRef": ref,
            "runnerEnvironment": "github-hosted",
            "sourceRepositoryVisibilityAtSigning": "private" if private else "public",
        }
        require(all(certificate.get(key) == value for key, value in expected_certificate.items()), "Verified certificate identity does not match release policy")
        invocation = rf"https://github\.com/{re.escape(repository)}/actions/runs/{re.escape(run_id)}(?:/attempts/[1-9][0-9]*)?"
        require(re.fullmatch(invocation, certificate.get("runInvocationURI", "")), "Attestation run ID mismatch")
        require(verification.get("verifiedTimestamps"), "Attestation has no verified timestamp")
        require(statement.get("_type") == "https://in-toto.io/Statement/v1" and statement.get("predicateType") == predicate_type, "Wrong attestation statement type")
        require(statement.get("predicate") == predicate, "Attestation predicate differs from manifest")
        actual_subjects = statement.get("subject")
        require(isinstance(actual_subjects, list) and len(actual_subjects) == len(subjects), "Unexpected attestation subjects")
        require(sorted(actual_subjects, key=lambda item: item.get("name", "")) == sorted(subjects, key=lambda item: item["name"]), "Attested subjects do not match artifacts")
    return results


def verify_generation(repo, commit):
    manifest, files = validate_tree(repo, commit)
    with tempfile.TemporaryDirectory(prefix="sdk-generation-verification-") as directory:
        content = Path(directory) / "canonical-content.json"
        bundle = Path(directory) / "generation.bundle.json"
        content.write_bytes(content_bytes(files))
        bundle.write_bytes(files[GEN_BUNDLE][1])
        verified_attestation(content, bundle, repository=CONFIG_REPO, workflow=GEN_WORKFLOW,
                             ref=GEN_REF, commit=manifest["config"]["commit"], predicate_type=GEN_PREDICATE,
                             predicate=manifest,
                             subjects=[{"name": f"plaid-go-v{manifest['sdk']['version']}-source", "digest": {"sha256": manifest["content_sha256"]}}],
                             run_id=manifest["generation"]["run_id"], private=False)
    return manifest, files


def api(path, paginate=False, raw=False):
    arguments = ["gh", "api", "--hostname", "github.com", path]
    if paginate:
        arguments += ["--paginate", "--slurp"]
    if raw:
        arguments += ["-H", "Accept: application/vnd.github.raw+json"]
    result = run(*arguments)
    if raw:
        return result
    value = parse_json(result)
    return [item for page in value for item in page] if paginate else value


def merged_release_pull(commit):
    hex_value(commit, 40, "merged SDK commit")
    branch = api(f"repos/{SDK_REPO}/branches/master")
    require(branch.get("protected") is True and branch.get("commit", {}).get("sha") == commit,
            "Release commit must be protected public master tip")
    pulls = api(f"repos/{SDK_REPO}/commits/{commit}/pulls?per_page=100", paginate=True)
    eligible = [pull for pull in pulls if pull.get("merged_at") and pull.get("merge_commit_sha") == commit
                and pull.get("base", {}).get("ref") == "master"
                and pull.get("base", {}).get("repo", {}).get("full_name") == SDK_REPO]
    require(len(eligible) == 1, "Release requires one merged SDK PR for this exact commit")
    return eligible[0]


def human_reviews(pull):
    reviews = api(f"repos/{SDK_REPO}/pulls/{pull['number']}/reviews?per_page=100", paginate=True)
    latest = {}
    for review in reviews:
        user = review.get("user", {})
        if (user.get("type") == "User" and review.get("author_association") in ("MEMBER", "OWNER", "COLLABORATOR")
                and review.get("state") in ("APPROVED", "CHANGES_REQUESTED", "DISMISSED")):
            latest[user["login"]] = review
    require(not any(review["state"] == "CHANGES_REQUESTED" for review in latest.values()),
            "Unresolved changes-requested review")
    return latest.values()


def require_generated_pull(pull):
    creator = os.environ.get("SDK_PR_APP_BOT_ID", "")
    decimal(creator, "SDK PR App bot ID")
    require(pull.get("user", {}).get("type") == "Bot"
            and pull.get("user", {}).get("id") == int(creator), "PR was not created by the configured generation App")
    require(pull.get("base", {}).get("repo", {}).get("full_name") == SDK_REPO
            and pull.get("base", {}).get("ref") == "master"
            and pull.get("head", {}).get("repo", {}).get("full_name") == SDK_REPO,
            "Generated PR must stay within the public SDK repository")
    require(re.fullmatch(r"sdk-generation/go-v[1-9][0-9]*-[0-9a-f]{20}", pull.get("head", {}).get("ref", "")),
            "Unexpected generated PR branch")


def require_approved_merge(commit):
    """Retain the explicitly reviewed manual fallback."""
    pull = merged_release_pull(commit)
    reviews = human_reviews(pull)
    require(any(review["state"] == "APPROVED" and review.get("commit_id") == pull.get("head", {}).get("sha")
                for review in reviews), "No human approval of the final PR head")
    return pull["number"]


def require_release_merge(commit):
    """Authorize either the configured automated integration or human review.

    This checks the merge identity only. The caller must independently verify
    the signed generation/release evidence and exact tree before publication.
    """
    if os.environ.get("SDK_AUTOMERGE_ENABLED") != "true":
        return require_approved_merge(commit)
    summary = merged_release_pull(commit)
    pull = api(f"repos/{SDK_REPO}/pulls/{summary['number']}")
    require(pull.get("merged_at") and pull.get("merge_commit_sha") == commit,
            "PR no longer identifies this exact release commit")
    human = human_reviews(pull)
    if any(review["state"] == "APPROVED" and review.get("commit_id") == pull.get("head", {}).get("sha")
           for review in human):
        return pull["number"]
    require_generated_pull(pull)
    merger = os.environ.get("SDK_INTEGRATION_APP_BOT_ID", "")
    decimal(merger, "SDK integration App bot ID")
    require(merger != os.environ.get("SDK_PR_APP_BOT_ID"), "Generation and integration must use distinct Apps")
    require(pull.get("merged_by", {}).get("type") == "Bot"
            and pull.get("merged_by", {}).get("id") == int(merger),
            "PR was not merged by the configured integration App")
    return pull["number"]


def verify_public_spec(manifest):
    source = manifest["source"]
    comparison = api(f"repos/{SPEC_REPO}/compare/{source['commit']}...master")
    require(comparison.get("status") in ("ahead", "identical"), "Spec commit is not on public spec master")
    content = api(f"repos/{SPEC_REPO}/contents/{source['path']}?ref={source['commit']}", raw=True)
    require(sha256(content) == source["sha256"], "Public spec bytes differ from generation manifest")


def require_context(environment, commit, control_commit):
    hex_value(commit, 40, "workflow SDK commit")
    hex_value(control_commit, 40, "trusted release control commit")
    dry_run = (environment.get("GITHUB_EVENT_NAME") == "workflow_dispatch"
               and environment.get("SDK_RELEASE_DRY_RUN") == "true")
    require(environment.get("SDK_RELEASE_ENABLED") == "true" or dry_run, "SDK release pipeline is disabled")
    require(environment.get("GITHUB_REPOSITORY") == SDK_REPO, "Workflow must run in the public SDK repository")
    require(environment.get("GITHUB_REF") == RELEASE_REF and environment.get("GITHUB_REF_PROTECTED") == "true", "Workflow must run on protected master")
    require(environment.get("GITHUB_SHA") == commit, "Workflow source SHA differs from release commit")
    require(environment.get("GITHUB_WORKFLOW_REF") == f"{SDK_REPO}/{RELEASE_WORKFLOW}@{RELEASE_REF}", "Untrusted release workflow context")
    require(environment.get("GITHUB_EVENT_NAME") in ("push", "workflow_dispatch"), "Unsupported release event")
    decimal(environment.get("GITHUB_RUN_ID"), "release run ID")
    decimal(environment.get("GITHUB_RUN_ATTEMPT"), "release attempt")


def require_control_workflow(sdk_repo, commit, control_repo, control_commit):
    require(git(control_repo, "rev-parse", "HEAD").decode().strip() == control_commit, "Wrong trusted control checkout")
    expected = git(control_repo, "show", f"{control_commit}:{RELEASE_WORKFLOW}")
    require(git(sdk_repo, "show", f"{commit}:{RELEASE_WORKFLOW}") == expected, "Release workflow changed since its reviewed control commit")
