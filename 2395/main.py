import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def findSubarrays(self, nums: List[int]) -> bool:
        sums = set()
        for i in range(len(nums) - 1):
            if nums[i] + nums[i + 1] in sums:
                return True
            sums.add(nums[i] + nums[i + 1])
        return False


run_tests(
    Solution().findSubarrays,
    [
        {"input": [[4, 2, 4]], "expected": True},
        {"input": [[1, 2, 3, 4, 5]], "expected": False},
        {"input": [[0, 0, 0]], "expected": True},
    ],
)
