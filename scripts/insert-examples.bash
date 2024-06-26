#!/bin/bash
set -euo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

go run "${SCRIPT_DIR}/insert-examples.go"

rm "README.template"
