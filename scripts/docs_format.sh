#!/bin/bash

set -euo pipefail

for filename in ./docs/*.md; do
  echo "$filename"
  sed -i -e 's/[[:space:]]*$//' "$filename"
done

echo "Fixing broken link in docs/FilesUploadResult.md"
sed -i -e 's/(Error.md)/(ModelError.md)/' "docs/FilesUploadResult.md"
