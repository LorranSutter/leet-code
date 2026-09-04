import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def removeDuplicates(self, s: str) -> str:
        i = 0
        while i < len(s) - 1:
            if s[i] == s[i + 1]:
                s = s[:i] + s[i + 2 :]
                i = max(0, i - 2)
            else:
                i += 1
        return s


run_tests(
    Solution().removeDuplicates,
    [
        {"input": ["abbaca"], "expected": "ca"},
        {"input": ["azxxzy"], "expected": "ay"},
        {"input": ["aaaaaaaa"], "expected": ""},
    ],
)
