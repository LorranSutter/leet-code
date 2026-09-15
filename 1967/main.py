import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def numOfStrings(self, patterns: List[str], word: str) -> int:
        return sum(p in word for p in patterns)


run_tests(
    Solution().numOfStrings,
    [
        {"input": [["a", "abc", "bc", "d"], "abc"], "expected": 3},
        {"input": [["a", "b", "c"], "aaaaabbbbb"], "expected": 2},
        {"input": [["a", "a", "a"], "ab"], "expected": 3},
    ],
)
