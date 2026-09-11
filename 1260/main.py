import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def shiftGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        length = m * n
        new_grid = [[0 for _ in range(n)] for _ in range(m)]

        for i in range(m):
            for j in range(n):
                r = ((i * n + j + k) % length) // n
                c = ((i * n + j + k) % length) % n
                new_grid[r][c] = grid[i][j]

        return new_grid


run_tests(
    Solution().shiftGrid,
    [
        {
            "input": [[[1, 2, 3], [4, 5, 6], [7, 8, 9]], 1],
            "expected": [[9, 1, 2], [3, 4, 5], [6, 7, 8]],
        },
        {
            "input": [
                [[3, 8, 1, 9], [19, 7, 2, 5], [4, 6, 11, 10], [12, 0, 21, 13]],
                4,
            ],
            "expected": [[12, 0, 21, 13], [3, 8, 1, 9], [19, 7, 2, 5], [4, 6, 11, 10]],
        },
        {
            "input": [[[1, 2, 3], [4, 5, 6], [7, 8, 9]], 9],
            "expected": [[1, 2, 3], [4, 5, 6], [7, 8, 9]],
        },
    ],
)
