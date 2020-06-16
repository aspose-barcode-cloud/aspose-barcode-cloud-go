#!/bin/bash

set -euo pipefail

for filename in ./docs/*.md; do
  echo "$filename"
  sed -i -e 's/[[:space:]]*$//' "$filename"
done
