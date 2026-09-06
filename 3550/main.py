import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def smallestIndex(self, nums: List[int]) -> int:
        def sum_digits(num: int) -> int:
            return sum(int(n) for n in str(num))

        for i in range(len(nums)):
            if i == sum_digits(nums[i]):
                return i
        return -1


run_tests(
    Solution().smallestIndex,
    [
        {"input": [[1, 3, 2]], "expected": 2},
        {"input": [[1, 10, 11]], "expected": 1},
        {"input": [[1, 2, 3]], "expected": -1},
    ],
)
