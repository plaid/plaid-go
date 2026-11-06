#!/usr/bin/env python3
"""Check generated SDK PRs; merge only with --merge and protected workflow context.

PR trees are inspected as Git data, never checked out or executed. The separate
write token is used only for the final head-pinned merge request.
"""

import argparse
import os
from pathlib import Path
import re
import sys
import tempfile

import policy
from policy import (SDK_REPO, PUBLIC_REMOTE, ReleaseError, api, canonical, decimal,
                    git, hex_value, parse_json, require, run, sha256, tree_files,
                    verify_generation, verify_public_spec)


WORKFLOW = ".github/workflows/integrate.yml"
CI_CONTEXT = "ci/circleci: build"
VERIFIED_CONTEXT = "sdk-generation/verified"
ACTIONS_APP_ID = 15368
ACTIONS_BOT_ID = 41898282
PRESERVED_PREFIXES = (".github/", "scripts/release/", ".circleci/", "legacy/tests/")
# Keep in sync with the public generation pipeline's preserved files. CHANGELOG
# has the one permitted deterministic prepend checked separately below.
PRESERVED_FILES = {
    ".fernignore", "README.md", "LICENSE", "internal/caller.go",
    "internal/destination.go", "internal/destination_test.go",
    "credential_destination_test.go", "core/custom_endpoint.go",
    "option/custom_endpoint.go", "retry_default_test.go", "wire_test_security_test.go",
    "transfer_diligence_upload_test.go",
}


# Reviewed generated surface. New endpoint packages need no policy changes, but
# new scripts, configuration formats or executable files require human review.
GENERATED_GO_PATH = re.compile(r"(?:[a-z][a-z0-9_]*/)*[a-z][a-z0-9_]*\.go")
GENERATED_RESOURCES = {
    "go.mod", "go.sum", ".fern/metadata.json", "reference.md",
    "wiremock/wiremock-mappings.json",
}


def approved_generated_changes(base_files, files):
    for name in base_files.keys() | files.keys():
        before, after = base_files.get(name), files.get(name)
        if (before == after or name in (policy.MANIFEST, policy.GEN_BUNDLE, "CHANGELOG.md")
                or name in PRESERVED_FILES or name.startswith(PRESERVED_PREFIXES)):
            continue  # Identity, maintained content and changelog have stricter checks.
        allowed = name in GENERATED_RESOURCES or (
            not name.startswith("scripts/") and GENERATED_GO_PATH.fullmatch(name) is not None)
        require(allowed and all(value[0] == "100644" for value in (before, after) if value is not None),
                f"Changed path is not approved generated content: {name}")


class NotReady(ReleaseError):
    pass


def integration_identity():
    decimal(os.environ.get("SDK_PR_APP_BOT_ID"), "SDK PR App bot ID")
    decimal(os.environ.get("SDK_INTEGRATION_APP_BOT_ID"), "SDK integration App bot ID")
    require(os.environ["SDK_INTEGRATION_APP_BOT_ID"] != os.environ["SDK_PR_APP_BOT_ID"],
            "Generation and integration must use distinct Apps")


def context(control, control_commit, *, status_only=False):
    dry_run = (os.environ.get("GITHUB_EVENT_NAME") == "workflow_dispatch"
               and os.environ.get("SDK_INTEGRATION_DRY_RUN") == "true")
    require(os.environ.get("SDK_AUTOMERGE_ENABLED") == "true" or dry_run or status_only, "SDK integration is disabled")
    require(os.environ.get("GITHUB_REPOSITORY") == SDK_REPO, "Untrusted integration repository")
    require(os.environ.get("GITHUB_REF") == "refs/heads/master"
            and os.environ.get("GITHUB_REF_PROTECTED") == "true", "Integration requires protected master")
    require(os.environ.get("GITHUB_WORKFLOW_REF") == f"{SDK_REPO}/{WORKFLOW}@refs/heads/master", "Untrusted integration workflow")
    require(os.environ.get("GITHUB_EVENT_NAME") in ("status", "schedule", "workflow_dispatch", "workflow_run"), "Untrusted integration event")
    hex_value(control_commit, 40, "control commit")
    hex_value(os.environ.get("GITHUB_SHA"), 40, "workflow commit")
    integration_identity()
    decimal(os.environ.get("SDK_CIRCLECI_STATUS_ACTOR_ID"), "CircleCI status actor ID")
    decimal(os.environ.get("SDK_REVIEW_RULESET_ID"), "SDK review ruleset ID")
    require(git(control, "rev-parse", "HEAD").decode().strip() == control_commit, "Wrong trusted control checkout")


