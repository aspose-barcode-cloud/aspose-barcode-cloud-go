#!/bin/bash

set -euo pipefail

go_mod_file="go.mod"
new_go_mod_file="go.mod.clean"

grep -v " // indirect$" "${go_mod_file}" > "${new_go_mod_file}"

# Replace go.mod file with cleaned up file
mv "${new_go_mod_file}" "${go_mod_file}"
