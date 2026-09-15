import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def check(self, nums: List[int]) -> bool:
        n = len(nums)
        minn = min(nums)
        min_ids = []

        # Find the index of all possible min values
        for i in range(n):
            if nums[i] == minn:
                min_ids.append(i)

        # Check if any of the min values results in a sorted array
        for idx in min_ids:
            is_sorted = True
            for i in range(idx, n + idx - 1):
                if nums[i % n] > nums[(i + 1) % n]:
                    is_sorted = False
                    break
            if is_sorted:
                return True

        return False


run_tests(
    Solution().check,
    [
        {"input": [[3, 4, 5, 1, 2]], "expected": True},
        {"input": [[2, 1, 3, 4]], "expected": False},
        {"input": [[1, 2, 3]], "expected": True},
        {"input": [[1, 2, 3, 4, 5, 1]], "expected": True},
    ],
)
