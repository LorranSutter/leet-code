import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def reverseDegree(self, s: str) -> int:
        return sum((i + 1) * (ord("z") - ord(s[i]) + 1) for i in range(len(s)))


run_tests(
    Solution().reverseDegree,
    [
        {"input": ["abc"], "expected": 148},
        {"input": ["zaza"], "expected": 160},
    ],
)
