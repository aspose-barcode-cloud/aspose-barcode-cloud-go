#!/bin/bash

set -euo pipefail

go_sum_file="go.sum"
new_go_sum_file="go.sum.clean"
cat /dev/null > "$new_go_sum_file"

started=false
while read -r line
do
    if [ "$line" == "require (" ]; then
        started=true
        continue
    fi
    if [ "$line" == ")" ]; then
        break
    fi

    if [ "$started" == true ]; then
        echo "$line"
        without_indirect=$(sed "s/ \/\/ indirect//g" <<< "$line")
        echo "\"$without_indirect\""
        grep -iF "$without_indirect" "$go_sum_file" >> "$new_go_sum_file"
    fi
done < "go.mod"

# Replace go.sum file with cleaned up file
mv "$new_go_sum_file" "$go_sum_file"
