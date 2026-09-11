import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        intersection = set(nums1) & set(nums2)

        if intersection != set():
            return min(intersection)
        return -1


run_tests(
    Solution().getCommon,
    [
        {"input": [[1, 2, 4], [2, 4]], "expected": 2},
        {"input": [[1, 2, 3, 6], [2, 3, 4, 5]], "expected": 2},
        {"input": [[1, 2, 3, 4], [5, 6, 7]], "expected": -1},
    ],
)
