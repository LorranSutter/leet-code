import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def thirdMax(self, nums: List[int]) -> int:
        nums = sorted(set(nums), reverse=True)

        if len(nums) >= 3:
            return nums[2]
        return nums[0]


run_tests(
    Solution().thirdMax,
    [
        {"input": [[3, 2, 1]], "expected": 1},
        {"input": [[1, 2]], "expected": 2},
        {"input": [[2, 2, 3, 1]], "expected": 1},
        {"input": [[-2147483648, 1, 2]], "expected": -2147483648},
    ],
)
