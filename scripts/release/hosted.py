"""Public-safe linkage from verified Fern output to Plaid's final SDK generation."""
import base64
import binascii
import json
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
FERN_PROVENANCE = ".sdk-release/fern-provenance.json"
FERN_PREDICATE = "https://plaid.com/attestations/fern-generation/v1"
FERN_KEY_SHA256 = "f2751e378c0e9a6bb48bd8f16d90e0b44c456a86374066f7cf2477e8c7626b8b"


def validate_fern_record(record, manifest=None):
    keys, hex_value = attestation.keys, attestation.hex_value
    keys(record, "schema_version sdk_version source config sdk_base generator cli evidence content_sha256 generation", "Fern record")
    require(type(record["schema_version"]) is int and record["schema_version"] == 2, "Unsupported Fern record schema")
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
    keys(record["evidence"], "receipt_sha256 key_sha256 output_mode git_tree inputs_sha256 job_id_sha256", "Fern evidence")
    require(record["evidence"]["output_mode"] == "downloadFiles", "Unexpected Fern output mode")
    for name, digest in record["evidence"].items():
        if name != "output_mode":
            hex_value(digest, 40 if name == "git_tree" else 64, "Fern evidence digest")
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


def validate_fern_provenance(data, record):
    """Bind the portable original receipt to Plaid's signed verification record.

    The producer verifies Fern's signature and complete downloaded tree. Consumers
    check its exact retained receipt and context, then verify Plaid's attestation.
    The unsigned vendor key identifier is deliberately absent from this envelope.
    """
    validate_fern_record(record)
    require(isinstance(data, bytes) and 0 < len(data) <= 1024 * 1024 and
            sha256(data) == record["evidence"]["receipt_sha256"], "Portable Fern receipt differs from signed record")
    try:
        keys = attestation.keys
        envelope = attestation.parse_json(data)
        keys(envelope, "payloadType payload signatures", "portable Fern envelope")
        require(json.dumps(envelope, sort_keys=True, separators=(",", ":"), ensure_ascii=True).encode() == data,
                "Portable Fern envelope is not canonical")
        require(envelope["payloadType"] == "application/vnd.in-toto+json" and
                isinstance(envelope["signatures"], list) and len(envelope["signatures"]) == 1,
                "Unexpected portable Fern signature envelope")
        keys(envelope["signatures"][0], "sig", "portable Fern signature")
        payload = base64.b64decode(envelope["payload"], validate=True)
        signature = base64.b64decode(envelope["signatures"][0]["sig"], validate=True)
        require(base64.b64encode(payload).decode() == envelope["payload"] and 0 < len(signature) <= 1024 and
                base64.b64encode(signature).decode() == envelope["signatures"][0]["sig"],
                "Portable Fern base64 is not canonical")
        statement = attestation.parse_json(payload)
        keys(statement, "_type subject predicateType predicate", "portable Fern statement")
        require(statement["_type"] == "https://in-toto.io/Statement/v1" and
                statement["predicateType"] == "https://slsa.dev/provenance/v1" and
                statement["subject"] == [{"name": "fern:generated-files", "digest": {"gitTree": record["evidence"]["git_tree"]}}],
                "Portable Fern statement or generated tree differs")
        predicate = statement["predicate"]
        keys(predicate, "buildDefinition runDetails", "portable Fern predicate")
        definition, details = predicate["buildDefinition"], predicate["runDetails"]
        keys(definition, "buildType externalParameters internalParameters resolvedDependencies", "portable Fern build")
        keys(details, "builder metadata", "portable Fern run")
        require(definition["buildType"] == "https://buildwithfern.com/provenance/sdk-generation/v1" and
                details["builder"] == {"id": "https://buildwithfern.com/fiddle"}, "Unexpected portable Fern builder")
        generator = record["generator"]
        require(definition["externalParameters"] == {"organization": "plaid", "api": "api",
                "generator": {key: generator[key] for key in ("image", "version")}, "output": {"mode": "downloadFiles"}},
                "Portable Fern output or generator differs")
        internal, metadata = definition["internalParameters"], details["metadata"]
        keys(internal, "jobId taskId", "portable Fern invocation")
        keys(metadata, "invocationId finishedOn", "portable Fern run metadata")
        uuid = r"[0-9a-f]{8}(?:-[0-9a-f]{4}){3}-[0-9a-f]{12}"
        require(all(isinstance(value, str) and re.fullmatch(uuid, value) for value in internal.values()) and
                sha256(internal["jobId"].encode()) == record["evidence"]["job_id_sha256"] and
                metadata["invocationId"] == internal["jobId"] + "/" + internal["taskId"] and
                isinstance(metadata["finishedOn"], str) and
                re.fullmatch(r"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d{1,9})?Z", metadata["finishedOn"]),
                "Portable Fern invocation differs")
        inputs, images = {}, []
        require(isinstance(definition["resolvedDependencies"], list), "Invalid portable Fern dependencies")
        for dependency in definition["resolvedDependencies"]:
            require(isinstance(dependency, dict) and set(dependency) in ({"name", "digest"}, {"uri", "digest"}),
                    "Unexpected portable Fern dependency")
            keys(dependency["digest"], "sha256", "portable Fern dependency digest")
            attestation.hex_value(dependency["digest"]["sha256"], 64, "portable Fern dependency digest")
            if "uri" in dependency:
                images.append(dependency)
            else:
                name = dependency["name"]
                require(isinstance(name, str) and name not in inputs, "Duplicate portable Fern input")
                inputs[name] = dependency["digest"]["sha256"]
        required = {"fern:../2020-09-14.yml", "fern:openapi-overlay.yml", "fern:generators.yml",
                    "fern:generators.yml#" + generator["image"]}
        require(images == [{"uri": "docker://docker.io/" + generator["image"] + ":" + generator["version"],
                            "digest": {"sha256": generator["sha256"]}}] and set(inputs) == required | {"fern:ir.json"} and
                inputs["fern:../2020-09-14.yml"] == record["source"]["sha256"] and
                sha256(canonical(dict(sorted((key, value) for key, value in inputs.items() if key != "fern:ir.json")))) ==
                record["evidence"]["inputs_sha256"], "Portable Fern inputs or image digest differ")
        return statement
    except (ValueError, TypeError, KeyError, IndexError, UnicodeError, binascii.Error, RecursionError):
        require(False, "Malformed portable Fern provenance")


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
