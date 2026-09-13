import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def mirrorDistance(self, n: int) -> int:
        return abs(n - int(str(n)[::-1]))


run_tests(
    Solution().mirrorDistance,
    [
        {"input": [25], "expected": 27},
        {"input": [10], "expected": 9},
        {"input": [7], "expected": 0},
    ],
)
