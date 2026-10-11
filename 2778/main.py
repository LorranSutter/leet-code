import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def sumOfSquares(self, nums: List[int]) -> int:
        n = len(nums)

        total = 0
        for i in range(n):
            if n % (i + 1) == 0:
                total += nums[i] * nums[i]

        return total


run_tests(
    Solution().sumOfSquares,
    [
        {"input": [[1, 2, 3, 4]], "expected": 21},
        {"input": [[2, 7, 1, 19, 18, 3]], "expected": 63},
    ],
)
