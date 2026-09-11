import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def maximumProduct(self, nums: List[int]) -> int:
        if len(nums) == 3:
            return nums[0] * nums[1] * nums[2]

        nums.sort(reverse=True)

        # 3 positives, 1 positive and 2 negatives
        res1 = nums[0] * nums[1] * nums[2]
        # 2 negatives and 1 positive
        res2 = nums[0] * nums[-1] * nums[-2]

        return max(res1, res2)


run_tests(
    Solution().maximumProduct,
    [
        {"input": [[1, 2, 3]], "expected": 6},
        {"input": [[-1, -2, -3]], "expected": -6},
        {"input": [[-1, -2, -3, 3]], "expected": 18},
        {"input": [[-1, -2, -3, 3, 4]], "expected": 24},
        {"input": [[-1, -2, -3, 0, 0]], "expected": 0},
        {"input": [[-100, -98, -1, 2, 3, 4]], "expected": 39200},
    ],
)
