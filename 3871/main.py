import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def countCommas(self, n: int) -> int:
        commas = 0
        thousand = 1000

        if n < thousand:
            return commas

        while n >= thousand:
            commas += n - thousand + 1
            thousand *= 1000

        return commas


run_tests(
    Solution().countCommas,
    [
        {"input": [1002], "expected": 3},
        {"input": [998], "expected": 0},
        {"input": [10002], "expected": 9003},
        {"input": [15002], "expected": 14003},
        {"input": [190085040002001], "expected": 759339159007008},
    ],
)