def fetch_objects(repo, *commits):
    for commit in commits:
        hex_value(commit, 40, "public SDK commit")
    environment = {key: os.environ[key] for key in ("HOME", "PATH", "TMPDIR") if key in os.environ}
    environment.update(GIT_CONFIG_GLOBAL="/dev/null", GIT_CONFIG_NOSYSTEM="1", GIT_CONFIG_COUNT="0", GIT_TERMINAL_PROMPT="0")
    run("git", "-C", repo, "-c", "credential.helper=", "-c", "core.hooksPath=/dev/null",
        "fetch", "--no-tags", PUBLIC_REMOTE, *dict.fromkeys(commits), env=environment)


def current_base():
    branch = api(f"repos/{SDK_REPO}/branches/master")
    require(branch.get("protected") is True, "Public master is not protected")
    commit = branch.get("commit", {}).get("sha")
    hex_value(commit, 40, "public master commit")
    return commit


def strict_status_protection():
    # GitHub's merge API pins only the head. Strict status checks must remain
    # enforced for the App to close the concurrent-base-update race server-side.
    protection = api(f"repos/{SDK_REPO}/branches/master/protection")
    checks = protection.get("required_status_checks") or {}
    contexts = set(checks.get("contexts", []))
    contexts.update(check.get("context") for check in checks.get("checks", []))
    require(checks.get("strict") is True and CI_CONTEXT in contexts,
            "Integration requires strict CircleCI branch protection")
    require([item.get("app_id") for item in checks.get("checks", []) if item.get("context") == VERIFIED_CONTEXT] == [ACTIONS_APP_ID],
            "Integration requires the generated-output status bound to GitHub Actions")
    require(not protection.get("required_pull_request_reviews"),
            "Move approval requirements to the PR-only review ruleset; classic approval bypass is not supported")


def approval_ruleset(token=None):
    ruleset_id = os.environ.get("SDK_REVIEW_RULESET_ID", "")
    decimal(ruleset_id, "SDK review ruleset ID")
    # Effective rules include inherited rules and omit disabled/evaluate rulesets.
    effective = api(f"repos/{SDK_REPO}/rules/branches/master?per_page=100", paginate=True, token=token)
    require(any(rule.get("type") == "pull_request" and rule.get("ruleset_id") == int(ruleset_id)
                for rule in effective), "Configured review ruleset is not active on master")
    ruleset = api(f"repos/{SDK_REPO}/rulesets/{ruleset_id}", token=token)
    require(ruleset.get("id") == int(ruleset_id) and ruleset.get("source_type") == "Repository"
            and ruleset.get("source") == SDK_REPO and ruleset.get("target") == "branch"
            and ruleset.get("enforcement") == "active", "Unexpected review ruleset identity or enforcement")
    refs = (ruleset.get("conditions") or {}).get("ref_name") or {}
    require("refs/heads/master" in refs.get("include", []) and refs.get("exclude") == [],
            "Review ruleset must explicitly include master without exclusions")
    rules = ruleset.get("rules") or []
    require(len(rules) == 1 and rules[0].get("type") == "pull_request",
            "Review bypass ruleset must contain only the pull-request rule")
    reviews = rules[0].get("parameters") or {}
    count = reviews.get("required_approving_review_count")
    require(type(count) is int and count >= 1 and reviews.get("require_code_owner_review") is True
            and reviews.get("dismiss_stale_reviews_on_push") is True,
            "Review ruleset requires human approval, code owners and stale-review dismissal")
    if token is not None:
        # This is the capability of the exact installation token about to merge.
        # The full bypass actor list may be hidden from this caller.
        require(ruleset.get("current_user_can_bypass") == "pull_requests_only",
                "Merger token must have PR-only approval bypass")


