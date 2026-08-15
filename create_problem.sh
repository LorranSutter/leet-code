#!/bin/bash

# Check if argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <number> [--ts|--py]"
    echo "Example: $0 23"
    echo "Example: $0 23 --ts"
    echo "Example: $0 23 --py"
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
import { runTests } from "../utils/utils.ts"

function solve(nums: number[]): number {
    // TODO Implement solution
    return 0
}

runTests(solve, [
    { input: [[1, 2, 3]], expected: 0 },
])
EOF

    echo "Created file: $folder_name/main.ts"
elif [ "$2" = "--py" ]; then
    # Create main.py file
    cat > "$folder_name/main.py" << 'EOF'
import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def solve(self, nums: list[int]) -> int:
        # TODO Implement solution
        return 0


run_tests(Solution().solve, [
    {"input": [[1, 2, 3]], "expected": 0},
])
EOF

    echo "Created file: $folder_name/main.py"
else
    # Create main.go file
    cat > "$folder_name/main.go" << 'EOF'
package main

import (
    "leetcode/utils"
)

func solve(nums []int) int {
    // TODO Implement solution
    return 0
}

func main() {
    utils.RunTests([]utils.TestCase[int]{
        {Input: []int{1, 2, 3}, Got: solve([]int{1, 2, 3}), Expected: 0},
    })
}
EOF

    echo "Created file: $folder_name/main.go"
fi

echo "Done!"