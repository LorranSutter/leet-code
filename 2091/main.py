import sys
import math
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def minimumDeletions(self, nums: List[int]) -> int:
        maxx = -math.inf
        minn = math.inf
        max_id, min_id = 0, 0

        for i in range(len(nums)):
            if nums[i] > maxx:
                maxx = nums[i]
                max_id = i
            if nums[i] < minn:
                minn = nums[i]
                min_id = i

        # Removal options
        sides = len(nums) - abs(max_id - min_id) + 1
        left = max(max_id, min_id) + 1
        right = len(nums) - min(max_id, min_id)

        return min(sides, left, right)


run_tests(
    Solution().minimumDeletions,
    [
        {"input": [[2, 10, 7, 5, 4, 1, 8, 6]], "expected": 5},
        {"input": [[0, -4, 19, 1, 8, -2, -3, 5]], "expected": 3},
        {"input": [[101]], "expected": 1},
    ],
)
