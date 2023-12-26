#!/bin/bash
init_module() {
    local dir=$1
    local mod_name=$2

    if [ -f "$dir/main.go" ]; then
        echo "Initialized Go module in $dir ($mod_name)"
        cd "$dir";go mod init "$mod_name";cd - > /dev/null
        echo "fin"
    fi
}

dir_queue=()

ROOT_DIR="."
dir_queue+=("$ROOT_DIR")

while [ ! -z "${dir_queue[*]}" ]; do
    current_dirs=${dir_queue[0]}
    dir_queue=("${dir_queue[@]:1}")

    for current_dir in "$current_dirs"/*; do
        if [ -d "$current_dir" ]; then
            dir_queue+=("$current_dir")
            abs_dir=$(realpath "$current_dir")
            dir_name=$(basename "$current_dir")
            
            init_module "$abs_dir" "$dir_name"
        fi
    done
done
