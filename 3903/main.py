import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def firstStableIndex(self, nums: List[int], k: int) -> int:
        max_nums = [nums[0] for _ in range(len(nums))]
        min_nums = [nums[-1] for _ in range(len(nums))]

        max_num = max_nums[0]
        for i in range(len(nums)):
            if nums[i] > max_num:
                max_num = nums[i]
            max_nums[i] = max_num

        min_num = min_nums[0]
        for i in range(len(nums) - 1, -1, -1):
            if nums[i] < min_num:
                min_num = nums[i]
            min_nums[i] = min_num

        for i in range(len(nums)):
            if max_nums[i] - min_nums[i] <= k:
                return i

        return -1


run_tests(
    Solution().firstStableIndex,
    [
        {"input": [[5, 0, 1, 4], 3], "expected": 3},
        {"input": [[3, 2, 1], 1], "expected": -1},
        {"input": [[0], 0], "expected": 0},
    ],
)
