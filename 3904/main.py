import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def firstStableIndex(self, nums: list[int], k: int) -> int:
        # TODO Implement solution
        return 0


run_tests(
    Solution().solve,
    [
        {"input": [[5, 0, 1, 4], 3], "expected": 3},
        {"input": [[3, 2, 1], 1], "expected": -1},
        {"input": [[0], 0], "expected": 0},
    ],
)