def require_identity(pull):
    policy.require_generated_pull(pull)
    require(pull.get("state") == "open" and pull.get("draft") is False, "Generated PR must be open and ready")
    require(pull.get("merged") is not True, "Generated PR is already merged")
    require(pull.get("commits") == 1, "Generated PR must contain exactly one commit")
    for side in ("base", "head"):
        hex_value(pull[side].get("sha"), 40, f"PR {side} commit")


def discover():
    candidates = []
    for pull in api(f"repos/{SDK_REPO}/pulls?state=open&base=master&per_page=100", paginate=True):
        try:
            policy.require_generated_pull(pull)
        except ReleaseError:
            continue
        if pull.get("draft") is False:
            candidates.append(pull["number"])
    require(len(candidates) <= 1, "Multiple generated PRs require reconciliation before integration")
    return candidates[0] if candidates else None


def ci_and_reviews(number, head, require_verified=False):
    combined = api(f"repos/{SDK_REPO}/commits/{head}/status")
    require(combined.get("sha") == head, "CI status is for another head")
    statuses = api(f"repos/{SDK_REPO}/commits/{head}/statuses?per_page=100", paginate=True)
    # GitHub returns newest status first. A later pending/failing status must
    # override an older success, including a rerun of the same job.
    matching = [status for status in statuses if status.get("context") == CI_CONTEXT]
    if not matching or matching[0].get("state") == "pending":
        raise NotReady("Waiting for CircleCI on the exact generated head")
    status = matching[0]
    require(status.get("state") == "success", "CircleCI failed on the generated head")
    require(str(status.get("creator", {}).get("id")) == os.environ["SDK_CIRCLECI_STATUS_ACTOR_ID"], "CircleCI status has an untrusted author")
    require(re.fullmatch(r"https://(?:circleci\.com/gh/plaid/plaid-go/[1-9][0-9]*|app\.circleci\.com/pipelines/github/plaid/plaid-go/[1-9][0-9]*(?:/workflows/[0-9a-f-]+)?)", status.get("target_url", "")), "Unexpected CircleCI build URL")
    # The verifier posts pending before evaluating this head. Exclude only its own
    # context to avoid a circular dependency, keeping every other latest status.
    latest_statuses = {}
    for item in statuses:
        latest_statuses.setdefault(item.get("context"), item)
    other = [item for name, item in latest_statuses.items() if name != VERIFIED_CONTEXT]
    if any(item.get("state") == "pending" for item in other):
        raise NotReady("Waiting for all reported commit statuses")
    require(other and all(item.get("state") == "success" for item in other), "A reported commit status failed")
    if require_verified:
        verified = latest_statuses.get(VERIFIED_CONTEXT) or {}
        require(verified.get("state") == "success" and verified.get("creator", {}).get("id") == ACTIONS_BOT_ID,
                "Missing trusted generated-output verification on the exact head")
        require(re.fullmatch(r"https://github\.com/plaid/plaid-go/actions/runs/[1-9][0-9]*", verified.get("target_url", "")),
                "Unexpected verification workflow URL")
    reviews = api(f"repos/{SDK_REPO}/pulls/{number}/reviews?per_page=100", paginate=True)
    latest = {}
    for review in reviews:
        user = review.get("user", {})
        if (user.get("type") == "User" and review.get("author_association") in ("MEMBER", "OWNER", "COLLABORATOR")
                and review.get("state") in ("APPROVED", "CHANGES_REQUESTED", "DISMISSED")):
            latest[user["login"]] = review
    require(not any(review["state"] == "CHANGES_REQUESTED" for review in latest.values()), "Human changes-requested review blocks integration")


