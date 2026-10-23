"""Bind the canonical Go module archive to its reviewed commit and workflow."""

from pathlib import Path
import re
import zipfile

from policy import (GEN_BUNDLE, MANIFEST, RELEASE_PREDICATE, RELEASE_REF,
                    RELEASE_WORKFLOW, SDK_REPO, canonical, decimal, git,
                    hex_value, keys, parse_json, require, sha256,
                    validate_manifest, validate_tree, verified_attestation)


def make_release_manifest(sdk_repo, commit, control_commit, run_id, run_attempt, archive):
    manifest, files = validate_tree(sdk_repo, commit)
    hex_value(control_commit, 40, "control commit")
    decimal(run_id, "release run ID")
    decimal(run_attempt, "release run attempt")
    archive = Path(archive)
    expected_name = f"plaid-go-v{manifest['sdk']['version']}.zip"
    require(archive.name == expected_name, "Unexpected module archive filename")
    value = {
        "schema_version": 1,
        "sdk": manifest["sdk"],
        "commit": commit,
        "git_tree": git(sdk_repo, "rev-parse", f"{commit}^{{tree}}").decode().strip(),
        "content_sha256": manifest["content_sha256"],
        "generation_manifest_sha256": sha256(files[MANIFEST][1]),
        "generation_bundle_sha256": sha256(files[GEN_BUNDLE][1]),
        "module_zip": {"name": expected_name, "sha256": sha256(archive.read_bytes())},
        "workflow": {
            "repository": SDK_REPO,
            "path": RELEASE_WORKFLOW,
            "ref": RELEASE_REF,
            "commit": commit,
            "control_commit": control_commit,
            "run_id": run_id,
            "run_attempt": run_attempt,
        },
    }
    output = archive.parent
    (output / "generation-manifest.json").write_bytes(files[MANIFEST][1])
    (output / "generation.bundle.json").write_bytes(files[GEN_BUNDLE][1])
    (output / "release-manifest.json").write_bytes(canonical(value) + b"\n")
    manifest_digest = sha256((output / "release-manifest.json").read_bytes())
    (output / "subjects.sha256").write_text(f"{value['module_zip']['sha256']}  {archive.name}\n{manifest_digest}  release-manifest.json\n")
    return value


def artifact_files(directory, commit, control_commit, run_id, run_attempt):
    directory = Path(directory)
    for path in directory.rglob("*"):
        require(not path.is_symlink(), "Artifact contains a symlink")
        require(path.is_file(), "Artifact contains an unexpected directory")
    release_path = directory / "release-manifest.json"
    value = parse_json(release_path.read_bytes())
    keys(value, "schema_version sdk commit git_tree content_sha256 generation_manifest_sha256 generation_bundle_sha256 module_zip workflow", "release manifest")
    require(type(value["schema_version"]) is int and value["schema_version"] == 1, "Unknown release manifest schema")
    hex_value(commit, 40, "release commit")
    hex_value(control_commit, 40, "control commit")
    require(value["commit"] == commit, "Artifact is for another SDK commit")
    hex_value(value["git_tree"], 40, "release tree")
    for name in ("content_sha256", "generation_manifest_sha256", "generation_bundle_sha256"):
        hex_value(value[name], 64, name)
    expected_workflow = {"repository": SDK_REPO, "path": RELEASE_WORKFLOW,
                         "ref": RELEASE_REF, "commit": commit,
                         "control_commit": control_commit, "run_id": run_id,
                         "run_attempt": run_attempt}
    decimal(run_id, "release run ID")
    decimal(run_attempt, "release attempt")
    require(value["workflow"] == expected_workflow, "Artifact workflow identity mismatch")
    generation_path = directory / "generation-manifest.json"
    generation = validate_manifest(parse_json(generation_path.read_bytes()))
    require(value["sdk"] == generation["sdk"] and value["content_sha256"] == generation["content_sha256"], "Release does not match generation")
    require(sha256(generation_path.read_bytes()) == value["generation_manifest_sha256"], "Generation manifest changed")
    require(sha256((directory / "generation.bundle.json").read_bytes()) == value["generation_bundle_sha256"], "Generation bundle changed")
    keys(value["module_zip"], "name sha256", "module ZIP")
    filename = f"plaid-go-v{generation['sdk']['version']}.zip"
    require(value["module_zip"]["name"] == filename, "Unexpected ZIP filename")
    hex_value(value["module_zip"]["sha256"], 64, "ZIP digest")
    archive = directory / filename
    require(sha256(archive.read_bytes()) == value["module_zip"]["sha256"], "Module ZIP was modified")
    allowed = {filename, "release-manifest.json", "generation-manifest.json", "generation.bundle.json", "subjects.sha256", "release.bundle.json"}
    require({path.name for path in directory.iterdir()} <= allowed, "Unexpected release artifact")
    prefix = generation["sdk"]["module"] + "@v" + generation["sdk"]["version"] + "/"
    with zipfile.ZipFile(archive) as zipped:
        names = zipped.namelist()
        require(names and len(names) == len(set(names)), "Empty ZIP or duplicate ZIP paths")
        require(sum(info.file_size for info in zipped.infolist()) <= 500 * 1024 * 1024, "Module ZIP exceeds Go's size limit")
        for name in names:
            require(name.startswith(prefix), "Wrong Go module ZIP prefix")
            relative = name[len(prefix):]
            require(relative and not relative.startswith("/") and "\\" not in relative
                    and all(part not in ("", ".", "..") for part in relative.split("/")), "Unsafe module ZIP path")
    subjects = [{"name": filename, "digest": {"sha256": value["module_zip"]["sha256"]}},
                {"name": "release-manifest.json", "digest": {"sha256": sha256(release_path.read_bytes())}}]
    expected_checksums = "".join(f"{subject['digest']['sha256']}  {subject['name']}\n" for subject in subjects)
    require((directory / "subjects.sha256").read_text() == expected_checksums, "Attestation checksum list mismatch")
    return value, generation, subjects


def verify_release(directory, commit, control_commit, run_id, run_attempt):
    directory = Path(directory)
    release, generation, subjects = artifact_files(directory, commit, control_commit, run_id, run_attempt)
    bundle = directory / "release.bundle.json"
    require(bundle.is_file(), "Missing signed release attestation")
    for name in (release["module_zip"]["name"], "release-manifest.json"):
        verified_attestation(directory / name, bundle, repository=SDK_REPO,
                             workflow=RELEASE_WORKFLOW, ref=RELEASE_REF, commit=commit,
                             predicate_type=RELEASE_PREDICATE, predicate=release,
                             subjects=subjects, run_id=run_id, private=False)
    return release, generation
