#!/bin/sh

set -eu

TEMP_DIR=$(mktemp -d)
TEST_DIR="$(cd "$(dirname "${0}")"; echo "$(pwd)")"

ALIASES=$(cd "${TEST_DIR}/../../..//dist"; echo "$(pwd)/aliases -c ${TEST_DIR}/aliases.yaml")
DIFF=$(if which colordiff >/dev/null; then echo "colordiff -Buw --strip-trailing-cr"; else echo "diff -Bw"; fi)
MASK="sed -e s|${HOME}|[HOME]|g -e s|${TEMP_DIR}|[TEMP_DIR]|g -e s|[0-9]*-[0-9]*-[0-9]*T[0-9]*:[0-9]*:[0-9]*Z|yyyy-mm-ddThh:MM:ssZ|g"

${ALIASES} gen --export-path "${TEMP_DIR}" | ${MASK} | sort | ${DIFF} ${TEST_DIR}/alias -
${ALIASES} gen --export --export-path "${TEMP_DIR}" | ${MASK} | ${DIFF} ${TEST_DIR}/export -
cat ${TEMP_DIR}/alpine1 | ${MASK} | ${DIFF} ${TEST_DIR}/alpine1 -
cat ${TEMP_DIR}/alpine2 | ${MASK} | ${DIFF} ${TEST_DIR}/alpine2 -

${TEMP_DIR}/alpine2 /bin/sh -c "alpine3 /bin/sh -c \"alpine5 /bin/sh -c \\\"echo 1\\\"\"" | ${MASK} | ${DIFF} ${TEST_DIR}/stdout -
${ALIASES} run /usr/local/bin/alpine2 /bin/sh -c "alpine3 /bin/sh -c \"alpine5 /bin/sh -c \\\"echo 1\\\"\"" | ${MASK} | ${DIFF} ${TEST_DIR}/stdout -