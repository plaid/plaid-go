"""Tag-only publication after attestation verification; no SDK execution."""

import os
from pathlib import Path
import re
import subprocess
import tempfile

from artifacts import verify_release
from policy import (MANIFEST, PUBLIC_REMOTE, SPEC_REPO, ReleaseError, api, git,
                    hex_value, parse_json, require, require_release_merge, run,
                    sha256, validate_manifest, validate_tree)


def network_git(repo, remote, *arguments, push_token=None):
    settings = subprocess.run(["git", "-C", str(repo), "config", "--null", "--includes", "--get-regexp", r"^url\..*\.(insteadof|pushinsteadof)$"], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    require(settings.returncode in (0, 1), "Cannot inspect Git URL rewrites")
    for record in settings.stdout.split(b"\0"):
        _, separator, prefix = record.partition(b"\n")
        require(not separator or not remote.encode().startswith(prefix), "Git URL rewrite would change the publication destination")
    environment = dict(os.environ, GH_HOST="github.com", GIT_TERMINAL_PROMPT="0")
    if push_token is not None:
        environment["GH_TOKEN"] = push_token
    return run("git", "-C", repo, "-c", "credential.helper=", "-c", "credential.helper=!gh auth git-credential", *arguments, env=environment)


def semver_key(tag):
    match = re.fullmatch(r"v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(?:-([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?", tag)
    if not match:
        return None
    major, minor, patch, prerelease = match.groups()
    if prerelease is None:
        suffix = (1, ())
    else:
        identifiers = []
        for value in prerelease.split("."):
            if value.isdigit():
                if len(value) > 1 and value.startswith("0"):
                    return None
                identifiers.append((0, int(value)))
            else:
                identifiers.append((1, value))
        suffix = (0, tuple(identifiers))
    return int(major), int(minor), int(patch), suffix


def remote_tags(repo, remote):
    tags = {}
    for line in network_git(repo, remote, "ls-remote", "--tags", remote).decode().splitlines():
        oid, ref = line.split("\t")
        hex_value(oid, 40, "remote tag object")
        if ref.startswith("refs/tags/") and not ref.endswith("^{}"):
            name = ref[len("refs/tags/"):]
            if semver_key(name) is not None:
                tags[name] = oid
    return tags


def install_verification_key(repo, public_key, directory):
    require(isinstance(public_key, str) and re.fullmatch(r"ssh-ed25519 [A-Za-z0-9+/]+={0,3}(?: [^\r\n]*)?", public_key.strip()), "Configure an Ed25519 release signing public key")
    allowed = Path(directory) / "allowed-signers"
    allowed.write_text("sdk-release@plaid.com " + public_key.strip() + "\n")
    git(repo, "config", "gpg.format", "ssh")
    git(repo, "config", "gpg.ssh.allowedSignersFile", str(allowed))
    git(repo, "config", "core.hooksPath", "/dev/null")
    git(repo, "config", "user.name", "Plaid SDK Release")
    git(repo, "config", "user.email", "sdk-release@plaid.com")


def verified_existing_tag(repo, remote, tag, commit):
    network_git(repo, remote, "fetch", "--no-tags", remote, f"refs/tags/{tag}:refs/tags/{tag}")
    require(git(repo, "cat-file", "-t", f"refs/tags/{tag}").strip() == b"tag", "Existing release tag is not annotated")
    raw = git(repo, "cat-file", "tag", f"refs/tags/{tag}")
    require(b"-----BEGIN SSH SIGNATURE-----" in raw, "Existing release tag lacks the required SSH signature")
    require(git(repo, "rev-parse", f"refs/tags/{tag}^{{commit}}").decode().strip() == commit, "Existing release version points to another commit")
    git(repo, "verify-tag", f"refs/tags/{tag}")


def publication_state(repo, remote, release, generation):
    tags = remote_tags(repo, remote)
    requested = "v" + release["sdk"]["version"]
    if requested in tags:
        verified_existing_tag(repo, remote, requested, release["commit"])
        return "already-published"
    if not tags:
        return "new"
    latest = max(tags, key=semver_key)
    require(semver_key(requested) > semver_key(latest), "A newer or equal SDK version is already published")
    if semver_key(latest)[0] < 49:
        return "new"
    network_git(repo, remote, "fetch", "--no-tags", remote, f"refs/tags/{latest}:refs/tags/{latest}")
    previous_commit = git(repo, "rev-parse", f"refs/tags/{latest}^{{commit}}").decode().strip()
    if git(repo, "ls-tree", "--name-only", previous_commit, MANIFEST).strip():
        previous_data = git(repo, "show", f"{previous_commit}:{MANIFEST}")
        previous = validate_manifest(parse_json(previous_data))
        require(previous["sdk"]["version"] == latest[1:], "Published tag and its generation manifest disagree")
        old_spec = previous["source"]["commit"]
    else:
        # A manual v49 release predates signed generation manifests. Bridge only
        # the exact predecessor and public spec reviewed by the release owner.
        predecessor = os.environ.get("SDK_RELEASE_MANUAL_PREDECESSOR_COMMIT", "")
        predecessor_tag = os.environ.get("SDK_RELEASE_MANUAL_PREDECESSOR_TAG", "")
        old_spec = os.environ.get("SDK_RELEASE_MANUAL_PREDECESSOR_SPEC_COMMIT", "")
        hex_value(predecessor, 40, "reviewed manual predecessor commit")
        hex_value(old_spec, 40, "reviewed manual predecessor spec commit")
        require(latest == predecessor_tag and previous_commit == predecessor,
                "Missing provenance for an unapproved manual predecessor")
    new_spec = generation["source"]["commit"]
    comparison = api(f"repos/{SPEC_REPO}/compare/{old_spec}...{new_spec}")
    require(comparison.get("status") in ("ahead", "identical"), "Candidate spec is older than or diverged from the last release")
    return "new"


def publish(directory, commit, control_commit, run_id, run_attempt, public_key, *, publish_requested=False, remote=PUBLIC_REMOTE):
    # Called directly by the trusted CLI. Always verify real signed bundles here;
    # there is no CLI option to inject a success result or a different remote.
    if publish_requested:
        require(os.environ.get("SDK_RELEASE_ENABLED") == "true", "SDK release pipeline is disabled")
        require(os.environ.get("SDK_RELEASE_DRY_RUN") != "true", "A workflow dry run cannot publish")
    release, generation = verify_release(directory, commit, control_commit, run_id, run_attempt)
    with tempfile.TemporaryDirectory(prefix="sdk-tag-only-") as temporary:
        repo = Path(temporary) / "repository.git"
        run("git", "init", "--bare", str(repo))
        install_verification_key(repo, public_key, temporary)
        state = publication_state(repo, remote, release, generation)
        if state == "already-published":
            return state
        require_release_merge(commit)
        network_git(repo, remote, "fetch", "--no-tags", remote, commit)
        require(git(repo, "rev-parse", f"{commit}^{{tree}}").decode().strip() == release["git_tree"], "Remote commit tree differs from the attested tree")
        actual_manifest, files = validate_tree(repo, commit)
        require(actual_manifest == generation, "Remote generation manifest differs from attested release")
        require(sha256(files[MANIFEST][1]) == release["generation_manifest_sha256"], "Remote generation manifest bytes changed")
        if not publish_requested:
            return "dry-run"
        require(os.environ.get("SDK_RELEASE_PUBLISH") == "true", "Publication requires explicit workflow authorization")
        token = os.environ.get("SDK_RELEASE_PUSH_TOKEN", "")
        private_key = os.environ.get("SDK_RELEASE_SIGNING_PRIVATE_KEY", "")
        require(token and private_key, "Missing protected release App token or signing key")
        key_path = Path(temporary) / "signing-key"
        key_path.write_text(private_key.rstrip() + "\n")
        key_path.chmod(0o600)
        actual_key = run("ssh-keygen", "-y", "-P", "", "-f", key_path).decode().strip().split()[:2]
        require(actual_key == public_key.strip().split()[:2], "Release private key does not match the pinned public key")
        git(repo, "config", "user.signingkey", str(key_path))
        tag = "v" + release["sdk"]["version"]
        message = f"Release plaid-go {tag}\n\nCommit: {commit}\nRelease-Manifest-SHA256: {sha256((Path(directory) / 'release-manifest.json').read_bytes())}\nModule-ZIP-SHA256: {release['module_zip']['sha256']}"
        git(repo, "tag", "-s", tag, commit, "-m", message)
        git(repo, "verify-tag", tag)
        # Recheck the master tip and version floor immediately before the one
        # permitted write. Workflow concurrency serializes this release path.
        require_release_merge(commit)
        require(publication_state(repo, remote, release, generation) == "new", "Release appeared while preparing the tag; retry verification")
        network_git(repo, remote, "push", remote, f"refs/tags/{tag}:refs/tags/{tag}", push_token=token)
        peeled = network_git(repo, remote, "ls-remote", remote, f"refs/tags/{tag}^{{}}")
        require(peeled.decode().split("\t", 1)[0] == commit, "Remote tag verification failed after push")
        return "published"
