import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def findClosestNumber(self, nums: List[int]) -> int:
        closest = nums[0]
        for num in nums:
            if abs(num) < abs(closest):
                closest = num
            elif abs(num) == abs(closest):
                closest = max(num, closest)
        return closest


run_tests(
    Solution().findClosestNumber,
    [
        {"input": [[-4, -2, 1, 4, 8]], "expected": 1},
        {"input": [[2, -1, 1]], "expected": 1},
    ],
)
