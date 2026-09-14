#!/bin/bash
set -euo pipefail

mapfile -t vet_packages < <(go list ./... | grep -v '/snippets')

go vet -v "${vet_packages[@]}"
