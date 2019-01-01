#!/bin/sh

set -eu

TEMP_DIR=$(mktemp -d)
TEST_DIR="$(cd "$(dirname "${0}")"; echo "$(pwd)")"

ALIASES=$(cd "${TEST_DIR}/../../..//dist"; echo "$(pwd)/aliases -c ${TEST_DIR}/aliases.yaml")
DIFF=$(if which colordiff >/dev/null; then echo "colordiff -Buw --strip-trailing-cr"; else echo "diff -Bw"; fi)
MASK="sed -e s|${HOME}|[HOME]|g -e s|${TEMP_DIR}|[TEMP_DIR]|g"

${ALIASES} gen --export-path "${TEMP_DIR}" | ${MASK} | sort | ${DIFF} ${TEST_DIR}/alias -
${ALIASES} gen --export --export-path "${TEMP_DIR}" | ${MASK} | ${DIFF} ${TEST_DIR}/export -
cat ${TEMP_DIR}/alpine | ${MASK} | ${DIFF} ${TEST_DIR}/alpine -

${TEMP_DIR}/alpine /bin/sh -c "echo foo" | ${DIFF} ${TEST_DIR}/stdout -
${ALIASES} run /usr/local/bin/alpine /bin/sh -c "echo foo" | ${DIFF} ${TEST_DIR}/stdout -