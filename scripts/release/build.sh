#!/usr/bin/env bash
# Called only in the read-only build job with an allowlisted environment.
set -euo pipefail
[[ $# == 1 ]] || { echo 'usage: build.sh <sdk-checkout>' >&2; exit 2; }
SDK="$(cd "$1" && pwd -P)"
TEST_DIR="$(mktemp -d "${TMPDIR:-/tmp}/sdk-build.XXXXXXXX")"
WIREMOCK_PID=""
cleanup() {
  local status=$?
  if [[ -n "$WIREMOCK_PID" ]]; then kill "$WIREMOCK_PID" 2>/dev/null || true; fi
  if [[ "$status" != 0 && -f "$TEST_DIR/wiremock.log" ]]; then tail -60 "$TEST_DIR/wiremock.log" >&2; fi
  rm -rf "$TEST_DIR"
}
trap cleanup EXIT

# The standalone distribution is the same WireMock 3.9.1 used by generated tests.
# Fetching the pinned JAR avoids Docker Hub's anonymous pull-rate failure mode.
curl --fail --location --retry 3 --output "$TEST_DIR/wiremock.jar" \
  https://repo.maven.apache.org/maven2/org/wiremock/wiremock-standalone/3.9.1/wiremock-standalone-3.9.1.jar
python3 - "$TEST_DIR/wiremock.jar" <<'PY'
import hashlib, pathlib, sys
expected = "8906708a6549b26d3c2a98fd0807f2edb5e59c3deaf6926ee880127f6952a704dac4f28cc07da9d06303e23f00cee1b5c5b24dfe8fa259fce0be9bb5d25dfadb"
if hashlib.sha512(pathlib.Path(sys.argv[1]).read_bytes()).hexdigest() != expected:
    sys.exit("WireMock distribution checksum mismatch")
PY
mkdir "$TEST_DIR/mappings"
cp "$SDK/wiremock/wiremock-mappings.json" "$TEST_DIR/mappings/wiremock-mappings.json"
"${SDK_JAVA_BIN:-java}" -Xmx1g -jar "$TEST_DIR/wiremock.jar" --port 8080 --root-dir "$TEST_DIR" >"$TEST_DIR/wiremock.log" 2>&1 &
WIREMOCK_PID=$!
ready=false
for _ in {1..60}; do
  if curl --fail --silent http://127.0.0.1:8080/__admin/health >/dev/null; then ready=true; break; fi
  kill -0 "$WIREMOCK_PID"
  sleep 1
done
[[ "$ready" == true ]] || { echo 'WireMock did not become healthy' >&2; exit 1; }
cd "$SDK"
export WIREMOCK_URL=http://127.0.0.1:8080
go build ./...
go vet ./...
go test -count=1 -p=1 ./...
