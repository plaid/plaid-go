#!/bin/bash

RED="\\033[31m"
GREEN="\\033[32m"
RESET="\\033[0m"

log() {
  echo -e "$GREEN[info]" "$@" "$RESET"
}

log_error() {
  echo -e "$RED[error]" "$@" "$RESET" 1>&2;
}

log_and_execute() {
  log "$@"
  "$@"
}
