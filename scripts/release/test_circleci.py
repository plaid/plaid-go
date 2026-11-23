"""Verify CircleCI job evidence at the HTTPS boundary."""

import json
import unittest
from unittest.mock import patch

import circleci
from policy import ReleaseError


class CircleCITests(unittest.TestCase):
    def setUp(self):
        self.branch = "sdk-generation/candidate+1"
        self.head = "a" * 40
        self.job = {
            "vcs_url": "https://github.com/plaid/plaid-go",
            "branch": self.branch,
            "vcs_revision": self.head,
            "workflows": {"workflow_name": "build", "job_name": "build"},
            "why": "github",
            "vcs_tag": None,
            "build_parameters": {},
            "ssh_users": [],
            "lifecycle": "finished",
            "status": "success",
            "outcome": "success",
        }
        self.https = self.enterContext(patch.object(circleci, "HTTPSConnection"))
        self.connection = self.https.return_value
        self.response = self.connection.getresponse.return_value
        self.response.status = 200
        self.response.read.return_value = json.dumps([self.job]).encode()

    def verify(self):
        circleci.verify_build(self.branch, self.head)

    def test_reads_latest_job_for_encoded_branch_from_fixed_https_origin(self):
        self.verify()
        self.https.assert_called_once_with("circleci.com", timeout=10)
        self.connection.request.assert_called_once_with(
            "GET", "/api/v1.1/project/github/plaid/plaid-go/tree/sdk-generation%2Fcandidate%2B1?limit=1",
            headers={"Accept": "application/json", "Cache-Control": "no-cache"},
        )
        self.connection.close.assert_called_once_with()

    def test_rejects_jobs_without_matching_successful_build_evidence(self):
        cases = [
            ({"branch": "master"}, "does not match"),
            ({"vcs_revision": "b" * 40}, "does not match"),
            ({"vcs_url": "https://github.com/plaid/another-repo"}, "does not match"),
            ({"workflows": {"workflow_name": "release", "job_name": "build"}}, "not the required build"),
            ({"workflows": {"workflow_name": "build", "job_name": "lint"}}, "not the required build"),
            ({"why": "retry"}, "normal VCS-triggered"),
            ({"vcs_tag": "v1.0.0"}, "normal VCS-triggered"),
            ({"build_parameters": {"mode": "preview"}}, "normal VCS-triggered"),
            ({"ssh_users": ["fixture"]}, "normal VCS-triggered"),
            ({"lifecycle": "running"}, "has not succeeded"),
            ({"status": "failed"}, "has not succeeded"),
            ({"outcome": "canceled"}, "has not succeeded"),
        ]
        for fields, error in cases:
            with self.subTest(fields=fields):
                self.response.read.return_value = json.dumps([self.job | fields]).encode()
                with self.assertRaisesRegex(ReleaseError, error):
                    self.verify()

    def test_requires_exactly_one_latest_job(self):
        for jobs in ([], {}, [None], [self.job | {"status": "failed"}, self.job]):
            with self.subTest(jobs=jobs):
                self.response.read.return_value = json.dumps(jobs).encode()
                with self.assertRaisesRegex(ReleaseError, "no confirmed latest job"):
                    self.verify()

    def test_rejects_http_and_payload_errors(self):
        for status, payload, error in (
            (302, b"", "HTTP 302"),
            (503, b"", "HTTP 503"),
            (200, b"invalid json", "Invalid JSON"),
            (200, b" " * (1024 * 1024 + 1), "too large"),
        ):
            with self.subTest(status=status, error=error):
                self.response.status = status
                self.response.read.return_value = payload
                with self.assertRaisesRegex(ReleaseError, error):
                    self.verify()
        self.response.read.assert_called_with(1024 * 1024 + 1)

    def test_transport_failure_closes_connection(self):
        self.connection.getresponse.side_effect = TimeoutError()
        with self.assertRaisesRegex(ReleaseError, "Could not read"):
            self.verify()
        self.connection.close.assert_called_once_with()
