import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def findKOr(self, nums: List[int], k: int) -> int:
        result = 0
        for i in range(31):
            count_ones = 0
            for num in nums:
                if (num >> i) & 1 == 1:
                    count_ones += 1
            if count_ones >= k:
                result |= 1 << i
        return result


run_tests(
    Solution().findKOr,
    [
        {"input": [[7, 12, 9, 8, 9, 15], 4], "expected": 9},
        {"input": [[2, 12, 1, 11, 4, 5], 6], "expected": 0},
        {"input": [[10, 8, 5, 9, 11, 6, 8], 1], "expected": 15},
    ],
)
