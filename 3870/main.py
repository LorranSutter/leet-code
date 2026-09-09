import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def countCommas(self, n: int) -> int:
        if n < 1000:
            return 0

        commas = n % 1000 + 1
        n = n // 1000

        if n > 1:
            commas += (n - 1) * 1000

        return commas


run_tests(
    Solution().countCommas,
    [
        {"input": [1002], "expected": 3},
        {"input": [998], "expected": 0},
        {"input": [10002], "expected": 9003},
        {"input": [15002], "expected": 14003},
    ],
)
