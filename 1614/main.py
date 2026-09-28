import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def maxDepth(self, s: str) -> int:
        count = 0
        max_count = 0
        for char in s:
            if char == "(":
                count += 1
            elif char == ")":
                max_count = max(max_count, count)
                count -= 1

        return max(max_count, count)


run_tests(
    Solution().maxDepth,
    [
        {"input": ["(1+(2*3)+((8)/4))+1"], "expected": 3},
        {"input": ["(1)+((2))+(((3)))"], "expected": 3},
        {"input": ["()(())((()()))"], "expected": 3},
    ],
)
