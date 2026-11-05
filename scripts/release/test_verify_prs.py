"""Required-status lifecycle; GitHub/signature boundaries are mocked explicitly."""
import copy
import json
import os
from pathlib import Path
import tempfile
import unittest
from unittest.mock import patch

import integrate
import policy
import verify_prs


class VerificationTests(unittest.TestCase):
    def setUp(self):
        self.head, self.base = "a" * 40, "b" * 40
        self.pull = {"number": 1, "state": "open", "draft": False, "merged": False,
                     "base": {"sha": self.base, "ref": "master", "repo": {"full_name": policy.SDK_REPO}},
                     "head": {"sha": self.head, "ref": "human-fix", "repo": {"full_name": policy.SDK_REPO}},
                     "user": {"type": "User", "id": 123}}
        self.reviews = [{"state": "APPROVED", "commit_id": self.head, "user": {"id": 456}}]
        self.decision = {"headRefOid": self.head, "baseRefName": "master", "reviewDecision": "APPROVED", "state": "OPEN", "isDraft": False}
        self.statuses = []
        self.generated = False
        self.patches = [
            patch.dict(os.environ, {"GITHUB_SHA": self.base, "SDK_PR_APP_BOT_ID": "999"}),
            patch.object(verify_prs, "api", side_effect=lambda *a, **kw: copy.deepcopy(self.pull)),
            patch.object(verify_prs, "report", side_effect=lambda head, state: self.statuses.append((head, state))),
            patch.object(verify_prs, "git", return_value=b"trusted workflow"),
            patch.object(integrate, "current_base", return_value=self.base),
            patch.object(integrate, "strict_status_protection"), patch.object(integrate, "approval_ruleset"),
            patch.object(integrate, "fetch_objects"), patch.object(integrate, "ci_and_reviews"),
            patch.object(policy, "human_reviews", side_effect=lambda _: self.reviews),
        ]
        for item in self.patches:
            item.start(); self.addCleanup(item.stop)
        real_run = verify_prs.run
        def run(*args, **kwargs):
            if args[:3] == ("gh", "api", "--hostname"):
                return json.dumps({"data": {"repository": {"pullRequest": self.decision}}}).encode()
            return real_run(*args, **kwargs)
        item = patch.object(verify_prs, "run", side_effect=run)
        item.start(); self.addCleanup(item.stop)

    def verify(self):
        verify_prs.verify_one(Path("trusted-control"), self.base, 1)

    def make_generated(self):
        self.pull["user"] = {"type": "Bot", "id": 999}
        self.pull["head"]["ref"] = "sdk-generation/go-v50-" + "a" * 20

    def test_human_path_requires_exact_head_approval_and_github_review_decision(self):
        self.verify()
        self.assertEqual(self.statuses, [(self.head, "pending"), (self.head, "success")])
        for review in ([], [{"state": "APPROVED", "commit_id": "c" * 40}]):
            self.statuses.clear(); self.reviews = review; self.verify()
            self.assertEqual(self.statuses, [(self.head, "pending")])
        self.reviews = [{"state": "APPROVED", "commit_id": self.head, "user": {"id": 456}}]
        for decision in (None, "REVIEW_REQUIRED", "CHANGES_REQUESTED"):
            self.statuses.clear(); self.decision["reviewDecision"] = decision; self.verify()
            self.assertEqual(self.statuses, [(self.head, "pending")])

    def test_generated_status_requires_existing_complete_verifier(self):
        self.make_generated()
        with patch.object(integrate, "check") as check:
            self.verify()
        self.assertEqual(check.call_args.kwargs, {"expected_base": self.base, "expected_head": self.head, "status_only": True})
        self.assertEqual(self.statuses[-1], (self.head, "success"))
        self.statuses.clear()
        with patch.object(integrate, "check", side_effect=policy.ReleaseError("bad signed output")), self.assertRaises(policy.ReleaseError):
            self.verify()
        self.assertEqual(self.statuses, [(self.head, "pending"), (self.head, "failure")])

    def test_pending_ci_clears_prior_success_without_self_status_deadlock(self):
        with patch.object(integrate, "ci_and_reviews", side_effect=integrate.NotReady("pending")):
            self.verify()
        self.assertEqual(self.statuses, [(self.head, "pending")])

    def test_stale_or_forged_human_review_response_never_succeeds(self):
        for mutation in ({"headRefOid": "c" * 40}, {"baseRefName": "other"}, {"state": "CLOSED"}, {"isDraft": True}):
            self.statuses.clear()
            with patch.dict(self.decision, mutation), self.assertRaises(policy.ReleaseError):
                self.verify()
            self.assertEqual(self.statuses[-1], (self.head, "failure"))

    def test_head_change_after_verification_marks_original_head_failed(self):
        moved = copy.deepcopy(self.pull); moved["head"]["sha"] = "c" * 40
        with patch.object(verify_prs, "api", side_effect=[copy.deepcopy(self.pull), moved]), self.assertRaisesRegex(policy.ReleaseError, "changed before"):
            self.verify()
        self.assertEqual(self.statuses, [(self.head, "pending"), (self.head, "failure")])

    def test_status_credentials_are_only_subprocess_environment(self):
        # Restore the real reporter while keeping its subprocess boundary mocked.
        self.patches[2].stop()
        with patch.dict(os.environ, {"SDK_VERIFICATION_STATUS_TOKEN": "status-fixture", "GITHUB_RUN_ID": "123"}), patch.object(verify_prs, "run", return_value=b"{}") as command:
            verify_prs.report(self.head, "pending")
        args, kwargs = command.call_args
        self.assertNotIn("status-fixture", args)
        self.assertNotIn(b"status-fixture", kwargs["input"])
        self.assertEqual(kwargs["env"]["GH_TOKEN"], "status-fixture")
        self.assertEqual(json.loads(kwargs["input"])["context"], integrate.VERIFIED_CONTEXT)

    def test_pr_author_approval_does_not_qualify(self):
        self.reviews = [{"state": "APPROVED", "commit_id": self.head, "user": {"id": 123}}]
        self.verify()
        self.assertEqual(self.statuses, [(self.head, "pending")])

    def test_manual_verification_continues_when_automerge_is_disabled(self):
        with patch.dict(os.environ, {"SDK_AUTOMERGE_ENABLED": "false"}), patch.object(integrate, "merge", side_effect=AssertionError("Must not merge")):
            self.verify()
        self.assertEqual(self.statuses[-1], (self.head, "success"))

    def test_changes_requested_or_dismissed_review_invalidates_success(self):
        with patch.object(policy, "human_reviews", side_effect=policy.ReleaseError("changes-requested")), self.assertRaises(policy.ReleaseError):
            self.verify()
        self.assertEqual(self.statuses[-1], (self.head, "failure"))
        self.statuses.clear()
        self.reviews = [{"state": "DISMISSED", "commit_id": self.head}]
        self.verify()
        self.assertEqual(self.statuses, [(self.head, "pending")])

    def test_status_write_failure_is_not_silently_treated_as_pr_failure(self):
        with patch('sys.argv', ['verify_prs', '--control', 'control', '--control-commit', self.base]), patch.object(integrate, 'context'), patch.object(verify_prs, 'api', return_value=[{'number': 1}]), patch.object(verify_prs, 'verify_one', side_effect=verify_prs.StatusWriteError('write failed')):
            with self.assertRaises(verify_prs.StatusWriteError):
                verify_prs.main()

    def test_review_signal_has_no_checkout_credentials_or_artifact_consumption(self):
        root = Path(__file__).resolve().parents[2]
        signal = (root / '.github/workflows/sdk-review-signal.yml').read_text()
        self.assertIn('permissions: {}', signal)
        self.assertIn('[submitted, dismissed, edited]', signal)
        for forbidden in ('uses:', 'secrets.', 'checkout', 'download-artifact'):
            self.assertNotIn(forbidden, signal)
        workflow = (root / '.github/workflows/integrate.yml').read_text()
        self.assertEqual(workflow.count('statuses: write'), 1)
        self.assertIn("github.event.context != 'sdk-generation/verified'", workflow)
        self.assertIn('workflows: [SDK review signal]', workflow)

    def test_dry_run_cannot_write_statuses(self):
        with patch('sys.argv', ['verify_prs', '--control', 'control', '--control-commit', self.base]), patch.object(integrate, 'context'), patch.dict(os.environ, {"SDK_INTEGRATION_DRY_RUN": "true"}):
            verify_prs.main()
        self.assertEqual(self.statuses, [])


if __name__ == '__main__':
    unittest.main()
