#!/bin/sh

set -eu

TEMP_DIR=$(mktemp -d)
TEST_DIR="$(cd "$(dirname "${0}")"; echo "$(pwd)")"

ALIASES=$(cd "${TEST_DIR}/../../..//dist"; echo "$(pwd)/aliases -c ${TEST_DIR}/aliases.yaml")
DIFF=$(if which colordiff >/dev/null; then echo "colordiff -Buw --strip-trailing-cr"; else echo "diff -Buw --strip-trailing-cr"; fi)
MASK="sed -e s|${HOME}|[HOME]|g -e s|${TEMP_DIR}|[TEMP_DIR]|g"

${DIFF} ${TEST_DIR}/alias   - <<<"$(${ALIASES} gen --export-path "${TEMP_DIR}" | ${MASK} | sort)"
${DIFF} ${TEST_DIR}/export  - <<<"$(${ALIASES} gen --export --export-path "${TEMP_DIR}" | ${MASK})"
${DIFF} ${TEST_DIR}/alpine  - <<<"$(cat ${TEMP_DIR}/alpine | ${MASK})"

${DIFF} ${TEST_DIR}/stdout  - <<<"$(${TEMP_DIR}/alpine /bin/sh -c "uname -a")"