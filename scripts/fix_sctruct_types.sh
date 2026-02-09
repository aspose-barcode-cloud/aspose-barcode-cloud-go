#!/bin/bash
set -euo pipefail

# Find all the files in the directory and apply the replacements
find "barcode" -type f -name "api_*.go" | while read -r file; do
  # Replace optional.string with optional.String and optional.float32 with optional.Float32
  sed -i 's/optional\.string/optional\.String/g' "$file"
  sed -i 's/optional\.float32/optional\.Float32/g' "$file"
  sed -i 's/optional\.int32/optional\.Int32/g' "$file"
  sed -i 's/optional\.bool/optional\.Bool/g' "$file"
done
