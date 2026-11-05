"""Public Fern handoff linkage; only the gh verification process is mocked."""
import copy
import json
from pathlib import Path
import tempfile
import unittest
from unittest.mock import patch

import policy as attestation
import policy as common
import policy
import integrate
import hosted


def record_fixture(manifest=None, *, sdk_base="b" * 40, index=None):
    """Build a public-safe synthetic record; callers may supply their manifest/index."""
    index = index if index is not None else common.canonical([
        [".fern/metadata.json", "100644", "1" * 64], ["client/client.go", "100644", "2" * 64]])
    record = {"schema_version": 1, "sdk_version": "50.0.0",
        "source": {"repository": attestation.SPEC_REPO, "commit": "a" * 40, "path": "2020-09-14.yml", "sha256": "c" * 64},
        "config": {"repository": attestation.CONFIG_REPO, "commit": "a" * 40, "sha256": "d" * 64},
        "sdk_base": sdk_base, "generator": {"image": "fernapi/fern-go-sdk", "version": "1.70.3", "sha256": "e" * 64},
        "cli": {"version": "5.193.2", "sha256": "f" * 64},
        "evidence": {"receipt_sha256": "1" * 64, "key_sha256": hosted.FERN_KEY_SHA256,
                     "git_commit": "2" * 40, "git_tree": "3" * 40, "inputs_sha256": "4" * 64, "job_id_sha256": "5" * 64},
        "content_sha256": common.sha256(index),
        "generation": {"repository": attestation.CONFIG_REPO, "workflow": attestation.GEN_WORKFLOW, "run_id": "123"}}
    if manifest is not None:
        record.update(sdk_version=manifest["sdk"]["version"],
                      **{name: copy.deepcopy(manifest[name]) for name in ("source", "config", "generation")})
        record["cli"] = {"version": manifest["generator"]["cli_version"], "sha256": manifest["generator"]["cli_sha256"]}
        record["generator"]["sha256"] = manifest["generator"]["image_digest"].removeprefix("sha256:")
    return record


def manifest_fixture(record=None):
    record = record or record_fixture()
    return {"schema_version": 2, "sdk": {"repository": attestation.SDK_REPO, "version": record["sdk_version"],
            "module": "github.com/plaid/plaid-go/v" + record["sdk_version"].split(".")[0]},
        **{name: copy.deepcopy(record[name]) for name in ("source", "config", "generation")},
        "generator": {"cli_version": record["cli"]["version"], "cli_sha256": record["cli"]["sha256"],
                      "image_digest": "sha256:" + record["generator"]["sha256"]},
        "legacy": {"source_commit": "a" * 40, "spec_sha256": record["source"]["sha256"],
                   "tree_sha256": "b" * 64, "generator_digest": "sha256:" + "c" * 64},
        "content_sha256": "d" * 64, "fern": record}


def verification_fixture(record):
    identity = f"https://github.com/{attestation.CONFIG_REPO}/{attestation.GEN_WORKFLOW}@refs/heads/master"
    certificate = {"subjectAlternativeName": identity, "issuer": "https://token.actions.githubusercontent.com",
        "buildSignerURI": identity, "buildSignerDigest": record["config"]["commit"],
        "sourceRepositoryURI": "https://github.com/" + attestation.CONFIG_REPO,
        "sourceRepositoryDigest": record["config"]["commit"], "sourceRepositoryRef": "refs/heads/master",
        "runnerEnvironment": "github-hosted", "sourceRepositoryVisibilityAtSigning": "public",
        "runInvocationURI": f"https://github.com/{attestation.CONFIG_REPO}/actions/runs/{record['generation']['run_id']}"}
    return [{"verificationResult": {"signature": {"certificate": certificate}, "verifiedTimestamps": [{"type": "fixture"}],
        "statement": {"_type": "https://in-toto.io/Statement/v1", "predicateType": hosted.FERN_PREDICATE,
                      "predicate": copy.deepcopy(record), "subject": [{"name": "plaid-go-v" + record["sdk_version"] + "-fern-source",
                      "digest": {"sha256": record["content_sha256"]}}]}}}]


