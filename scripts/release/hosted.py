"""Public-safe linkage from verified Fern output to Plaid's final SDK generation."""
import re
from pathlib import Path

import policy as attestation
from policy import canonical, require, sha256

GENERATED_GO_PATH = re.compile(r"(?:[a-z][a-z0-9_]*/)*[a-z][a-z0-9_]*\.go")
GENERATED_DATA_FILES = {"go.mod", "go.sum", ".fern/metadata.json", "reference.md", "wiremock/wiremock-mappings.json"}
# The generated index must not claim ownership of maintained SDK Go files.
PRESERVED_GO = {
    "internal/caller.go", "internal/destination.go", "internal/destination_test.go",
    "credential_destination_test.go", "core/custom_endpoint.go", "option/custom_endpoint.go",
    "retry_default_test.go", "wire_test_security_test.go", "transfer_diligence_upload_test.go",
}


def protected(path):
    return path in PRESERVED_GO

FERN_INDEX = ".sdk-release/fern-content.json"
FERN_BUNDLE = ".sdk-release/fern.bundle.json"
FERN_PREDICATE = "https://plaid.com/attestations/fern-generation/v1"
FERN_KEY_SHA256 = "f2751e378c0e9a6bb48bd8f16d90e0b44c456a86374066f7cf2477e8c7626b8b"


def validate_fern_record(record, manifest=None):
    keys, hex_value = attestation.keys, attestation.hex_value
    keys(record, "schema_version sdk_version source config sdk_base generator cli evidence content_sha256 generation", "Fern record")
    require(type(record["schema_version"]) is int and record["schema_version"] == 1, "Unsupported Fern record schema")
    attestation.version(record["sdk_version"])
    keys(record["source"], "repository commit path sha256", "Fern source")
    keys(record["config"], "repository commit sha256", "Fern config")
    require(record["source"]["repository"] == attestation.SPEC_REPO and
            record["source"]["path"] == "2020-09-14.yml" and
            record["config"]["repository"] == attestation.CONFIG_REPO, "Untrusted Fern source/config identity")
    for name in ("source", "config"):
        hex_value(record[name]["commit"], 40, "Fern source/config commit")
        hex_value(record[name]["sha256"], 64, "Fern source/config digest")
    require(record["source"]["commit"] == record["config"]["commit"], "Fern spec/config commits differ")
    hex_value(record["sdk_base"], 40, "Fern SDK base")
    keys(record["generator"], "image version sha256", "Fern generator")
    generator = record["generator"]
    require(generator["image"] == "fernapi/fern-go-sdk" and isinstance(generator["version"], str) and
            re.fullmatch(r"[0-9]+\.[0-9]+\.[0-9]+", generator["version"]), "Unexpected Fern generator")
    hex_value(generator["sha256"], 64, "Fern generator digest")
    keys(record["cli"], "version sha256", "Fern CLI")
    require(isinstance(record["cli"]["version"], str) and
            re.fullmatch(r"[0-9]+\.[0-9]+\.[0-9]+", record["cli"]["version"]), "Invalid Fern CLI version")
    hex_value(record["cli"]["sha256"], 64, "Fern CLI digest")
    keys(record["evidence"], "receipt_sha256 key_sha256 git_commit git_tree inputs_sha256 job_id_sha256", "Fern evidence")
    for name, digest in record["evidence"].items():
        hex_value(digest, 40 if name in {"git_commit", "git_tree"} else 64, "Fern evidence digest")
    require(record["evidence"]["key_sha256"] == FERN_KEY_SHA256, "Unexpected Fern signing key")
    hex_value(record["content_sha256"], 64, "Fern content digest")
    keys(record["generation"], "repository workflow run_id", "Fern verification workflow")
    require(record["generation"]["repository"] == attestation.CONFIG_REPO and
            record["generation"]["workflow"] == attestation.GEN_WORKFLOW, "Untrusted Fern verification workflow")
    attestation.decimal(record["generation"]["run_id"], "Fern verification run")
    if manifest is not None:
        require(record["sdk_version"] == manifest["sdk"]["version"], "Fern proof SDK version differs from final generation")
        require(all(record[name] == manifest[name] for name in ("source", "config", "generation")),
                "Fern proof source, config or run differs from final generation")
        require(record["cli"] == {"version": manifest["generator"]["cli_version"],
                                   "sha256": manifest["generator"]["cli_sha256"]} and
                "sha256:" + generator["sha256"] == manifest["generator"]["image_digest"],
                "Fern proof generator/CLI differs from final generation")
    return record


def validate_generation_mode(manifest, settings):
    require(isinstance(settings, dict), "Invalid reviewed generation settings")
    mode = settings.get("generation_mode", "local")
    require(isinstance(mode, str) and mode in {"local", "fern-hosted"}, "Unknown reviewed generation mode")
    require(manifest["schema_version"] == (2 if mode == "fern-hosted" else 1),
            "Generation evidence does not match reviewed generation mode")
    if mode == "fern-hosted":
        record = validate_fern_record(manifest["fern"], manifest)
        expected = {"image": settings.get("generator_image"), "version": settings.get("generator_version"),
                    "sha256": record["generator"]["sha256"]}
        require(record["generator"] == expected and settings.get("generator_digest") == "sha256:" + expected["sha256"] and record["cli"] ==
                {"version": settings.get("cli_version"), "sha256": settings.get("cli_sha256")},
                "Fern proof differs from reviewed generator settings")


def validate_fern_index(data, record):
    validate_fern_record(record)
    require(sha256(data) == record["content_sha256"], "Fern file index differs from signed record")
    index = attestation.parse_json(data)
    require(isinstance(index, list) and index, "Empty or invalid Fern file index")
    paths = []
    for item in index:
        require(isinstance(item, list) and len(item) == 3, "Invalid Fern file index entry")
        path, mode, digest = item
        require(isinstance(path, str) and (path in GENERATED_DATA_FILES or GENERATED_GO_PATH.fullmatch(path)) and
                path not in {"go.mod", "go.sum"} and not path.startswith(("scripts/", "legacy/")) and
                not protected(path), "Unapproved Fern file index path")
        require(mode == "100644", "Executable Fern file index entry")
        attestation.hex_value(digest, 64, "Fern file digest")
        paths.append(path)
    require(paths == sorted(set(paths)) and ".fern/metadata.json" in paths and
            any(GENERATED_GO_PATH.fullmatch(path) for path in paths),
            "Missing, duplicate or unordered Fern file index entries")
    require(canonical(index) == data, "Fern file index is not canonical")
    return index


def verify_fern_evidence(index: Path, bundle: Path, record: dict):
    validate_fern_index(index.read_bytes(), record)
    require(bundle.is_file() and not bundle.is_symlink(), "Missing Fern verification signature")
    attestation.verified_attestation(index, bundle, repository=attestation.CONFIG_REPO,
        workflow=attestation.GEN_WORKFLOW, ref="refs/heads/master", commit=record["config"]["commit"],
        predicate_type=FERN_PREDICATE, predicate=record,
        subjects=[{"name": "plaid-go-v" + record["sdk_version"] + "-fern-source",
                   "digest": {"sha256": record["content_sha256"]}}],
        run_id=record["generation"]["run_id"], private=False)
