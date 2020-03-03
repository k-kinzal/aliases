#!/bin/sh

set -euv

TEMP_DIR=$(mktemp -d /tmp/XXXX)
TEST_DIR="$(cd "$(dirname "${0}")"; echo "$(pwd)")"

ALIASES=$(cd "${TEST_DIR}/../../..//dist"; echo "$(pwd)/aliases -c ${TEST_DIR}/aliases.yaml")
DIFF=$(if which colordiff >/dev/null; then echo "colordiff -Buw --strip-trailing-cr"; else echo "diff -Bw"; fi)
MASK="sed -e s|${HOME}|[HOME]|g -e s|${TEMP_DIR}|[TEMP_DIR]|g -e s|\[HOME\]/.aliases/tmp/[a-z0-9]*|[ALIASE_TEMP_DIR]|"

export FOO_PASS_ENV1="1"
export FOO_PASS_ENV2="2"
export FOO_PASS_ENV3="3"

${ALIASES} gen --export-path "${TEMP_DIR}" | ${MASK} | sort | ${DIFF} ${TEST_DIR}/alias -
${ALIASES} gen --export --export-path "${TEMP_DIR}" | ${MASK} | ${DIFF} ${TEST_DIR}/export -
cat ${TEMP_DIR}/alpine1
cat ${TEMP_DIR}/alpine1 | ${MASK} | ${DIFF} ${TEST_DIR}/alpine1 -
cat ${TEMP_DIR}/alpine2
cat ${TEMP_DIR}/alpine2 | ${MASK} | ${DIFF} ${TEST_DIR}/alpine2 -

${TEMP_DIR}/alpine2 /bin/sh -c 'alpine1 sh -c "env"' | grep PASS_ENV | ${DIFF} ${TEST_DIR}/stdout -
${ALIASES} run /usr/local/bin/alpine2 /bin/sh -c 'alpine1 sh -c "env"' | ${MASK} | grep PASS_ENV | ${DIFF} ${TEST_DIR}/stdout -