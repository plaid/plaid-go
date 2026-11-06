"""Integration policy tests: real Git trees, mocked GitHub boundaries, no writes."""

import copy
import json
import os
from pathlib import Path
import subprocess
import tempfile
import unittest
from unittest.mock import patch

import integrate
import policy


class IntegrationTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix="sdk-integration-test-")
        self.addCleanup(self.temp.cleanup)
        self.repo = Path(self.temp.name) / "sdk"
        self.control = Path(self.temp.name) / "control"
        self.env = patch.dict(os.environ, {
            "HOME": os.environ["HOME"], "PATH": os.environ["PATH"],
            "GIT_CONFIG_GLOBAL": "/dev/null", "GIT_CONFIG_NOSYSTEM": "1",
            "GIT_AUTHOR_NAME": "Fixture", "GIT_COMMITTER_NAME": "Fixture",
            "GIT_AUTHOR_EMAIL": "fixture@example.com", "GIT_COMMITTER_EMAIL": "fixture@example.com",
            "SDK_AUTOMERGE_ENABLED": "true", "SDK_PR_APP_BOT_ID": "123", "SDK_CIRCLECI_STATUS_ACTOR_ID": "456",
            "SDK_INTEGRATION_APP_BOT_ID": "789", "SDK_REVIEW_RULESET_ID": "12345",
            "GITHUB_REPOSITORY": policy.SDK_REPO, "GITHUB_REF": "refs/heads/master", "GITHUB_REF_PROTECTED": "true",
            "GITHUB_WORKFLOW_REF": f"{policy.SDK_REPO}/{integrate.WORKFLOW}@refs/heads/master", "GITHUB_EVENT_NAME": "schedule",
            "GH_TOKEN": "read-only-fixture",
        }, clear=True)
        self.env.start()
        self.addCleanup(self.env.stop)
        self.cmd("git", "init", "-q", "--initial-branch=master", self.repo)
        for path, value in {integrate.WORKFLOW: "name: trusted\n", ".github/CODEOWNERS": "* @plaid/developer-experience\n",
                            "scripts/release/tool.py": "# trusted\n", ".circleci/config.yml": "trusted-ci\n",
                            ".fernignore": ".github\nscripts/release\n", "LICENSE": "license\n", "README.md": "docs\n",
                            "internal/caller.go": "package internal\n", "CHANGELOG.md": "## [v49.0.0]\n\nOld release.\n",
                            "go.mod": "module github.com/plaid/plaid-go/v49\n", "sdk.go": "package plaid\n", ".fern/metadata.json": '{"generatorVersion":"1.70.3"}'}.items():
            self.write(path, value)
        self.base = self.commit()
        self.cmd("git", "clone", "-q", self.repo, self.control)
        os.environ["GITHUB_SHA"] = self.base
        self.spec = b"openapi: 3.0.3\ninfo:\n  title: Plaid\n  version: 2020-09-14_1.767.1\n"
        self.manifest = {"schema_version": 1, "sdk": {"version": "49.1.0"}, "source": {"repository": policy.SPEC_REPO,
                         "commit": "a" * 40, "path": "2020-09-14.yml", "sha256": policy.sha256(self.spec)},
                         "config": {"sha256": "b" * 64}}
        self.write("sdk.go", "package plaid\nconst Changed = true\n")
        self.head = self.commit()
        branch_identity = policy.sha256(policy.canonical(["49.1.0", policy.sha256(self.spec), "b" * 64, self.base]))[:20]
        self.pull = {"number": 42, "state": "open", "draft": False, "merged": False, "commits": 1,
                     "user": {"type": "Bot", "id": 123},
                     "base": {"sha": self.base, "ref": "master", "repo": {"full_name": policy.SDK_REPO}},
                     "head": {"sha": self.head, "ref": "sdk-generation/go-v49-" + branch_identity,
                              "repo": {"full_name": policy.SDK_REPO}}}
        self.branch = {"protected": True, "commit": {"sha": self.base}}
        self.protection = {"required_status_checks": {"strict": True, "contexts": [integrate.CI_CONTEXT],
                          "checks": [{"context": integrate.VERIFIED_CONTEXT, "app_id": integrate.ACTIONS_APP_ID}]}}
        self.effective_rules = [{"type": "pull_request", "ruleset_id": 12345}]
        self.ruleset = {"id": 12345, "source_type": "Repository", "source": policy.SDK_REPO,
                        "target": "branch", "enforcement": "active",
                        "conditions": {"ref_name": {"include": ["refs/heads/master"], "exclude": []}},
                        "rules": [{"type": "pull_request", "parameters": {"required_approving_review_count": 1,
                                   "require_code_owner_review": True, "dismiss_stale_reviews_on_push": True}}],
                        "current_user_can_bypass": "pull_requests_only"}
        self.ruleset_tokens = []
        self.statuses = [{"context": integrate.CI_CONTEXT, "state": "success", "creator": {"id": 456},
                          "target_url": "https://circleci.com/gh/plaid/plaid-go/123"}]
        self.statuses.append({"context": integrate.VERIFIED_CONTEXT, "state": "success",
                              "creator": {"id": integrate.ACTIONS_BOT_ID},
                              "target_url": "https://github.com/plaid/plaid-go/actions/runs/123"})
        self.combined = {"sha": self.head, "state": "success"}
        self.reviews = []
        self.api_calls = []
        self.checks = []
        self.patches = [patch.object(integrate, "api", side_effect=self.api),
                        patch.object(integrate, "fetch_objects"),
                        patch.object(integrate, "verify_generation", side_effect=self.verified_generation),
                        patch.object(integrate, "verify_public_spec")]
        for item in self.patches:
            item.start()
            self.addCleanup(item.stop)

    def cmd(self, *args):
        result = subprocess.run([str(arg) for arg in args], capture_output=True)
        self.assertEqual(result.returncode, 0, result.stderr.decode())
        return result.stdout.decode().strip()

    def write(self, name, value):
        target = self.repo / name
        target.parent.mkdir(parents=True, exist_ok=True)
        target.write_text(value)

    def commit(self):
        self.cmd("git", "-C", self.repo, "add", "-A")
        self.cmd("git", "-C", self.repo, "commit", "-q", "-m", "Fixture")
        return self.cmd("git", "-C", self.repo, "rev-parse", "HEAD")

    def amend_candidate(self, path, content):
        self.write(path, content)
        self.cmd("git", "-C", self.repo, "add", "-A")
        self.cmd("git", "-C", self.repo, "commit", "-q", "--amend", "--no-edit")
        self.head = self.cmd("git", "-C", self.repo, "rev-parse", "HEAD")
        self.pull["head"]["sha"] = self.head
        self.combined["sha"] = self.head

    def api(self, path, **kwargs):
        self.api_calls.append(path)
        if path.endswith("/branches/master"):
            return copy.deepcopy(self.branch)
        if path.endswith("/branches/master/protection"):
            return copy.deepcopy(self.protection)
        if "/rules/branches/master?" in path:
            return copy.deepcopy(self.effective_rules)
        if path.endswith("/rulesets/12345"):
            self.ruleset_tokens.append(kwargs.get("token"))
            return copy.deepcopy(self.ruleset)
        if path.endswith("/pulls/42"):
            return copy.deepcopy(self.pull)
        if "/pulls?" in path:
            return [copy.deepcopy(self.pull)]
        if path.endswith("/status"):
            return copy.deepcopy(self.combined)
        if "/statuses?" in path:
            return copy.deepcopy(self.statuses)
        if "/reviews?" in path:
            return copy.deepcopy(self.reviews)
        if "/contents/" in path:
            return self.spec
        if "/git/commits/" in path:
            return {"parents": [{"sha": self.base}], "tree": {"sha": self.tree()}}
        self.fail("Unexpected API path: " + path)

    def verified_generation(self, repo, head):
        self.checks.append((head, os.environ.get("GH_TOKEN")))
        return self.manifest, policy.tree_files(repo, head)

    def check(self, **kwargs):
        return integrate.check(self.repo, self.control, self.base, 42, **kwargs)

    def tree(self):
        return self.cmd("git", "-C", self.repo, "rev-parse", f"{self.head}^{{tree}}")

    def test_check_only_validates_exact_git_tree_without_write(self):
        with patch.object(integrate, "run", side_effect=AssertionError("No write command expected")):
            result = self.check()
        self.assertEqual(result["head"], self.head)
        self.assertEqual(result["tree"], self.tree())
        self.assertEqual(self.checks, [(self.head, "read-only-fixture")])
        self.assertNotIn("/merge", " ".join(self.api_calls))

    def test_branch_identity_binds_the_current_sdk_base(self):
        def branch(base):
            identity = policy.sha256(policy.canonical([
                self.manifest["sdk"]["version"], self.manifest["source"]["sha256"],
                self.manifest["config"]["sha256"], base]))[:20]
            return "sdk-generation/go-v49-" + identity
        self.pull["head"]["ref"] = branch(self.base)
        self.assertEqual(self.check()["head"], self.head)
        self.pull["head"]["ref"] = branch("f" * 40)
        with self.assertRaisesRegex(policy.ReleaseError, "branch identity"):
            self.check()

    def test_discovery_accepts_only_the_pinned_bot_and_same_repo_branch(self):
        self.assertEqual(integrate.discover(), 42)
        for field, value in (("id", 999), ("type", "User")):
            old = self.pull["user"][field]
            self.pull["user"][field] = value
            self.assertIsNone(integrate.discover())
            self.pull["user"][field] = old
        self.pull["head"]["repo"]["full_name"] = "attacker/plaid-go"
        self.assertIsNone(integrate.discover())

    def test_signature_failure_stops_before_merge(self):
        with patch.object(integrate, "verify_generation", side_effect=policy.ReleaseError("Invalid signature")):
            with self.assertRaisesRegex(policy.ReleaseError, "Invalid signature"):
                self.check()

    def test_public_verification_uses_normal_read_token(self):
        with patch.dict(os.environ, {"SDK_ATTESTATION_READ_TOKEN": "private-read-fixture"}):
            self.check()
        self.assertEqual(self.checks, [(self.head, "read-only-fixture")])
        self.assertEqual(os.environ["GH_TOKEN"], "read-only-fixture")

    def test_changed_branch_identity_or_multiple_commits_is_rejected(self):
        self.pull["head"]["ref"] = "sdk-generation/go-v49-" + "0" * 20
        with self.assertRaisesRegex(policy.ReleaseError, "signed inputs"):
            self.check()
        self.pull["commits"] = 2
        with self.assertRaisesRegex(policy.ReleaseError, "exactly one commit"):
            self.check()

    def test_pr_head_or_base_movement_is_rejected(self):
        with self.assertRaisesRegex(policy.ReleaseError, "head changed"):
            self.check(expected_head="c" * 40)
        self.branch["commit"]["sha"] = "d" * 40
        with self.assertRaisesRegex(policy.ReleaseError, "base moved"):
            self.check()

    def test_extra_commit_cannot_masquerade_as_single_generated_change(self):
        self.write("extra.go", "package plaid\n")
        self.head = self.commit()
        self.pull["head"]["sha"] = self.head
        self.combined["sha"] = self.head
        with self.assertRaisesRegex(policy.ReleaseError, "direct child"):
            self.check()

    def test_protected_files_cannot_be_added_deleted_or_changed(self):
        before = policy.tree_files(self.repo, self.base)
        for path in ("scripts/release/tool.py", ".github/new.yml", ".circleci/new.yml", "LICENSE", ".fernignore", "internal/caller.go", "transfer_diligence_upload_test.go"):
            candidate = copy.deepcopy(before)
            candidate[path] = ("100644", b"changed")
            with self.subTest(path=path), self.assertRaisesRegex(policy.ReleaseError, "protected"):
                integrate.preserved_content(before, candidate, self.manifest)
        candidate = copy.deepcopy(before)
        del candidate["LICENSE"]
        with self.assertRaisesRegex(policy.ReleaseError, "protected"):
            integrate.preserved_content(before, candidate, self.manifest)

    def test_unexpected_signed_script_is_not_an_approved_generated_path(self):
        before = policy.tree_files(self.repo, self.base)
        candidate = copy.deepcopy(before)
        candidate["setup.sh"] = ("100644", b"echo unexpected\n")
        with self.assertRaisesRegex(policy.ReleaseError, "approved generated"):
            integrate.preserved_content(before, candidate, self.manifest)

    def test_generated_path_policy_allows_new_endpoints_and_generated_deletions(self):
        before = policy.tree_files(self.repo, self.base)
        candidate = copy.deepcopy(before)
        for name in ("newproduct/client.go", "newproduct/report/raw_client.go", "new_model.go"):
            candidate[name] = ("100644", b"package fixture\n")
        for name in integrate.GENERATED_RESOURCES:
            if name not in candidate:
                candidate[name] = ("100644", b"fixture")
        del candidate["sdk.go"]
        integrate.preserved_content(before, candidate, self.manifest)

    def test_generated_path_policy_rejects_config_scripts_and_executable_modes(self):
        before = policy.tree_files(self.repo, self.base)
        for name in ("Makefile", ".env", "newproduct/config.json", "newproduct/setup.py", "scripts/tool.go", "NewClient.go"):
            candidate = copy.deepcopy(before)
            candidate[name] = ("100644", b"unexpected")
            with self.subTest(name=name), self.assertRaisesRegex(policy.ReleaseError, "approved generated"):
                integrate.preserved_content(before, candidate, self.manifest)
        for name in ("newproduct/client.go", "go.mod", "reference.md"):
            candidate = copy.deepcopy(before)
            candidate[name] = ("100755", b"unexpected")
            with self.subTest(mode=name), self.assertRaisesRegex(policy.ReleaseError, "approved generated"):
                integrate.preserved_content(before, candidate, self.manifest)

    def test_unlisted_existing_files_are_frozen_including_deletion(self):
        before = policy.tree_files(self.repo, self.base)
        for name in (".gitattributes", ".fern/replay.lock", "CONTRIBUTING.md", "MIGRATION.md", "wiremock/docker-compose.test.yml"):
            before[name] = ("100644", b"maintained")
            integrate.preserved_content(before, before, self.manifest)
            for action in ("edit", "delete", "chmod"):
                candidate = copy.deepcopy(before)
                if action == "delete":
                    del candidate[name]
                else:
                    candidate[name] = ("100755" if action == "chmod" else "100644", b"changed")
                with self.subTest(name=name, action=action), self.assertRaisesRegex(policy.ReleaseError, "approved generated"):
                    integrate.preserved_content(before, candidate, self.manifest)

    def test_legacy_tests_are_maintained_with_only_major_import_rewrites(self):
        before = policy.tree_files(self.repo, self.base)
        name = "legacy/tests/client_test.go"
        before[name] = ("100644", b'package tests\nimport "github.com/plaid/plaid-go/v49/legacy"\n')
        candidate = copy.deepcopy(before)
        candidate[name] = ("100644", before[name][1].replace(b"/v49/", b"/v50/"))
        manifest = copy.deepcopy(self.manifest)
        manifest["sdk"]["version"] = "50.0.0"
        integrate.preserved_content(before, candidate, manifest)
        for action in ("add", "delete", "edit", "chmod"):
            changed = copy.deepcopy(candidate)
            if action == "add":
                changed["legacy/tests/new_test.go"] = ("100644", b"package tests\n")
            elif action == "delete":
                del changed[name]
            else:
                changed[name] = ("100755" if action == "chmod" else "100644", b"package changed\n")
            with self.subTest(action=action), self.assertRaisesRegex(policy.ReleaseError, "protected"):
                integrate.preserved_content(before, changed, manifest)
        with self.assertRaisesRegex(policy.ReleaseError, "protected"):
            integrate.preserved_content(before, candidate, self.manifest)

    def test_only_exact_changelog_prepend_preserves_history(self):
        before = policy.tree_files(self.repo, self.base)
        candidate = copy.deepcopy(before)
        prefix = b"## [v49.1.0]\n\nRegenerate the new client and legacy package from OpenAPI `2020-09-14_1.767.1` using Fern Go `1.70.3` for the new client.\n\n"
        candidate["CHANGELOG.md"] = ("100644", prefix + before["CHANGELOG.md"][1])
        integrate.preserved_content(before, candidate, self.manifest)
        candidate["CHANGELOG.md"] = ("100644", prefix + b"history removed")
        with self.assertRaisesRegex(policy.ReleaseError, "exact release entry"):
            integrate.preserved_content(before, candidate, self.manifest)

    def test_major_transition_allows_only_exact_maintained_module_substitutions(self):
        before = policy.tree_files(self.repo, self.base)
        before["internal/caller.go"] = ("100644", b'package internal\nimport "github.com/plaid/plaid-go/v49/core"\n')
        before["README.md"] = ("100644", b"go get github.com/plaid/plaid-go/v49@latest\n")
        candidate = copy.deepcopy(before)
        for path in ("go.mod", "internal/caller.go", "README.md"):
            mode, data = candidate[path]
            candidate[path] = (mode, data.replace(b"plaid-go/v49", b"plaid-go/v50"))
        manifest = copy.deepcopy(self.manifest)
        manifest["sdk"].update(version="50.0.0", module="github.com/plaid/plaid-go/v50")
        integrate.preserved_content(before, candidate, manifest)
        for path in ("internal/caller.go", "README.md", "scripts/release/tool.py", ".circleci/config.yml"):
            changed = copy.deepcopy(candidate)
            mode, data = changed[path]
            changed[path] = (mode, data + b"unexpected change\n")
            with self.subTest(path=path), self.assertRaisesRegex(policy.ReleaseError, "protected"):
                integrate.preserved_content(before, changed, manifest)
        changed = copy.deepcopy(candidate)
        changed["internal/caller.go"] = ("100755", candidate["internal/caller.go"][1])
        with self.assertRaisesRegex(policy.ReleaseError, "protected"):
            integrate.preserved_content(before, changed, manifest)

    def test_major_candidate_passes_full_integration_policy_with_exact_branch_identity(self):
        self.manifest["sdk"].update(version="50.0.0", module="github.com/plaid/plaid-go/v50")
        self.amend_candidate("go.mod", "module github.com/plaid/plaid-go/v50\n")
        identity = policy.sha256(policy.canonical(["50.0.0", policy.sha256(self.spec), "b" * 64, self.base]))[:20]
        self.pull["head"]["ref"] = "sdk-generation/go-v50-" + identity
        self.assertEqual(self.check()["version"], "50.0.0")
        self.statuses[0]["state"] = "failure"
        with self.assertRaisesRegex(policy.ReleaseError, "CircleCI failed"):
            self.check()

    def test_compatible_release_after_major_requires_unchanged_maintained_paths(self):
        before = policy.tree_files(self.repo, self.base)
        before["go.mod"] = ("100644", b"module github.com/plaid/plaid-go/v50\n")
        before["internal/caller.go"] = ("100644", b'package internal\nimport "github.com/plaid/plaid-go/v50/core"\n')
        manifest = copy.deepcopy(self.manifest)
        manifest["sdk"].update(version="50.1.0", module="github.com/plaid/plaid-go/v50")
        integrate.preserved_content(before, before, manifest)
        candidate = copy.deepcopy(before)
        candidate["internal/caller.go"] = ("100644", before["internal/caller.go"][1].replace(b"/v50/", b"/v51/"))
        with self.assertRaisesRegex(policy.ReleaseError, "protected"):
            integrate.preserved_content(before, candidate, manifest)

    def test_major_transition_rejects_skips_downgrades_and_noninitial_versions(self):
        before = policy.tree_files(self.repo, self.base)
        for version in ("48.1.0", "51.0.0", "50.1.0", "50.0.1", "50.0.0-rc.1"):
            manifest = copy.deepcopy(self.manifest)
            manifest["sdk"]["version"] = version
            with self.subTest(version=version), self.assertRaises(policy.ReleaseError):
                integrate.preserved_content(before, before, manifest)

    def test_module_rewrite_never_changes_similarly_named_dependencies(self):
        before = policy.tree_files(self.repo, self.base)
        text = b'github.com/plaid/plaid-go/v49/core github.com/plaid/plaid-go/v490/core github.com/plaid/plaid-go/v49-extra/core'
        before["README.md"] = ("100644", text)
        candidate = copy.deepcopy(before)
        candidate["README.md"] = ("100644", text.replace(b"/v49/core", b"/v50/core"))
        manifest = copy.deepcopy(self.manifest)
        manifest["sdk"]["version"] = "50.0.0"
        integrate.preserved_content(before, candidate, manifest)
        candidate["README.md"] = ("100644", text.replace(b"/v49", b"/v50"))
        with self.assertRaisesRegex(policy.ReleaseError, "protected"):
            integrate.preserved_content(before, candidate, manifest)
        with self.assertRaisesRegex(policy.ReleaseError, "protected"):
            integrate.preserved_content(before, candidate, self.manifest)

    def test_strict_required_circleci_protection_is_mandatory(self):
        for change in ({"strict": False}, {"contexts": []}):
            original = copy.deepcopy(self.protection)
            self.protection["required_status_checks"].update(change)
            with self.subTest(change=change), self.assertRaisesRegex(policy.ReleaseError, "strict"):
                self.check()
            self.protection = original

    def test_status_only_checks_remain_available_without_automatic_merging(self):
        with patch.dict(os.environ, {"SDK_AUTOMERGE_ENABLED": "false"}):
            verified = self.check(status_only=True)
            with patch.object(integrate, "run", side_effect=AssertionError("Must not merge")), self.assertRaisesRegex(policy.ReleaseError, "disabled"):
                integrate.merge(verified)

    def test_native_verification_status_requires_github_actions_binding(self):
        for app_id in (None, -1, 999):
            self.protection["required_status_checks"]["checks"][0]["app_id"] = app_id
            with self.subTest(app_id=app_id), self.assertRaisesRegex(policy.ReleaseError, "bound to GitHub Actions"):
                self.check()

    def test_own_pending_verification_is_not_a_circular_ci_dependency(self):
        self.statuses[1]["state"] = "pending"
        self.combined["state"] = "pending"
        self.check()
        self.statuses.insert(0, {"context": "security-scan", "state": "failure"})
        with self.assertRaisesRegex(policy.ReleaseError, "reported commit status failed"):
            self.check()

    def test_unverified_or_forged_status_cannot_reach_merge(self):
        verified = self.check()
        for change in ({"state": "pending"}, {"state": "failure"}, {"creator": {"id": 789}},
                       {"target_url": "https://attacker.example/verified"}):
            original = copy.deepcopy(self.statuses[1])
            self.statuses[1].update(change)
            with self.subTest(change=change), patch.object(integrate, "run", side_effect=AssertionError("Must not merge")), self.assertRaises(policy.ReleaseError):
                integrate.merge(verified)
            self.statuses[1] = original
        self.statuses.pop()
        with patch.object(integrate, "run", side_effect=AssertionError("Must not merge")), self.assertRaises(policy.ReleaseError):
            integrate.merge(verified)

    def test_review_ruleset_must_be_active_applicable_and_approval_only(self):
        original = copy.deepcopy(self.ruleset)
        mutations = [
            lambda rule: rule.update(id=999),
            lambda rule: rule.update(source="other/plaid-go"),
            lambda rule: rule.update(source_type="Organization"),
            lambda rule: rule.update(target="tag"),
            lambda rule: rule.update(enforcement="disabled"),
            lambda rule: rule.update(enforcement="evaluate"),
            lambda rule: rule["conditions"]["ref_name"].update(include=["refs/heads/other"]),
            lambda rule: rule["conditions"]["ref_name"].update(exclude=["refs/heads/master"]),
            lambda rule: rule["rules"].append({"type": "required_status_checks"}),
            lambda rule: rule["rules"][0]["parameters"].update(required_approving_review_count=0),
            lambda rule: rule["rules"][0]["parameters"].update(require_code_owner_review=False),
            lambda rule: rule["rules"][0]["parameters"].update(dismiss_stale_reviews_on_push=False),
        ]
        for index, mutate in enumerate(mutations):
            self.ruleset = copy.deepcopy(original)
            mutate(self.ruleset)
            with self.subTest(case=index), self.assertRaises(policy.ReleaseError):
                self.check()
        self.ruleset = original
        self.effective_rules = [{"type": "pull_request", "ruleset_id": 999}]
        with self.assertRaisesRegex(policy.ReleaseError, "not active on master"):
            self.check()

    def test_missing_review_ruleset_id_or_classic_review_bypass_is_rejected(self):
        for value in ("", "0", "not-an-id"):
            with self.subTest(value=value), patch.dict(os.environ, {"SDK_REVIEW_RULESET_ID": value}), self.assertRaises(policy.ReleaseError):
                self.check()
        self.protection["required_pull_request_reviews"] = {"required_approving_review_count": 1}
        with self.assertRaisesRegex(policy.ReleaseError, "classic approval bypass"):
            self.check()

    def test_broad_or_missing_merger_bypass_stops_before_write(self):
        verified = self.check()
        for capability in (None, "never", "always", "exempt", "pull_request", "unknown"):
            self.ruleset["current_user_can_bypass"] = capability
            with self.subTest(capability=capability), patch.dict(os.environ, {"SDK_INTEGRATION_TOKEN": "write-fixture"}):
                with patch.object(integrate, "run", side_effect=AssertionError("Must not merge")), self.assertRaisesRegex(policy.ReleaseError, "PR-only approval bypass"):
                    integrate.merge(verified)

    def test_read_phase_does_not_assume_capability_but_write_token_must_prove_it(self):
        original_api = self.api
        def token_specific(path, **kwargs):
            response = original_api(path, **kwargs)
            if "/rulesets/" in path:
                response.pop("bypass_actors", None)
                response["current_user_can_bypass"] = "pull_requests_only" if kwargs.get("token") == "write-fixture" else "never"
            return response
        with patch.object(integrate, "api", side_effect=token_specific):
            verified = self.check()
            with patch.dict(os.environ, {"SDK_INTEGRATION_TOKEN": "write-fixture"}), patch.object(integrate, "run", return_value=json.dumps({"merged": True, "sha": "e" * 40}).encode()):
                integrate.merge(verified)
        self.assertEqual(self.ruleset_tokens, [None, "write-fixture"])
        self.assertEqual(os.environ["GH_TOKEN"], "read-only-fixture")

    def test_api_token_override_is_scoped_to_read_subprocess_environment(self):
        with patch.object(policy, "run", return_value=b'{}') as command:
            policy.api("repos/plaid/plaid-go/rulesets/12345", token="write-fixture")
        args, kwargs = command.call_args
        self.assertNotIn("write-fixture", args)
        self.assertEqual(kwargs["env"]["GH_TOKEN"], "write-fixture")
        self.assertEqual(os.environ["GH_TOKEN"], "read-only-fixture")

    def test_pending_failed_stale_or_wrong_author_ci_cannot_merge(self):
        original = copy.deepcopy(self.statuses[0])
        for change in ({"state": "pending"}, {"state": "failure"}, {"creator": {"id": 999}}, {"target_url": "https://attacker.example/success"}):
            self.statuses.insert(0, {**original, **change})
            with self.subTest(change=change), self.assertRaises(policy.ReleaseError):
                self.check()
            self.statuses.pop(0)
        self.combined["sha"] = "c" * 40
        with self.assertRaisesRegex(policy.ReleaseError, "another head"):
            self.check()

    def test_human_changes_request_remains_a_veto(self):
        self.reviews = [{"user": {"type": "User", "login": "reviewer"}, "author_association": "MEMBER", "state": "CHANGES_REQUESTED"}]
        with self.assertRaisesRegex(policy.ReleaseError, "Human changes-requested"):
            self.check()

    def test_disabled_or_untrusted_workflow_cannot_integrate(self):
        for change in ({"SDK_AUTOMERGE_ENABLED": "false"}, {"GITHUB_EVENT_NAME": "pull_request_target"},
                       {"GITHUB_REF_PROTECTED": "false"}, {"GITHUB_REPOSITORY": "attacker/plaid-go"},
                       {"SDK_PR_APP_BOT_ID": ""}, {"SDK_CIRCLECI_STATUS_ACTOR_ID": ""}):
            with self.subTest(change=change), patch.dict(os.environ, change), self.assertRaises(policy.ReleaseError):
                self.check()

    def test_merge_is_single_sha_pinned_squash_request_with_scoped_token(self):
        verified = self.check()
        with patch.dict(os.environ, {"SDK_INTEGRATION_TOKEN": "write-fixture"}), patch.object(integrate, "run", return_value=json.dumps({"merged": True, "sha": "e" * 40}).encode()) as write:
            self.assertEqual(integrate.merge(verified), "e" * 40)
        args, kwargs = write.call_args
        self.assertEqual(args[args.index("--method") + 1], "PUT")
        self.assertIn("repos/plaid/plaid-go/pulls/42/merge", args)
        self.assertEqual(json.loads(kwargs["input"])["sha"], self.head)
        self.assertEqual(json.loads(kwargs["input"])["merge_method"], "squash")
        self.assertEqual(kwargs["env"]["GH_TOKEN"], "write-fixture")
        self.assertEqual(os.environ["GH_TOKEN"], "read-only-fixture")

    def test_final_recheck_catches_head_base_or_review_race_without_write(self):
        verified = self.check()
        original = copy.deepcopy(self.pull)
        for mutation in (lambda: self.pull["head"].update(sha="c" * 40), lambda: self.branch["commit"].update(sha="d" * 40)):
            mutation()
            with patch.object(integrate, "run", side_effect=AssertionError("Must not write")), self.assertRaises(policy.ReleaseError):
                integrate.merge(verified)
            self.pull = copy.deepcopy(original)
            self.branch["commit"]["sha"] = self.base
        self.reviews = [{"user": {"type": "User", "login": "reviewer"}, "author_association": "OWNER", "state": "CHANGES_REQUESTED"}]
        with patch.object(integrate, "run", side_effect=AssertionError("Must not write")), self.assertRaises(policy.ReleaseError):
            integrate.merge(verified)

    def test_disabled_integration_stops_at_mutation_boundary(self):
        verified = self.check()
        with patch.dict(os.environ, {"SDK_AUTOMERGE_ENABLED": "false", "SDK_INTEGRATION_TOKEN": "write-fixture"}):
            with patch.object(integrate, "run", side_effect=AssertionError("Must not write")), self.assertRaisesRegex(policy.ReleaseError, "disabled"):
                integrate.merge(verified)

    def test_missing_or_shared_integration_identity_blocks_check_and_merge(self):
        verified = self.check()
        for identity in ("", "123", "not-an-id"):
            with self.subTest(identity=identity), patch.dict(os.environ, {"SDK_INTEGRATION_APP_BOT_ID": identity}):
                with self.assertRaises(policy.ReleaseError):
                    self.check()
                with patch.object(integrate, "run", side_effect=AssertionError("Must not write")), self.assertRaises(policy.ReleaseError):
                    integrate.merge(verified)

    def test_disabled_explicit_manual_dry_run_checks_but_never_merges(self):
        settings = {"SDK_AUTOMERGE_ENABLED": "false", "GITHUB_EVENT_NAME": "workflow_dispatch", "SDK_INTEGRATION_DRY_RUN": "true"}
        with patch.dict(os.environ, settings):
            verified = self.check()
            with patch.object(integrate, "run", side_effect=AssertionError("Must not write")), self.assertRaisesRegex(policy.ReleaseError, "disabled"):
                integrate.merge(verified)
            for change in ({"GITHUB_EVENT_NAME": "schedule"}, {"GITHUB_EVENT_NAME": "status"}, {"SDK_INTEGRATION_DRY_RUN": "false"}):
                with self.subTest(change=change), patch.dict(os.environ, change), self.assertRaises(policy.ReleaseError):
                    self.check()

    def test_enabled_manual_dry_run_still_cannot_merge(self):
        settings = {"GITHUB_EVENT_NAME": "workflow_dispatch", "SDK_INTEGRATION_DRY_RUN": "true", "SDK_INTEGRATION_TOKEN": "write-fixture"}
        with patch.dict(os.environ, settings):
            verified = self.check()
            with patch.object(integrate, "run", side_effect=AssertionError("Must not write")), self.assertRaisesRegex(policy.ReleaseError, "dry run cannot merge"):
                integrate.merge(verified)

    def test_unexpected_merged_tree_is_reported_as_failure(self):
        verified = self.check()
        original_api = self.api

        def different_tree(path, **kwargs):
            if "/git/commits/" in path:
                return {"parents": [{"sha": self.base}], "tree": {"sha": "f" * 40}}
            return original_api(path, **kwargs)

        with patch.dict(os.environ, {"SDK_INTEGRATION_TOKEN": "write-fixture"}), patch.object(integrate, "api", side_effect=different_tree), patch.object(integrate, "run", return_value=json.dumps({"merged": True, "sha": "e" * 40}).encode()):
            with self.assertRaisesRegex(policy.ReleaseError, "Merged result differs"):
                integrate.merge(verified)


if __name__ == "__main__":
    unittest.main()