def preserved_content(base_files, files, manifest):
    approved_generated_changes(base_files, files)
    protected = lambda name: name in PRESERVED_FILES or name.startswith(PRESERVED_PREFIXES)
    before = {name: value for name, value in base_files.items() if protected(name)}
    after = {name: value for name, value in files.items() if protected(name)}
    require("go.mod" in base_files, "Missing base module")
    modules = re.findall(rb"^module\s+github\.com/plaid/plaid-go/v([1-9][0-9]*)\s*$",
                         base_files["go.mod"][1], re.MULTILINE)
    require(len(modules) == 1 and int(modules[0]) >= 49, "Invalid base SDK module")
    old_major = int(modules[0])
    new_major = policy.major_transition(old_major, manifest["sdk"]["version"])
    if new_major != old_major:
        # Only the exact old module token changes in maintained Go code/examples.
        # CI, workflows, release controls and every file mode remain byte-identical.
        module = re.compile(rb"github\.com/plaid/plaid-go/v" + str(old_major).encode() + rb"(?![0-9A-Za-z_.-])")
        replacement = f"github.com/plaid/plaid-go/v{new_major}".encode()
        before = {name: (mode, module.sub(replacement, data)
                         if (name in PRESERVED_FILES or name.startswith("legacy/tests/"))
                         and (name.endswith(".go") or name == "README.md") else data)
                  for name, (mode, data) in before.items()}
    require(before == after, "Generated PR changed protected automation or custom SDK files")
    require("CHANGELOG.md" in base_files and "CHANGELOG.md" in files, "Missing preserved changelog")
    old_mode, old = base_files["CHANGELOG.md"]
    mode, new = files["CHANGELOG.md"]
    require(mode == old_mode, "Changed changelog file mode")
    if new == old:
        return
    source = manifest["source"]
    spec = api(f"repos/{source['repository']}/contents/{source['path']}?ref={source['commit']}", raw=True)
    require(sha256(spec) == source["sha256"], "Changelog spec does not match signed input")
    info = re.search(rb"^info:\s*\n((?:[ \t].*\n|\s*\n)*)", spec, re.MULTILINE)
    require(info is not None, "Missing public spec version")
    versions = re.findall(rb"^  version:\s*['\"]?(\d{4}-\d{2}-\d{2}_\d+\.\d+\.\d+)['\"]?\s*$", info[1], re.MULTILINE)
    require(len(versions) == 1, "Ambiguous public spec version")
    metadata = parse_json(files[".fern/metadata.json"][1])
    prefix = (f"## [v{manifest['sdk']['version']}]\n\nRegenerate the new client and legacy package from OpenAPI `{versions[0].decode()}` "
              f"using Fern Go `{metadata['generatorVersion']}` for the new client.\n\n").encode()
    require(new == prefix + old, "Generated changelog must only prepend the exact release entry")


def check(repo, control, control_commit, number, expected_base=None, expected_head=None, *, status_only=False):
    context(control, control_commit, status_only=status_only)
    pull = api(f"repos/{SDK_REPO}/pulls/{number}")
    require_identity(pull)
    base, head = pull["base"]["sha"], pull["head"]["sha"]
    require(current_base() == base, "PR base moved; regenerate on current master")
    require(expected_base is None or base == expected_base, "PR base changed after verification")
    require(expected_head is None or head == expected_head, "PR head changed after verification")
    strict_status_protection()
    approval_ruleset()
    fetch_objects(repo, base, head, os.environ["GITHUB_SHA"])
    expected_workflow = git(control, "show", f"{control_commit}:{WORKFLOW}")
    require(git(repo, "show", f"{os.environ['GITHUB_SHA']}:{WORKFLOW}") == expected_workflow
            and git(repo, "show", f"{base}:{WORKFLOW}") == expected_workflow, "Integration workflow differs from reviewed controls")
    parents = git(repo, "show", "-s", "--format=%P", head).decode().strip().split()
    require(parents == [base], "Generated commit must be the direct child of current master")
    manifest, files = verify_generation(repo, head)
    verify_public_spec(manifest)
    identity = sha256(canonical([manifest["sdk"]["version"], manifest["source"]["sha256"], manifest["config"]["sha256"]]))[:20]
    require(pull["head"]["ref"] == f"sdk-generation/go-v{manifest['sdk']['version'].split('.')[0]}-{identity}", "Generated branch identity does not match signed inputs")
    preserved_content(tree_files(repo, base), files, manifest)
    ci_and_reviews(number, head)
    return {"number": number, "base": base, "head": head, "version": manifest["sdk"]["version"],
            "tree": git(repo, "rev-parse", f"{head}^{{tree}}").decode().strip()}


