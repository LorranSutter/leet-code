import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def countLatticePoints(self, circles: List[List[int]]) -> int:
        def inside_circle(p: List[int], x: int, y: int, r: int) -> bool:
            return (x - p[0]) ** 2 + (y - p[1]) ** 2 <= r * r

        points_inside = set()
        for x, y, r in circles:
            min_x, max_x = x - r, x + r
            min_y, max_y = y - r, y + r

            for i in range(min_x, max_x + 1):
                for j in range(min_y, max_y + 1):
                    if (i, j) in points_inside:
                        continue
                    if inside_circle([i, j], x, y, r):
                        points_inside.add((i, j))

        return len(points_inside)


run_tests(
    Solution().countLatticePoints,
    [
        {"input": [[[2, 2, 1]]], "expected": 5},
        {"input": [[[2, 2, 2], [3, 4, 1]]], "expected": 16},
    ],
)
