import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        p1 = 0
        p2 = len(nums)

        while p1 < p2:
            if nums[p1] == val:
                nums[p1], nums[p2 - 1] = nums[p2 - 1], nums[p1]
                p2 -= 1
            else:
                p1 += 1

        return p2


run_tests(
    Solution().removeElement,
    [
        {"input": [[3, 2, 2, 3], 3], "expected": 2},
        {"input": [[0, 1, 2, 2, 3, 0, 4, 2], 2], "expected": 5},
        {"input": [[1], 1], "expected": 0},
    ],
)
