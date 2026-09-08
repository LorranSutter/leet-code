import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        p1, p2 = m - 1, n - 1
        for i in range(m + n - 1, -1, -1):
            if p1 < 0:
                nums1[i] = nums2[p2]
                p2 -= 1
            elif p2 < 0:
                break
            elif nums1[p1] > nums2[p2]:
                nums1[i] = nums1[p1]
                p1 -= 1
            else:
                nums1[i] = nums2[p2]
                p2 -= 1


run_tests(
    Solution().merge,
    [
        {"input": [[1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3], "expected": None},
        {"input": [[1], 1, [], 0], "expected": None},
        {"input": [[0], 0, [1], 1], "expected": None},
        {"input": [[2, 0], 1, [1], 1], "expected": None},
    ],
)
