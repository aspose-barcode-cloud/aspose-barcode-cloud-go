#!/bin/bash
set -euo pipefail

mapfile -t vet_packages < <(go list ./... | grep -v '/snippets')

go vet -v "${vet_packages[@]}"

# Snippet directories contain multiple standalone programs with duplicate main functions,
# so each snippet must be vetted independently instead of using go vet ./....
while IFS= read -r -d '' snippet; do
    go vet -v "${snippet}"
done < <(find snippets -type f -name '*.go' -print0)
