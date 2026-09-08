import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        k = 1
        for i in range(1, len(nums)):
            if nums[i] != nums[i - 1]:
                nums[k] = nums[i]
                k += 1

        return k


run_tests(
    Solution().removeDuplicates,
    [
        {"input": [[1, 1, 2]], "expected": 2},
        {"input": [[0, 0, 1, 1, 1, 2, 2, 3, 3, 4]], "expected": 5},
        {"input": [[1, 1]], "expected": 1},
    ],
)
