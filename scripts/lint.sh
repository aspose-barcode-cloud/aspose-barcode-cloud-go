#!/bin/bash
set -euo pipefail

mapfile -t packages < <(go list ./... | grep -v '/snippets' | grep -v '/test$')

go vet -v "${packages[@]}"
staticcheck "${packages[@]}"
