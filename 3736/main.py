import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def minMoves(self, nums: List[int]) -> int:
        max_element = max(nums)
        return sum((max_element - n) for n in nums)


run_tests(
    Solution().minMoves,
    [
        {"input": [[2, 1, 3]], "expected": 3},
        {"input": [[4, 4, 5]], "expected": 2},
    ],
)
