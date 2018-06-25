#!/bin/bash
# shellcheck source=./scripts/logs.sh
#
# This script asserts that the "dep" available on path complies with the
# version as specified by the DEP_RELEASE_TAG environment variable. It is
# designed to work both inside a linux container and Max OS X.
#
# The DEP_RELEASE_TAG value is used by golang/dep/master/install.sh

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source "$DIR"/logs.sh
set -euo pipefail

if [ -z "${DEP_RELEASE_TAG:-}" ] ; then
  log_error "DEP_RELEASE_TAG is required"
  exit 1
fi

dep=$(which dep || true)
if [ ! -z "$dep" ]; then
  if dep version | grep -q "$DEP_RELEASE_TAG"; then
    log "dep version is $DEP_RELEASE_TAG... nothing to do"
    exit 0
  else
    log "dep version does not match constraint, reinstalling"
  fi
fi

log "curl -v https://raw.githubusercontent.com/golang/dep/master/install.sh | sh"
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
log "successfully installed dep@$DEP_RELEASE_TAG"
