import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def isPalindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1

        while left <= right:
            if not s[left].isalnum():
                left += 1
                continue
            if not s[right].isalnum():
                right -= 1
                continue
            if s[left].upper() != s[right].upper():
                return False
            left += 1
            right -= 1

        return True


run_tests(
    Solution().isPalindrome,
    [
        {"input": ["A man, a plan, a canal: Panama"], "expected": True},
        {"input": ["race a car"], "expected": False},
        {"input": [" "], "expected": True},
        {"input": ["0P"], "expected": False},
    ],
)
