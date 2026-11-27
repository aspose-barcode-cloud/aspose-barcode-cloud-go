#!/bin/bash
set -euo pipefail

mapfile -t vet_packages < <(go list ./... | grep -v '/snippets')

go vet -v "${vet_packages[@]}"
go run honnef.co/go/tools/cmd/staticcheck@v0.8.1 "${vet_packages[@]}"

# Snippet directories contain multiple standalone programs with duplicate main functions,
# so each snippet must be vetted independently instead of using go vet ./....
while IFS= read -r -d '' snippet; do
    go vet -v "${snippet}"
done < <(find snippets -type f -name '*.go' -print0)

# lint/ is a separate module, so the golang.org/x/tools dependency of its
# analyzers stays out of the SDK's go.mod.
go test -C lint ./...
go build -C lint -o bin/binarydecode ./cmd/binarydecode
go vet -vettool="$(pwd)/lint/bin/binarydecode" ./barcode/...
