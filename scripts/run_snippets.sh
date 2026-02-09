#!/bin/bash

set -euo pipefail

RUN_DIR="snippets_test"
SNIPPETS_DIR="snippets"
SCRIPT_DIR="scripts"
CONFIG_FILE_PATH="test/configuration.json"

rm -rf "${RUN_DIR}" || true
mkdir -p "${RUN_DIR}"


for file in $(find "${SNIPPETS_DIR}" -name "*.go"); do
    ${SCRIPT_DIR}/run_snippet.sh "$file" $RUN_DIR $SCRIPT_DIR $CONFIG_FILE_PATH || { echo "Error processing $file"; exit 1; }
done

rm -rf "${RUN_DIR}" || true