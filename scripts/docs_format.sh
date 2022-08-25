#!/bin/bash
set -euo pipefail

format_md_file () {
  echo "$1"
  sed -i -e 's/[[:space:]]*$//' "$1"
}

format_md_file "README.md"

for filename in ./docs/*.md; do
  format_md_file "$filename"
done

echo "Fixing broken link in docs/FilesUploadResult.md"
sed -i -e 's/(Error.md)/(ModelError.md)/' "docs/FilesUploadResult.md"
