#!/usr/bin/env python3
"""Trusted entry points for the disabled-by-default automated Go SDK release."""

import argparse
import os
from pathlib import Path
import sys

from artifacts import artifact_files, make_release_manifest, verify_release
from policy import (GEN_BUNDLE, MANIFEST, ReleaseError, canonical, git,
                    require, require_release_merge, require_context,
                    require_control_workflow, run, sha256, validate_tree,
                    verify_generation, verify_public_spec)
from publish import publish


def output(values):
    if os.environ.get("GITHUB_OUTPUT"):
        with open(os.environ["GITHUB_OUTPUT"], "a") as stream:
            for name, value in values.items():
                require("\n" not in str(value), "Multiline workflow output")
                stream.write(f"{name}={value}\n")
    print(canonical(values).decode())


def context(args):
    require_context(os.environ, args.commit, args.control_commit)


def exact_checkout(repo, commit):
    require(git(repo, "rev-parse", "HEAD").decode().strip() == commit, "Wrong SDK checkout")
    require(not git(repo, "status", "--porcelain", "--untracked-files=all").strip()
            and not git(repo, "ls-files", "--others").strip(), "SDK checkout must be completely clean")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("operation", choices=("gate", "inspect", "package", "check-artifacts", "verify-release", "publish"))
    parser.add_argument("--commit", required=True)
    parser.add_argument("--control-commit", required=True)
    parser.add_argument("--sdk", type=Path)
    parser.add_argument("--control", type=Path)
    parser.add_argument("--output", type=Path)
    parser.add_argument("--zip-tool", type=Path)
    parser.add_argument("--publish", action="store_true")
    args = parser.parse_args()
    context(args)
    if args.operation == "gate":
        require(args.sdk is not None and args.control is not None, "SDK and control checkouts are required")
        tracked = git(args.sdk, "ls-tree", "--name-only", args.commit, MANIFEST).strip()
        if not tracked:
            output({"ready": "false", "reason": "No generation manifest; release remains inactive"})
            return
        require_control_workflow(args.sdk, args.commit, args.control, args.control_commit)
        pr = require_release_merge(args.commit)
        manifest, files = verify_generation(args.sdk, args.commit)
        verify_public_spec(manifest)
        output({"ready": "true", "version": manifest["sdk"]["version"],
                "content_sha256": manifest["content_sha256"],
                "manifest_sha256": sha256(files[MANIFEST][1]), "pull_request": str(pr)})
    elif args.operation in ("inspect", "package"):
        require(args.sdk is not None, "SDK checkout is required")
        exact_checkout(args.sdk, args.commit)
        manifest, files = validate_tree(args.sdk, args.commit)
        require(manifest["content_sha256"] == os.environ.get("VERIFIED_CONTENT_SHA256")
                and sha256(files[MANIFEST][1]) == os.environ.get("VERIFIED_MANIFEST_SHA256"), "SDK differs from verified generation job")
        if args.operation == "inspect":
            output({"version": manifest["sdk"]["version"], "module": manifest["sdk"]["module"]})
            return
        require(args.output is not None and args.zip_tool is not None, "Output directory and canonical ZIP tool are required")
        args.output.mkdir(parents=True, exist_ok=False)
        archive = args.output / f"plaid-go-v{manifest['sdk']['version']}.zip"
        run(args.zip_tool.resolve(), args.sdk.resolve(), manifest["sdk"]["module"], "v" + manifest["sdk"]["version"], archive.resolve())
        make_release_manifest(args.sdk, args.commit, args.control_commit, os.environ["GITHUB_RUN_ID"], os.environ["GITHUB_RUN_ATTEMPT"], archive)
        output({"zip": archive.name})
    else:
        require(args.output is not None, "Release artifact directory is required")
        parameters = (args.output, args.commit, args.control_commit, os.environ["GITHUB_RUN_ID"], os.environ["GITHUB_RUN_ATTEMPT"])
        if args.operation == "check-artifacts":
            artifact_files(*parameters)
            output({"checked": "true"})
        elif args.operation == "verify-release":
            verify_release(*parameters)
            output({"verified": "true"})
        else:
            state = publish(*parameters, os.environ.get("SDK_RELEASE_SIGNING_PUBLIC_KEY", ""), publish_requested=args.publish)
            output({"state": state})


if __name__ == "__main__":
    try:
        main()
    except (ReleaseError, OSError, ValueError, KeyError) as error:
        sys.exit(str(error))