class HostedContractTests(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory()
        self.addCleanup(self.temporary.cleanup)
        self.root = Path(self.temporary.name)
        self.index = common.canonical([[".fern/metadata.json", "100644", "1" * 64], ["client/client.go", "100644", "2" * 64]])
        self.record = record_fixture(index=self.index)
        self.manifest = manifest_fixture(self.record)
        self.index_path, self.bundle = self.root / "index.json", self.root / "bundle.json"
        self.index_path.write_bytes(self.index)
        self.bundle.write_text('{"fixture": "not a signature"}')

    def test_schema2_links_every_input(self):
        self.assertEqual(attestation.validate_manifest(self.manifest), self.manifest)
        changes = [("sdk_version", "50.1.0"), ("source", {**self.record["source"], "sha256": "0" * 64}),
                   ("config", {**self.record["config"], "sha256": "0" * 64}),
                   ("generation", {**self.record["generation"], "run_id": "999"}),
                   ("cli", {**self.record["cli"], "version": "9.0.0"}),
                   ("generator", {**self.record["generator"], "sha256": "0" * 64})]
        for field, value in changes:
            manifest = copy.deepcopy(self.manifest)
            manifest["fern"][field] = value
            with self.subTest(field=field), self.assertRaises(common.ReleaseError):
                attestation.validate_manifest(manifest)

    def test_schema1_historical_fallback_and_downgrade_guard(self):
        local = copy.deepcopy(self.manifest)
        local["schema_version"] = 1
        del local["fern"]
        attestation.validate_manifest(local)
        hosted.validate_generation_mode(local, {})
        hosted.validate_generation_mode(local, {"generation_mode": "local"})
        with self.assertRaisesRegex(common.ReleaseError, "reviewed generation mode"):
            hosted.validate_generation_mode(local, {"generation_mode": "fern-hosted"})
        with self.assertRaises(common.ReleaseError):
            hosted.validate_generation_mode(self.manifest, {})

    def test_hosted_mode_checks_reviewed_settings(self):
        generator = self.record["generator"]
        settings = {"generation_mode": "fern-hosted", "generator_image": generator["image"],
                    "generator_version": generator["version"], "generator_digest": "sha256:" + generator["sha256"],
                    "cli_version": self.record["cli"]["version"], "cli_sha256": self.record["cli"]["sha256"]}
        hosted.validate_generation_mode(self.manifest, settings)
        for field in settings:
            with self.subTest(field=field), self.assertRaises(common.ReleaseError):
                hosted.validate_generation_mode(self.manifest, {**settings, field: "unexpected"})

    def test_missing_or_private_record_fields_and_wrong_key_rejected(self):
        for field in self.record:
            record = copy.deepcopy(self.record)
            del record[field]
            with self.subTest(field=field), self.assertRaises(common.ReleaseError):
                hosted.validate_fern_record(record)
        for extra in ("repository_url", "branch", "receipt"):
            with self.subTest(extra=extra), self.assertRaises(common.ReleaseError):
                hosted.validate_fern_record({**self.record, extra: "not public"})
        record = copy.deepcopy(self.record)
        record["evidence"]["key_sha256"] = "0" * 64
        with self.assertRaisesRegex(common.ReleaseError, "signing key"):
            hosted.validate_fern_record(record)

    def test_index_requires_exact_hash_canonical_paths_and_modes(self):
        hosted.validate_fern_index(self.index, self.record)
        with self.assertRaisesRegex(common.ReleaseError, "index differs"):
            hosted.validate_fern_index(self.index + b"\n", self.record)
        original = json.loads(self.index)
        cases = [original + [original[0]], list(reversed(original)), [original[0]],
                 [[".fern/metadata.json", "100755", "1" * 64], original[1]]]
        for path in ("go.mod", "go.sum", "scripts/run.go", "legacy/client.go", "internal/caller.go",
                     "wiremock/docker-compose.test.yml", "../secret.go", ".sdk-release/extra.json"):
            cases.append(sorted(original + [[path, "100644", "3" * 64]]))
        for index in cases:
            data = common.canonical(index)
            record = {**self.record, "content_sha256": common.sha256(data)}
            with self.subTest(index=index), self.assertRaises(common.ReleaseError):
                hosted.validate_fern_index(data, record)

    def test_evidence_requires_exact_gha_signature_identity_and_subject(self):
        result = verification_fixture(self.record)
        with patch.object(attestation, "run", return_value=common.canonical(result)) as command:
            hosted.verify_fern_evidence(self.index_path, self.bundle, self.record)
        self.assertIn("--source-digest", command.call_args.args)
        self.assertIn(hosted.FERN_PREDICATE, command.call_args.args)
        self.assertNotIn("--no-public-good", command.call_args.args)
        for field, replacement in (("runInvocationURI", "https://github.com/plaid/plaid-openapi/actions/runs/999"),
                                   ("sourceRepositoryDigest", "0" * 40), ("sourceRepositoryVisibilityAtSigning", "private")):
            bad = copy.deepcopy(result)
            bad[0]["verificationResult"]["signature"]["certificate"][field] = replacement
            with self.subTest(field=field), patch.object(attestation, "run", return_value=common.canonical(bad)), self.assertRaises(common.ReleaseError):
                hosted.verify_fern_evidence(self.index_path, self.bundle, self.record)
        bad = copy.deepcopy(result)
        bad[0]["verificationResult"]["statement"]["subject"][0]["digest"]["sha256"] = "0" * 64
        with patch.object(attestation, "run", return_value=common.canonical(bad)), self.assertRaises(common.ReleaseError):
            hosted.verify_fern_evidence(self.index_path, self.bundle, self.record)
        with patch.object(attestation, "run", side_effect=common.ReleaseError("Signature rejected")), self.assertRaises(common.ReleaseError):
            hosted.verify_fern_evidence(self.index_path, self.bundle, self.record)
        self.bundle.unlink()
        with self.assertRaisesRegex(common.ReleaseError, "signature"):
            hosted.verify_fern_evidence(self.index_path, self.bundle, self.record)


class HostedGenerationTests(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory()
        self.addCleanup(self.temporary.cleanup)
        self.repo = Path(self.temporary.name) / "sdk"
        self.repo.mkdir()
        policy.git(self.repo, "init", "-q")
        policy.git(self.repo, "config", "user.name", "Fixture")
        policy.git(self.repo, "config", "user.email", "fixture@example.invalid")
        policy.git(self.repo, "config", "commit.gpgsign", "false")
        policy.git(self.repo, "config", "core.hooksPath", "/dev/null")
        self.write("go.mod", b"module github.com/plaid/plaid-go/v50\n\ngo 1.21\n")
        self.write("go.sum", b"")
        self.write("legacy/client.go", b"package plaid\n")
        self.write("legacy/configuration.go", b"package plaid\n")
        self.write("client/client.go", b"package client\n")
        self.write(".fern/metadata.json", policy.canonical({"cliVersion": "5.193.2", "generatorName": "fernapi/fern-go-sdk",
                                                          "generatorVersion": "1.70.3", "sdkVersion": "v50.0.0"}))
        self.base = self.commit()
        files = policy.tree_files(self.repo, self.base)
        self.index = policy.canonical([[path, mode, policy.sha256(data)] for path, (mode, data) in sorted(files.items())
                                       if path in {".fern/metadata.json", "client/client.go"}])
        self.record = record_fixture(sdk_base=self.base, index=self.index)
        self.manifest = manifest_fixture(self.record)
        self.manifest["content_sha256"] = policy.sha256(policy.content_bytes(files))
        self.manifest["legacy"]["tree_sha256"] = policy.sha256(policy.content_bytes(files, legacy_only=True))
        self.fern_bundle, self.gen_bundle = b'{"fixture":"fern"}', b'{"fixture":"generation"}'
        self.write(policy.FERN_INDEX, self.index)
        self.write(policy.FERN_BUNDLE, self.fern_bundle)
        self.write(policy.GEN_BUNDLE, self.gen_bundle)
        self.write(policy.MANIFEST, policy.canonical(self.manifest))
        self.sha = self.commit()
        self.settings = {"generation_mode": "fern-hosted", "generator_image": self.record["generator"]["image"],
                         "generator_version": self.record["generator"]["version"],
                         "generator_digest": "sha256:" + self.record["generator"]["sha256"],
                         "cli_version": self.record["cli"]["version"], "cli_sha256": self.record["cli"]["sha256"]}
        self.calls = []

    def write(self, path, content):
        target = self.repo / path
        target.parent.mkdir(parents=True, exist_ok=True)
        target.write_bytes(content)

    def commit(self):
        policy.git(self.repo, "add", "-A")
        policy.git(self.repo, "commit", "-qm", "Fixture")
        return policy.git(self.repo, "rev-parse", "HEAD").decode().strip()

    def rewrite_manifest(self):
        self.write(policy.MANIFEST, policy.canonical(self.manifest))
        self.sha = self.commit()

    def verifier(self, *arguments, **kwargs):
        if arguments[:3] != ("gh", "attestation", "verify"):
            return self.real_run(*arguments, **kwargs)
        predicate_type = arguments[arguments.index("--predicate-type") + 1]
        self.calls.append(predicate_type)
        bundle = Path(arguments[arguments.index("--bundle") + 1]).read_bytes()
        expected_bundle = self.fern_bundle if predicate_type == hosted.FERN_PREDICATE else self.gen_bundle
        policy.require(bundle == expected_bundle, "Fixture cryptographic verifier rejected swapped signature")
        result = verification_fixture(self.record)
        if predicate_type == policy.GEN_PREDICATE:
            statement = result[0]["verificationResult"]["statement"]
            statement.update(predicateType=predicate_type, predicate=copy.deepcopy(self.manifest),
                             subject=[{"name": "plaid-go-v50.0.0-source", "digest": {"sha256": self.manifest["content_sha256"]}}])
        return policy.canonical(result)

    def api(self, path, **kwargs):
        self.assertEqual(path, f"repos/{policy.CONFIG_REPO}/contents/pipeline/generation/settings.json?ref={self.manifest['config']['commit']}")
        self.assertEqual(kwargs, {"raw": True})
        return policy.canonical(self.settings)

    def verify(self):
        self.real_run = policy.run
        with patch.object(policy, "run", side_effect=self.verifier), patch.object(policy, "api", side_effect=self.api):
            return policy.verify_generation(self.repo, self.sha)

    def test_hosted_generation_checks_both_attestations_and_config_mode(self):
        manifest, files = self.verify()
        self.assertEqual(manifest, self.manifest)
        self.assertEqual(self.calls, [hosted.FERN_PREDICATE, policy.GEN_PREDICATE])
        self.assertEqual(policy.sha256(policy.content_bytes(files)), manifest["content_sha256"])
        changed = {**files, policy.FERN_BUNDLE: ("100644", b"different valid encoding")}
        self.assertEqual(policy.content_bytes(files), policy.content_bytes(changed))

    def test_missing_or_swapped_proof_stops_generation(self):
        for name in (policy.FERN_INDEX, policy.FERN_BUNDLE):
            files = policy.tree_files(self.repo, self.sha)
            del files[name]
            with self.subTest(name=name), patch.object(policy, "tree_files", return_value=files), self.assertRaises(policy.ReleaseError):
                policy.validate_tree(self.repo, self.sha)
        files = policy.tree_files(self.repo, self.sha)
        files[policy.FERN_BUNDLE] = ("100644", b"swapped signature")
        with patch.object(policy, "tree_files", return_value=files), self.assertRaisesRegex(policy.ReleaseError, "swapped signature"):
            self.verify()

    def test_changed_index_and_wrong_generator_metadata_rejected(self):
        files = policy.tree_files(self.repo, self.sha)
        for name, replacement in ((policy.FERN_INDEX, self.index + b"\n"),
                                  (".fern/metadata.json", policy.canonical({"cliVersion": "5.193.2", "generatorName": "fernapi/fern-go-sdk",
                                                                        "generatorVersion": "9.0.0", "sdkVersion": "v50.0.0"}))):
            changed = {**files, name: ("100644", replacement)}
            if name == ".fern/metadata.json":
                manifest = copy.deepcopy(self.manifest)
                manifest["content_sha256"] = policy.sha256(policy.content_bytes(changed))
                changed[policy.MANIFEST] = ("100644", policy.canonical(manifest))
            with self.subTest(name=name), patch.object(policy, "tree_files", return_value=changed), self.assertRaises(policy.ReleaseError):
                policy.validate_tree(self.repo, self.sha)

    def test_schema1_cannot_carry_unverified_hosted_files(self):
        self.manifest["schema_version"] = 1
        del self.manifest["fern"]
        self.rewrite_manifest()
        with self.assertRaisesRegex(policy.ReleaseError, "unverified Fern proof"):
            policy.validate_tree(self.repo, self.sha)

    def test_hosted_settings_reject_schema1_downgrade_but_historical_local_works(self):
        self.manifest["schema_version"] = 1
        del self.manifest["fern"]
        (self.repo / policy.FERN_INDEX).unlink()
        (self.repo / policy.FERN_BUNDLE).unlink()
        self.rewrite_manifest()
        with self.assertRaisesRegex(policy.ReleaseError, "reviewed generation mode"):
            self.verify()
        self.assertEqual(self.calls, [])
        self.settings = {}
        self.verify()
        self.assertEqual(self.calls, [policy.GEN_PREDICATE])

    def test_wrong_request_base_stops_generation(self):
        self.record["sdk_base"] = "0" * 40
        self.rewrite_manifest()
        with self.assertRaisesRegex(policy.ReleaseError, "SDK base"):
            self.verify()

    def test_evidence_allowlist_is_exact_and_non_executable(self):
        self.assertEqual(policy.EVIDENCE_FILES, (policy.MANIFEST, policy.GEN_BUNDLE, policy.FERN_INDEX, policy.FERN_BUNDLE))
        integrate.approved_generated_changes({}, {name: ("100644", b"proof") for name in policy.EVIDENCE_FILES})
        with self.assertRaises(policy.ReleaseError):
            integrate.approved_generated_changes({}, {".sdk-release/unverified.json": ("100644", b"proof")})
        self.write(policy.FERN_INDEX, self.index)
        (self.repo / policy.FERN_INDEX).chmod(0o755)
        with self.assertRaisesRegex(policy.ReleaseError, "executable"):
            policy.tree_files(self.repo, self.commit())

    def test_raw_index_maintained_paths_match_integration_policy(self):
        expected = {name for name in integrate.PRESERVED_FILES if name.endswith(".go")}
        self.assertEqual(hosted.PRESERVED_GO, expected)


if __name__ == "__main__":
    unittest.main()
