import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def sumCounts(self, nums: List[int]) -> int:
        n = len(nums)
        i, size = 0, 1
        result = 0
        while size <= n:
            while i + size <= n:
                distinct = len(set(nums[i : i + size]))
                result += distinct**2
                i += 1
            i = 0
            size += 1

        return result


run_tests(
    Solution().sumCounts,
    [
        {"input": [[1, 2, 1]], "expected": 15},
        {"input": [[1, 1]], "expected": 3},
        {
            "input": [
                [43, 9, 37, 14, 24, 20, 6, 33, 25, 38, 46, 32, 10, 35, 18, 19, 31]
            ],
            "expected": 8721,
        },
    ],
)
