#!/usr/bin/env python3
"""Exercise the exact production verifier against GitHub CLI's signed fixture.

This verifies existing official test data; it creates no attestation or release.
It does not prove that Plaid's production repositories or permissions are ready.
"""

import argparse
import base64
from pathlib import Path
import tempfile
import urllib.request

from policy import ReleaseError, canonical, parse_json, require, sha256, verified_attestation


SOURCE = "https://raw.githubusercontent.com/cli/cli/54a671d9a8f8c7a785c132d6f6a25375dc96acc1/pkg/cmd/attestation/test/data/"
ARTIFACT = "github_provenance_demo-0.0.12-py3-none-any.whl"
BUNDLE = "github_provenance_demo-0.0.12-py3-none-any-bundle.jsonl"
ARTIFACT_SHA = "ae57936def59bc4c75edd3a837d89bcefc6d3a5e31d55a6fa7a71624f92c3c3b"


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--fixture-dir", type=Path, required=True)
    parser.add_argument("--download", action="store_true", help="Download the two official fixture files if absent")
    args = parser.parse_args()
    if args.download:
        args.fixture_dir.mkdir(parents=True, exist_ok=True)
        for name in (ARTIFACT, BUNDLE):
            target = args.fixture_dir / name
            if not target.exists():
                target.write_bytes(urllib.request.urlopen(SOURCE + name, timeout=30).read())
    artifact, bundle = args.fixture_dir / ARTIFACT, args.fixture_dir / BUNDLE
    require(sha256(artifact.read_bytes()) == ARTIFACT_SHA, "Official fixture artifact differs")
    statement = parse_json(base64.b64decode(parse_json(bundle.read_bytes())["dsseEnvelope"]["payload"]))
    parameters = {
        "repository": "actions/attest-demo",
        "workflow": ".github/workflows/build-python.yml",
        "ref": "refs/heads/main",
        "commit": "a6c23b9806c593664f68637c8f9d45dfcf98b2db",
        "predicate_type": "https://slsa.dev/provenance/v1",
        "predicate": statement["predicate"],
        "subjects": [{"name": ARTIFACT, "digest": {"sha256": ARTIFACT_SHA}}],
        "run_id": "8788389601",
        "private": True,
    }
    verified = verified_attestation(artifact, bundle, **parameters)
    with tempfile.TemporaryDirectory(prefix="sdk-attestation-negative-") as temporary:
        changed = Path(temporary) / ARTIFACT
        changed.write_bytes(artifact.read_bytes() + b"tampered")
        try:
            verified_attestation(changed, bundle, **parameters)
        except ReleaseError:
            pass
        else:
            raise ReleaseError("Verifier accepted modified artifact bytes")
    print(canonical({"official_fixture_verified": True, "tampered_artifact_rejected": True,
                     "verified_records": len(verified), "production_private_oidc_proven": False}).decode())


if __name__ == "__main__":
    main()
