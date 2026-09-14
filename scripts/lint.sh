#!/bin/bash
set -euo pipefail

mapfile -t vet_packages < <(go list ./... | grep -v '/snippets')
mapfile -t staticcheck_packages < <(printf '%s\n' "${vet_packages[@]}" | grep -v '/test$')

go vet -v "${vet_packages[@]}"
staticcheck "${staticcheck_packages[@]}"
