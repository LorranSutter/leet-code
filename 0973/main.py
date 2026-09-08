import sys
import math
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests, equal_matrix_unordered


class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        points = sorted(points, key=lambda p: math.sqrt(p[0] ** 2 + p[1] ** 2))

        return points[:k]


run_tests(
    Solution().kClosest,
    [
        {"input": [[[1, 3], [-2, 2]], 1], "expected": [[-2, 2]]},
        {"input": [[[3, 3], [5, -1], [-2, 4]], 2], "expected": [[3, 3], [-2, 4]]},
    ],
    is_equal=equal_matrix_unordered,
)
