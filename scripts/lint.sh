#!/bin/bash
set -euo pipefail

mapfile -t lint_packages < <(go list ./... | grep -v '/snippets')

go vet -v "${lint_packages[@]}"
go run honnef.co/go/tools/cmd/staticcheck@v0.8.1 "${lint_packages[@]}"
