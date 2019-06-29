#!/usr/bin/env bash

load './bin/bats-support/load'
load './bin/bats-assert/load'

TEST_DIR="."

BIN_UNDER_TEST='./dist/simulator'

_global_setup() {
  [ ! -f ${BATS_PARENT_TMPNAME}.skip ] || skip "skip remaining tests"
  export SIMULATOR_MANIFEST_PATH='./fixtures/valid'
  export SIMULATOR_TF_DIR='./fixtures/noop-tf-dir'
  export SIMULATOR_CLI_TEST_OUTPUT='test.debug'
}

_global_teardown() {
  if [ ! -n "$BATS_TEST_COMPLETED" ]; then
    touch ${BATS_PARENT_TMPNAME}.skip
  fi
}

_app() {
  local ARGS="${@:-}"
  ./../${BIN_UNDER_TEST:-false} "${ARGS}";
}
