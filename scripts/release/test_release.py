#!/usr/bin/env python3
"""Local release integration. Fake verification outputs never claim real OIDC.

Git, module ZIP creation, SSH signing and signature checks are real. Only GitHub
API responses and the gh cryptographic-verifier process boundary are mocked.
"""

import copy
import json
import os
from pathlib import Path
import subprocess
import tempfile
import unittest
from unittest.mock import patch
import zipfile

import artifacts
import policy
import publish


class ReleaseTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.zip_tool = Path(os.environ["MODULEZIP_BIN"]).resolve()
        if not cls.zip_tool.is_file():
            raise RuntimeError("Build scripts/release/modulezip and set MODULEZIP_BIN first")

    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory(prefix="sdk-release-policy-test-")
        self.addCleanup(self.temporary.cleanup)
        self.root = Path(self.temporary.name).resolve()
        self.repo = self.root / "sdk"
        self.control = self.root / "control"
        self.remote = self.root / "public.git"
        self.out = self.root / "artifacts"
        self.out.mkdir()
        self.environment = {
            "HOME": os.environ["HOME"], "PATH": os.environ["PATH"], "TMPDIR": str(self.root),
            "GIT_CONFIG_GLOBAL": "/dev/null", "GIT_CONFIG_NOSYSTEM": "1",
            "GIT_AUTHOR_NAME": "Release fixture", "GIT_COMMITTER_NAME": "Release fixture",
            "GIT_AUTHOR_EMAIL": "fixture@example.com", "GIT_COMMITTER_EMAIL": "fixture@example.com",
            "SDK_RELEASE_ENABLED": "true", "GITHUB_REPOSITORY": policy.SDK_REPO,
            "GITHUB_REF": policy.RELEASE_REF, "GITHUB_REF_PROTECTED": "true",
            "GITHUB_WORKFLOW_REF": f"{policy.SDK_REPO}/{policy.RELEASE_WORKFLOW}@{policy.RELEASE_REF}",
            "GITHUB_EVENT_NAME": "push", "GITHUB_RUN_ID": "1234", "GITHUB_RUN_ATTEMPT": "1",
        }
        self.environ_patch = patch.dict(os.environ, self.environment, clear=True)
        self.environ_patch.start()
        self.addCleanup(self.environ_patch.stop)
        self.command("git", "init", "-q", "--initial-branch=master", self.repo)
        self.command("git", "init", "-q", "--bare", "--initial-branch=master", self.remote)
        self.write("go.mod", "module github.com/plaid/plaid-go/v49\n\ngo 1.24\n")
        self.write("go.sum", "")
        self.write("sdk.go", "package plaid\n\nconst Version = 49\n")
        self.write("legacy/client.go", "package legacy\n")
        self.write("legacy/configuration.go", "package legacy\n")
        self.write("LICENSE", "Fixture license\n")
        self.write("README.md", "Fixture SDK\n")
        self.write(policy.RELEASE_WORKFLOW, "name: trusted-release\n")
        self.metadata = {"cliVersion": "5.134.0", "generatorName": "fernapi/fern-go-sdk", "generatorVersion": "1.70.3", "sdkVersion": "v49.0.0"}
        self.write(".fern/metadata.json", json.dumps(self.metadata))
        self.control_commit = self.commit()
        self.command("git", "clone", "-q", str(self.repo), str(self.control))
        files = policy.tree_files(self.repo, self.control_commit)
        self.spec = b"openapi: 3.0.3\ninfo:\n  title: fixture\n"
        self.manifest = {
            "schema_version": 1,
            "sdk": {"repository": policy.SDK_REPO, "version": "49.0.0", "module": "github.com/plaid/plaid-go/v49"},
            "source": {"repository": policy.SPEC_REPO, "commit": "a" * 40, "path": "2020-09-14.yml", "sha256": policy.sha256(self.spec)},
            "config": {"repository": policy.CONFIG_REPO, "commit": "a" * 40, "sha256": "c" * 64},
            "generator": {"cli_version": "5.134.0", "cli_sha256": "d" * 64, "image_digest": "sha256:" + "e" * 64},
            "legacy": {"source_commit": "f" * 40, "spec_sha256": policy.sha256(self.spec), "tree_sha256": policy.sha256(policy.content_bytes(files, legacy_only=True)), "generator_digest": "sha256:" + "1" * 64},
            "content_sha256": policy.sha256(policy.content_bytes(files)),
            "generation": {"repository": policy.CONFIG_REPO, "workflow": policy.GEN_WORKFLOW, "run_id": "9876"},
        }
        self.write(policy.MANIFEST, json.dumps(self.manifest))
        self.write(policy.GEN_BUNDLE, '{"fixture_only":"not a real Sigstore attestation"}\n')
        self.sha = self.commit()
        os.environ["GITHUB_SHA"] = self.sha
        self.command("git", "-C", self.repo, "remote", "add", "origin", self.remote)
        self.command("git", "-C", self.repo, "push", "-q", "origin", "master")
        self.key = self.root / "release-key"
        self.command("ssh-keygen", "-q", "-t", "ed25519", "-N", "", "-f", self.key)
        self.public_key = self.key.with_suffix(".pub").read_text().strip()
        self.reviews = [{"user": {"type": "User", "login": "reviewer"}, "author_association": "MEMBER", "state": "APPROVED", "commit_id": "2" * 40}]
        self.gh_calls = []
        self.pull = {"number": 10, "merged_at": "2026-10-12T12:00:00Z", "merge_commit_sha": self.sha,
                     "base": {"ref": "master", "repo": {"full_name": policy.SDK_REPO}}, "head": {"sha": "2" * 40}}

    def command(self, *args):
        result = subprocess.run([str(value) for value in args], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        self.assertEqual(result.returncode, 0, result.stderr.decode())
        return result.stdout

    def write(self, name, value):
        path = self.repo / name
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(value)

    def commit(self):
        self.command("git", "-C", self.repo, "add", "-A")
        self.command("git", "-C", self.repo, "commit", "-q", "-m", "Fixture")
        return self.command("git", "-C", self.repo, "rev-parse", "HEAD").decode().strip()

    def refresh_content(self):
        interim = self.commit()
        files = policy.tree_files(self.repo, interim)
        self.manifest["content_sha256"] = policy.sha256(policy.content_bytes(files))
        self.manifest["legacy"]["tree_sha256"] = policy.sha256(policy.content_bytes(files, legacy_only=True))
        self.write(policy.MANIFEST, json.dumps(self.manifest))
        self.sha = self.commit()
        os.environ["GITHUB_SHA"] = self.sha
        return self.sha

    def fake_api(self, path, paginate=False, raw=False):
        if path.endswith("/branches/master"):
            return {"protected": True, "commit": {"sha": self.sha}}
        if "/commits/" in path and "/pulls" in path:
            return [dict(self.pull, merge_commit_sha=self.sha)]
        if path.endswith("/pulls/10"):
            return dict(self.pull, merge_commit_sha=self.sha)
        if "/reviews?" in path:
            return self.reviews
        if "/compare/" in path:
            return {"status": "ahead"}
        if raw:
            return self.spec
        self.fail("Unexpected GitHub API path: " + path)

    def verification(self, *, generation=False):
        if generation:
            repo, workflow, ref = policy.CONFIG_REPO, policy.GEN_WORKFLOW, policy.GEN_REF
            commit, run_id = self.manifest["config"]["commit"], self.manifest["generation"]["run_id"]
            predicate, predicate_type = self.manifest, policy.GEN_PREDICATE
            subjects = [{"name": f"plaid-go-v{self.manifest['sdk']['version']}-source", "digest": {"sha256": self.manifest["content_sha256"]}}]
        else:
            repo, workflow, ref = policy.SDK_REPO, policy.RELEASE_WORKFLOW, policy.RELEASE_REF
            commit, run_id = self.sha, "1234"
            predicate, _, subjects = artifacts.artifact_files(self.out, self.sha, self.control_commit, "1234", "1")
            predicate_type = policy.RELEASE_PREDICATE
        identity = f"https://github.com/{repo}/{workflow}@{ref}"
        certificate = {
            "subjectAlternativeName": identity, "issuer": "https://token.actions.githubusercontent.com",
            "buildSignerURI": identity, "buildSignerDigest": commit,
            "sourceRepositoryURI": f"https://github.com/{repo}", "sourceRepositoryDigest": commit,
            "sourceRepositoryRef": ref, "runnerEnvironment": "github-hosted",
            "sourceRepositoryVisibilityAtSigning": "public",
            "runInvocationURI": f"https://github.com/{repo}/actions/runs/{run_id}/attempts/1",
        }
        return [{"verificationResult": {"signature": {"certificate": certificate},
                 "verifiedTimestamps": [{"type": "Tlog", "timestamp": "fixture"}],
                 "statement": {"_type": "https://in-toto.io/Statement/v1", "predicateType": predicate_type,
                               "predicate": copy.deepcopy(predicate), "subject": copy.deepcopy(subjects)}}}]

    def mock_verifier(self, *, generation=False, mutate=None, fail=False):
        real_run = policy.run

        def verified_process(*arguments, **kwargs):
            if arguments[:3] == ("gh", "attestation", "verify"):
                self.gh_calls.append(arguments)
                if fail:
                    raise policy.ReleaseError("Fixture verifier rejected the signature")
                value = self.verification(generation=generation)
                if mutate:
                    mutate(value)
                return json.dumps(value).encode()
            return real_run(*arguments, **kwargs)

        return patch.object(policy, "run", side_effect=verified_process)

    def prepare_artifacts(self):
        archive = self.out / f"plaid-go-v{self.manifest['sdk']['version']}.zip"
        self.command(self.zip_tool, self.repo, self.manifest["sdk"]["module"], "v" + self.manifest["sdk"]["version"], archive)
        result = artifacts.make_release_manifest(self.repo, self.sha, self.control_commit, "1234", "1", archive)
        (self.out / "release.bundle.json").write_text('{"fixture_only":"not a real attestation"}\n')
        return result

    def publish_call(self, *, publish_requested=False, public_key=None):
        with patch.object(policy, "api", side_effect=self.fake_api), patch.object(publish, "api", side_effect=self.fake_api):
            return publish.publish(self.out, self.sha, self.control_commit, "1234", "1", public_key or self.public_key,
                                   publish_requested=publish_requested, remote=str(self.remote))

    def prepare_manual_predecessor(self):
        self.command("git", "-C", self.repo, "tag", "-a", "v49.0.0", self.control_commit, "-m", "Manual fixture release")
        self.command("git", "-C", self.repo, "push", "-q", "origin", "refs/tags/v49.0.0")
        self.manifest["sdk"]["version"] = "49.1.0"
        self.metadata["sdkVersion"] = "v49.1.0"
        self.write(".fern/metadata.json", json.dumps(self.metadata))
        self.refresh_content()
        self.command("git", "-C", self.repo, "push", "-q", "origin", "master")
        self.prepare_artifacts()
        return {"SDK_RELEASE_MANUAL_PREDECESSOR_TAG": "v49.0.0",
                "SDK_RELEASE_MANUAL_PREDECESSOR_COMMIT": self.control_commit,
                "SDK_RELEASE_MANUAL_PREDECESSOR_SPEC_COMMIT": "3" * 40}

    def test_canonical_content_and_legacy_match_contract(self):
        manifest, files = policy.validate_tree(self.repo, self.sha)
        expected = [[path, mode, policy.sha256(content)] for path, (mode, content) in sorted(files.items()) if path not in (policy.MANIFEST, policy.GEN_BUNDLE)]
        self.assertEqual(policy.content_bytes(files), json.dumps(expected, separators=(",", ":"), ensure_ascii=True).encode())
        self.assertEqual(policy.sha256(policy.content_bytes(files)), manifest["content_sha256"])
        self.assertNotIn(b"fixture_only", policy.content_bytes(files))

    def test_manifest_rejects_wrong_repository_version_sha_and_extra_fields(self):
        changes = [("sdk", "repository", "other/plaid-go"), ("sdk", "version", "49.01.0"),
                   ("config", "repository", "example/untrusted-config"), ("config", "commit", "main"),
                   ("source", "sha256", "not-a-digest"), ("generation", "workflow", "attacker.yml"),
                   ("legacy", "spec_sha256", "9" * 64)]
        for section, field, replacement in changes:
            with self.subTest(section=section, field=field):
                value = copy.deepcopy(self.manifest)
                value[section][field] = replacement
                with self.assertRaises(policy.ReleaseError):
                    policy.validate_manifest(value)
        value = copy.deepcopy(self.manifest)
        value["private_config"] = {"secret": "must not be published"}
        with self.assertRaises(policy.ReleaseError):
            policy.validate_manifest(value)

    def test_spec_and_config_cannot_be_signed_from_different_commits(self):
        manifest = copy.deepcopy(self.manifest)
        manifest["config"]["commit"] = "b" * 40
        with self.assertRaisesRegex(policy.ReleaseError, "same public commit"):
            policy.validate_manifest(manifest)

    def test_modified_sdk_bytes_are_rejected(self):
        self.write("sdk.go", "package plaid\nconst Tampered = true\n")
        sha = self.commit()
        with self.assertRaisesRegex(policy.ReleaseError, "changed after generation"):
            policy.validate_tree(self.repo, sha)

    def test_private_generator_config_cannot_be_published_in_metadata(self):
        self.metadata["generatorConfig"] = {"private": "configuration"}
        self.write(".fern/metadata.json", json.dumps(self.metadata))
        self.refresh_content()
        with self.assertRaisesRegex(policy.ReleaseError, "public Fern metadata"):
            policy.validate_tree(self.repo, self.sha)

    def test_extra_release_file_and_symlink_are_rejected(self):
        self.write(".sdk-release/unverified.json", "{}")
        with self.assertRaisesRegex(policy.ReleaseError, "Unexpected .sdk-release"):
            policy.tree_files(self.repo, self.commit())
        (self.repo / ".sdk-release/unverified.json").unlink()
        (self.repo / "symlink").symlink_to("/tmp")
        with self.assertRaisesRegex(policy.ReleaseError, "Symlinks"):
            policy.tree_files(self.repo, self.commit())

    def test_generation_verifier_binds_real_command_and_exact_predicate(self):
        with self.mock_verifier(generation=True):
            policy.verify_generation(self.repo, self.sha)
        call = self.gh_calls[0]
        for flag in ("--cert-identity", "--source-digest", "--signer-digest", "--source-ref", "--deny-self-hosted-runners"):
            self.assertIn(flag, call)
        self.assertNotIn("--no-public-good", call)
        self.assertIn(policy.GEN_PREDICATE, call)
        self.assertIn(self.manifest["config"]["commit"], call)
        def mutate(value):
            value[0]["verificationResult"]["statement"]["predicate"]["sdk"]["version"] = "49.0.1"
        with self.mock_verifier(generation=True, mutate=mutate), self.assertRaisesRegex(policy.ReleaseError, "predicate differs"):
            policy.verify_generation(self.repo, self.sha)

    def test_wrong_signer_ref_run_visibility_and_failed_crypto_are_rejected(self):
        replacements = {"subjectAlternativeName": "https://github.com/attacker/workflow.yml@refs/heads/main",
                        "sourceRepositoryDigest": "0" * 40, "sourceRepositoryRef": "refs/heads/unreviewed",
                        "runnerEnvironment": "self-hosted", "sourceRepositoryVisibilityAtSigning": "private",
                        "runInvocationURI": "https://github.com/plaid/plaid-openapi/actions/runs/111"}
        for field, replacement in replacements.items():
            with self.subTest(field=field):
                def mutate(value):
                    value[0]["verificationResult"]["signature"]["certificate"][field] = replacement
                with self.mock_verifier(generation=True, mutate=mutate), self.assertRaises(policy.ReleaseError):
                    policy.verify_generation(self.repo, self.sha)
        with self.mock_verifier(generation=True, fail=True), self.assertRaises(policy.ReleaseError):
            policy.verify_generation(self.repo, self.sha)

    def test_unreviewed_or_outside_approval_cannot_authorize_release(self):
        with patch.object(policy, "api", side_effect=self.fake_api):
            self.assertEqual(policy.require_approved_merge(self.sha), 10)
            self.reviews[0]["author_association"] = "CONTRIBUTOR"
            with self.assertRaisesRegex(policy.ReleaseError, "No human approval"):
                policy.require_approved_merge(self.sha)
            self.reviews[0]["author_association"] = "MEMBER"
            self.reviews[0]["commit_id"] = "0" * 40
            with self.assertRaisesRegex(policy.ReleaseError, "No human approval"):
                policy.require_approved_merge(self.sha)
            self.reviews[0]["state"] = "CHANGES_REQUESTED"
            with self.assertRaisesRegex(policy.ReleaseError, "changes-requested"):
                policy.require_approved_merge(self.sha)

    def automatic_merge(self):
        self.reviews = []
        self.pull.update(user={"type": "Bot", "id": 1001}, merged_by={"type": "Bot", "id": 1002})
        self.pull["head"].update(repo={"full_name": policy.SDK_REPO}, ref="sdk-generation/go-v49-" + "a" * 20)
        return {"SDK_AUTOMERGE_ENABLED": "true", "SDK_PR_APP_BOT_ID": "1001", "SDK_INTEGRATION_APP_BOT_ID": "1002"}

    def test_configured_automatic_merge_needs_no_human_approval(self):
        with patch.dict(os.environ, self.automatic_merge()), patch.object(policy, "api", side_effect=self.fake_api):
            self.assertEqual(policy.require_release_merge(self.sha), 10)
            with self.assertRaisesRegex(policy.ReleaseError, "No human approval"):
                policy.require_approved_merge(self.sha)

    def test_automatic_merge_disabled_or_wrong_identities_fail_closed(self):
        settings = self.automatic_merge()
        cases = [{"SDK_AUTOMERGE_ENABLED": "false"}, {"SDK_PR_APP_BOT_ID": ""},
                 {"SDK_PR_APP_BOT_ID": "1003"}, {"SDK_INTEGRATION_APP_BOT_ID": "1003"},
                 {"SDK_INTEGRATION_APP_BOT_ID": "1001"}]
        for override in cases:
            with self.subTest(override=override), patch.dict(os.environ, dict(settings, **override)), patch.object(policy, "api", side_effect=self.fake_api):
                with self.assertRaises(policy.ReleaseError):
                    policy.require_release_merge(self.sha)

    def test_automatic_merge_rejects_fork_human_edits_and_review_veto(self):
        settings = self.automatic_merge()
        with patch.dict(os.environ, settings), patch.object(policy, "api", side_effect=self.fake_api):
            self.pull["head"]["repo"]["full_name"] = "other/plaid-go"
            with self.assertRaises(policy.ReleaseError):
                policy.require_release_merge(self.sha)
            self.automatic_merge()
            self.pull["head"]["ref"] = "human-fix"
            with self.assertRaises(policy.ReleaseError):
                policy.require_release_merge(self.sha)
            self.automatic_merge()
            self.pull["merged_by"]["type"] = "User"
            with self.assertRaises(policy.ReleaseError):
                policy.require_release_merge(self.sha)
            self.automatic_merge()
            self.reviews = [{"user": {"type": "User", "login": "reviewer"}, "author_association": "MEMBER", "state": "CHANGES_REQUESTED"}]
            with self.assertRaisesRegex(policy.ReleaseError, "changes-requested"):
                policy.require_release_merge(self.sha)

    def test_real_signed_tag_after_automatic_merge_without_approval(self):
        settings = self.automatic_merge()
        settings.update(SDK_RELEASE_PUBLISH="true", SDK_RELEASE_PUSH_TOKEN="fixture-token",
                        SDK_RELEASE_SIGNING_PRIVATE_KEY=self.key.read_text())
        self.prepare_artifacts()
        with self.mock_verifier(), patch.dict(os.environ, settings):
            self.assertEqual(self.publish_call(publish_requested=True), "published")
            self.assertEqual(self.publish_call(publish_requested=True), "already-published")
        self.assertEqual(self.command("git", "-C", self.remote, "rev-parse", "v49.0.0^{}").decode().strip(), self.sha)

    def test_wrong_public_spec_is_rejected(self):
        with patch.object(policy, "api", side_effect=self.fake_api):
            policy.verify_public_spec(self.manifest)
            self.spec = b"tampered public spec"
            with self.assertRaisesRegex(policy.ReleaseError, "Public spec bytes"):
                policy.verify_public_spec(self.manifest)

    def test_workflow_context_and_control_are_pinned(self):
        policy.require_context(os.environ, self.sha, self.control_commit)
        policy.require_control_workflow(self.repo, self.sha, self.control, self.control_commit)
        for field, replacement in (("SDK_RELEASE_ENABLED", "false"), ("GITHUB_REPOSITORY", "fork/plaid-go"),
                                   ("GITHUB_EVENT_NAME", "pull_request_target"), ("GITHUB_REF_PROTECTED", "false")):
            with self.subTest(field=field), patch.dict(os.environ, {field: replacement}), self.assertRaises(policy.ReleaseError):
                policy.require_context(os.environ, self.sha, self.control_commit)
        self.write(policy.RELEASE_WORKFLOW, "name: malicious-workflow\n")
        with self.assertRaisesRegex(policy.ReleaseError, "workflow changed"):
            policy.require_control_workflow(self.repo, self.commit(), self.control, self.control_commit)

    def test_disabled_pipeline_allows_only_explicit_trusted_manual_dry_run(self):
        settings = {"SDK_RELEASE_ENABLED": "false", "GITHUB_EVENT_NAME": "workflow_dispatch", "SDK_RELEASE_DRY_RUN": "true"}
        with patch.dict(os.environ, settings):
            policy.require_context(os.environ, self.sha, self.control_commit)
            for change in ({"GITHUB_EVENT_NAME": "push"}, {"SDK_RELEASE_DRY_RUN": "false"},
                           {"GITHUB_REF_PROTECTED": "false"}, {"GITHUB_REF": "refs/heads/unreviewed"}):
                with self.subTest(change=change), patch.dict(os.environ, change), self.assertRaises(policy.ReleaseError):
                    policy.require_context(os.environ, self.sha, self.control_commit)

    def test_disabled_or_explicit_dry_run_cannot_enter_publication(self):
        self.prepare_artifacts()
        settings = {"SDK_RELEASE_ENABLED": "false", "GITHUB_EVENT_NAME": "workflow_dispatch",
                    "SDK_RELEASE_DRY_RUN": "true", "SDK_RELEASE_PUBLISH": "true"}
        with patch.dict(os.environ, settings), self.mock_verifier():
            self.assertEqual(self.publish_call(), "dry-run")
            with self.assertRaisesRegex(policy.ReleaseError, "pipeline is disabled"):
                self.publish_call(publish_requested=True)
            with patch.dict(os.environ, {"SDK_RELEASE_ENABLED": "true"}), self.assertRaisesRegex(policy.ReleaseError, "dry run cannot publish"):
                self.publish_call(publish_requested=True)
        self.assertEqual(self.command("git", "-C", self.remote, "tag"), b"")

    def test_real_canonical_module_zip_and_release_subjects(self):
        self.prepare_artifacts()
        release, _, subjects = artifacts.artifact_files(self.out, self.sha, self.control_commit, "1234", "1")
        archive = self.out / release["module_zip"]["name"]
        second = self.root / "second.zip"
        self.command(self.zip_tool, self.repo, self.manifest["sdk"]["module"], "v49.0.0", second)
        self.assertEqual(archive.read_bytes(), second.read_bytes())
        with zipfile.ZipFile(archive) as zipped:
            self.assertIn("github.com/plaid/plaid-go/v49@v49.0.0/legacy/client.go", zipped.namelist())
        with self.mock_verifier():
            artifacts.verify_release(self.out, self.sha, self.control_commit, "1234", "1")
        self.assertEqual(len(self.gh_calls), 2)
        self.assertEqual({item["name"] for item in subjects}, {"plaid-go-v49.0.0.zip", "release-manifest.json"})

    def test_modified_zip_bundle_manifest_or_wrong_run_is_rejected(self):
        self.prepare_artifacts()
        for name in ("plaid-go-v49.0.0.zip", "generation-manifest.json", "generation.bundle.json"):
            path = self.out / name
            original = path.read_bytes()
            path.write_bytes(original + b"tampered")
            with self.subTest(name=name), self.assertRaises((policy.ReleaseError, ValueError)):
                artifacts.artifact_files(self.out, self.sha, self.control_commit, "1234", "1")
            path.write_bytes(original)
        with self.assertRaises(policy.ReleaseError):
            artifacts.artifact_files(self.out, self.sha, self.control_commit, "9999", "1")
        with self.mock_verifier(fail=True), self.assertRaises(policy.ReleaseError):
            artifacts.verify_release(self.out, self.sha, self.control_commit, "1234", "1")

    def test_dry_run_does_not_access_signing_secret_or_publish(self):
        self.prepare_artifacts()
        with self.mock_verifier():
            self.assertEqual(self.publish_call(), "dry-run")
        self.assertEqual(self.command("git", "-C", self.remote, "tag"), b"")

    def test_real_signed_tag_publication_and_idempotent_retry(self):
        self.prepare_artifacts()
        settings = {"SDK_RELEASE_PUBLISH": "true", "SDK_RELEASE_PUSH_TOKEN": "fixture-token",
                    "SDK_RELEASE_SIGNING_PRIVATE_KEY": self.key.read_text()}
        with self.mock_verifier(), patch.dict(os.environ, settings):
            self.assertEqual(self.publish_call(publish_requested=True), "published")
            self.assertEqual(self.publish_call(publish_requested=True), "already-published")
        self.assertEqual(self.command("git", "-C", self.remote, "rev-parse", "v49.0.0^{}").decode().strip(), self.sha)
        self.assertEqual(self.command("git", "-C", self.remote, "cat-file", "-t", "v49.0.0").strip(), b"tag")

    def test_wrong_private_key_cannot_publish(self):
        self.prepare_artifacts()
        other = self.root / "other-key"
        self.command("ssh-keygen", "-q", "-t", "ed25519", "-N", "", "-f", other)
        with self.mock_verifier(), patch.dict(os.environ, {"SDK_RELEASE_PUBLISH": "true", "SDK_RELEASE_PUSH_TOKEN": "fixture", "SDK_RELEASE_SIGNING_PRIVATE_KEY": other.read_text()}):
            with self.assertRaisesRegex(policy.ReleaseError, "private key does not match"):
                self.publish_call(publish_requested=True)
        self.assertEqual(self.command("git", "-C", self.remote, "tag"), b"")

    def test_conflicting_or_unsigned_existing_tag_is_rejected(self):
        self.prepare_artifacts()
        self.command("git", "-C", self.repo, "tag", "v49.0.0", self.control_commit)
        self.command("git", "-C", self.repo, "push", "-q", "origin", "refs/tags/v49.0.0")
        with self.mock_verifier(), self.assertRaisesRegex(policy.ReleaseError, "not annotated"):
            self.publish_call()

    def test_validly_signed_tag_for_another_commit_is_rejected(self):
        self.prepare_artifacts()
        publish.install_verification_key(self.repo, self.public_key, self.root)
        self.command("git", "-C", self.repo, "config", "user.signingkey", self.key)
        self.command("git", "-C", self.repo, "tag", "-s", "v49.0.0", self.control_commit, "-m", "Conflicting fixture")
        self.command("git", "-C", self.repo, "push", "-q", "origin", "refs/tags/v49.0.0")
        with self.mock_verifier(), self.assertRaisesRegex(policy.ReleaseError, "another commit"):
            self.publish_call()

    def test_lower_version_is_rejected_even_when_requested_tag_is_absent(self):
        self.prepare_artifacts()
        self.command("git", "-C", self.repo, "tag", "v49.2.0")
        self.command("git", "-C", self.repo, "push", "-q", "origin", "refs/tags/v49.2.0")
        with self.mock_verifier(), self.assertRaisesRegex(policy.ReleaseError, "newer or equal SDK version"):
            self.publish_call()

    def test_manual_first_transition_requires_all_exact_reviewed_pins(self):
        settings = self.prepare_manual_predecessor()
        with self.mock_verifier(), self.assertRaisesRegex(policy.ReleaseError, "reviewed manual predecessor commit"):
            self.publish_call()
        with self.mock_verifier(), patch.dict(os.environ, settings):
            self.assertEqual(self.publish_call(), "dry-run")
            for field in settings:
                with self.subTest(missing=field), patch.dict(os.environ, {field: ""}), self.assertRaises(policy.ReleaseError):
                    self.publish_call()
            for field, incorrect in (("SDK_RELEASE_MANUAL_PREDECESSOR_TAG", "v49.0.1"),
                                     ("SDK_RELEASE_MANUAL_PREDECESSOR_COMMIT", "4" * 40)):
                with self.subTest(wrong=field), patch.dict(os.environ, {field: incorrect}), self.assertRaisesRegex(policy.ReleaseError, "unapproved manual predecessor"):
                    self.publish_call()
        self.assertEqual(self.command("git", "-C", self.remote, "tag").strip(), b"v49.0.0")

    def test_manual_first_transition_rejects_public_spec_rollback(self):
        settings = self.prepare_manual_predecessor()
        real_api = self.fake_api
        checked = []

        def reject_ancestry(path, **kwargs):
            if "/compare/" in path:
                checked.append(path)
                return {"status": "diverged"}
            return real_api(path, **kwargs)

        with self.mock_verifier(), patch.dict(os.environ, settings), patch.object(policy, "api", side_effect=real_api), patch.object(publish, "api", side_effect=reject_ancestry):
            with self.assertRaisesRegex(policy.ReleaseError, "older than or diverged"):
                publish.publish(self.out, self.sha, self.control_commit, "1234", "1", self.public_key, remote=str(self.remote))
        self.assertEqual(checked, [f"repos/{policy.SPEC_REPO}/compare/{'3' * 40}...{'a' * 40}"])

    def test_manual_first_transition_can_publish_only_the_new_signed_version(self):
        settings = self.prepare_manual_predecessor()
        settings.update({"SDK_RELEASE_PUBLISH": "true", "SDK_RELEASE_PUSH_TOKEN": "fixture-token",
                         "SDK_RELEASE_SIGNING_PRIVATE_KEY": self.key.read_text()})
        with self.mock_verifier(), patch.dict(os.environ, settings):
            self.assertEqual(self.publish_call(publish_requested=True), "published")
            self.assertEqual(self.publish_call(publish_requested=True), "already-published")
        self.assertEqual(self.command("git", "-C", self.remote, "rev-parse", "v49.0.0^{}").decode().strip(), self.control_commit)
        self.assertEqual(self.command("git", "-C", self.remote, "rev-parse", "v49.1.0^{}").decode().strip(), self.sha)

    def test_git_url_rewrites_fail_before_network_in_all_config_scopes(self):
        remote = policy.PUBLIC_REMOTE
        key = "url.https://unapproved.example/.insteadOf"
        for scope in ("local", "global", "environment"):
            with self.subTest(scope=scope):
                settings = {}
                if scope == "local":
                    self.command("git", "-C", self.repo, "config", key, remote)
                elif scope == "global":
                    config = self.root / "global.gitconfig"
                    self.command("git", "config", "--file", config, key, remote)
                    settings["GIT_CONFIG_GLOBAL"] = str(config)
                else:
                    settings = {"GIT_CONFIG_COUNT": "1", "GIT_CONFIG_KEY_0": key, "GIT_CONFIG_VALUE_0": remote}
                with patch.dict(os.environ, settings), patch.object(publish, "run", side_effect=AssertionError("Network must not run")):
                    with self.assertRaisesRegex(policy.ReleaseError, "URL rewrite"):
                        publish.network_git(self.repo, remote, "ls-remote", remote)
                if scope == "local":
                    self.command("git", "-C", self.repo, "config", "--unset", key)

    def test_semver_order_including_prereleases(self):
        values = ["v49.0.0-alpha.2", "v49.0.0-alpha.11", "v49.0.0-beta", "v49.0.0", "v49.1.0", "v50.0.0"]
        self.assertEqual(sorted(reversed(values), key=publish.semver_key), values)
        self.assertIsNone(publish.semver_key("v49.01.0"))


if __name__ == "__main__":
    unittest.main()