def merge(verified):
    require(os.environ.get("SDK_AUTOMERGE_ENABLED") == "true", "SDK integration is disabled")
    require(os.environ.get("SDK_INTEGRATION_DRY_RUN") != "true", "An integration dry run cannot merge")
    integration_identity()
    number, head, base = verified["number"], verified["head"], verified["base"]
    pull = api(f"repos/{SDK_REPO}/pulls/{number}")
    require_identity(pull)
    require(pull["head"]["sha"] == head and pull["base"]["sha"] == base and current_base() == base, "PR changed immediately before merge")
    strict_status_protection()
    ci_and_reviews(number, head, require_verified=True)
    token = os.environ.get("SDK_INTEGRATION_TOKEN", "")
    require(token, "Missing scoped integration App token")
    approval_ruleset(token=token)
    body = canonical({"sha": head, "merge_method": "squash", "commit_title": f"Generate plaid-go v{verified['version']} (#{number})"})
    environment = dict(os.environ, GH_TOKEN=token, GH_HOST="github.com")
    result = parse_json(run("gh", "api", "--hostname", "github.com", "--method", "PUT",
                            f"repos/{SDK_REPO}/pulls/{number}/merge", "--input", "-", input=body, env=environment))
    require(result.get("merged") is True, "GitHub did not merge the exact generated head")
    commit = result.get("sha")
    hex_value(commit, 40, "merged SDK commit")
    merged = api(f"repos/{SDK_REPO}/git/commits/{commit}")
    require([parent.get("sha") for parent in merged.get("parents", [])] == [base]
            and merged.get("tree", {}).get("sha") == verified["tree"], "Merged result differs from the verified base and tree; stop publication")
    return commit


def output(value):
    if os.environ.get("GITHUB_OUTPUT"):
        with open(os.environ["GITHUB_OUTPUT"], "a") as stream:
            for key, item in value.items():
                require("\n" not in str(item), "Invalid workflow output")
                stream.write(f"{key}={item}\n")
    print(canonical(value).decode())


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--control", required=True, type=Path)
    parser.add_argument("--control-commit", required=True)
    parser.add_argument("--number", type=int)
    parser.add_argument("--base")
    parser.add_argument("--head")
    parser.add_argument("--merge", action="store_true")
    args = parser.parse_args()
    context(args.control, args.control_commit)
    require(not args.merge or (args.number and args.base and args.head), "Merge requires the previously verified PR, base and head")
    number = args.number or discover()
    if number is None:
        output({"ready": "false", "reason": "No ready generated PR"})
        return
    require(number > 0, "Invalid PR number")
    with tempfile.TemporaryDirectory(prefix="sdk-integrate-") as temporary:
        repo = Path(temporary) / "objects.git"
        run("git", "init", "--bare", repo)
        verified = check(repo, args.control, args.control_commit, number, args.base, args.head)
        if args.merge:
            output({"merged": "true", "commit": merge(verified)})
        else:
            output({"ready": "true", **verified})


if __name__ == "__main__":
    try:
        main()
    except NotReady as error:
        output({"ready": "false", "reason": str(error)})
    except (ReleaseError, OSError, ValueError, KeyError) as error:
        sys.exit(str(error))
