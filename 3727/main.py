import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        nums = sorted([abs(num) for num in nums], reverse=True)

        n = len(nums)
        result = 0
        for i in range(n // 2):
            result += nums[i] ** 2 - nums[n - i - 1] ** 2

        if n % 2 != 0:
            result += nums[n // 2] ** 2

        return result


run_tests(
    Solution().maxAlternatingSum,
    [
        {"input": [[1, 2, 3]], "expected": 12},
        {"input": [[1, -1, 2, -2, 3, -3]], "expected": 16},
    ],
)
