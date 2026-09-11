import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def maxDistance(self, colors: List[int]) -> int:
        for i in range(len(colors) - 1, -1, -1):
            if colors[0] != colors[i]:
                return i
            if colors[-1] != colors[len(colors) - i]:
                return i - 1
        return 1


run_tests(
    Solution().maxDistance,
    [
        {"input": [[1, 1, 1, 6, 1, 1, 1]], "expected": 3},
        {"input": [[1, 8, 3, 8, 3]], "expected": 4},
        {"input": [[0, 1]], "expected": 1},
        {"input": [[1, 1, 1, 1, 6, 1, 1]], "expected": 4},
        {"input": [[4, 4, 4, 11, 4, 4, 11, 4, 4, 4, 4, 4]], "expected": 8},
    ],
)
