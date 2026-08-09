#!/bin/bash

# Check if argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <number> [--ts]"
    echo "Example: $0 23"
    echo "Example: $0 23 --ts"
    exit 1
fi

# Validate input is a number
if ! [[ "$1" =~ ^[0-9]+$ ]]; then
    echo "Error: Input must be a number"
    exit 1
fi

# Format number with leading zeros (4 digits)
folder_name=$(printf "%04d" "$1")

# Create folder
if [ -d "$folder_name" ]; then
    echo "Error: Folder $folder_name already exists"
    exit 1
fi

mkdir "$folder_name"
echo "Created folder: $folder_name"

if [ "$2" = "--ts" ]; then
    # Create main.ts file
    cat > "$folder_name/main.ts" << 'EOF'
function solve(): void {
    console.log("Hello, LeetCode!");
}

solve();
EOF

    echo "Created file: $folder_name/main.ts"
else
    # Create main.go file
    cat > "$folder_name/main.go" << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello, LeetCode!")
}
EOF

    echo "Created file: $folder_name/main.go"
fi

echo "Done!"