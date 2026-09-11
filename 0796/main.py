import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        for _ in range(len(s)):
            if s == goal:
                return True
            s = s[1:] + s[0]
        return False


run_tests(
    Solution().rotateString,
    [
        {"input": ["abcde", "cdeab"], "expected": True},
        {"input": ["abcde", "abced"], "expected": False},
    ],
)
