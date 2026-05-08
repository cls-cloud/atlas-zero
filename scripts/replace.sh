#!/bin/bash
set -euo pipefail

target=${1:-.}

if [ -d "$target" ]; then
    echo "Scanning directory: $target"
    find "$target" -type f -path "*/model/*" -name "*gen.go" | while read -r file; do
        bash "$0" "$file"
    done
    echo "Cleaning up backup files (*.bak)..."
    find "$target" -type f -name "*.bak" -delete
    echo "All replacements done. Backups removed."
    exit 0
fi

file="$target"
tmp_file="${file}.tmp"

if [ ! -f "$file" ]; then
    echo "Target is not a regular file: $file" >&2
    exit 1
fi

if grep -qE '^[[:space:]]*DelFlag[[:space:]]+[a-zA-Z0-9_]+' "$file"; then
    echo "Processing: $file"

    sed -i.bak -E \
        's/^[[:space:]]*DelFlag[[:space:]]+[a-zA-Z0-9_]+[[:space:]]+`[^`]*`.*$/\tDelFlag soft_delete.DeletedAt `gorm:"softDelete:flag"`/' \
        "$file"

    if ! grep -q 'gorm.io/plugin/soft_delete' "$file"; then
        echo "Adding import for soft_delete in $file"
        if ! awk '
            /^import[[:space:]]*\(/ && !added {
                print
                print "\t\"gorm.io/plugin/soft_delete\""
                added=1
                next
            }
            /^import[[:space:]]+"[^"]+"/ && !added {
                sub(/^import[[:space:]]+/, "")
                print "import ("
                print "\t\"gorm.io/plugin/soft_delete\""
                print "\t" $0
                print ")"
                added=1
                next
            }
            { print }
            END {
                if (!added) {
                    exit 2
                }
            }
        ' "$file" > "$tmp_file"; then
            rm -f "$tmp_file"
            echo "Failed to add soft_delete import in $file" >&2
            exit 1
        fi
        mv "$tmp_file" "$file"
    fi

    gofmt -w "$file"
fi
