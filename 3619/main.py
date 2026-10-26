import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def countIslands(self, grid: List[List[int]], k: int) -> int:
        m, n = len(grid), len(grid[0])

        def dfs(r: int, c: int) -> int:
            if r < 0 or c < 0 or r >= m or c >= n:
                return 0
            if grid[r][c] == 0:
                return 0

            value = grid[r][c]
            grid[r][c] = 0
            return value + dfs(r + 1, c) + dfs(r - 1, c) + dfs(r, c + 1) + dfs(r, c - 1)

        num_of_islands = 0
        for i in range(m):
            for j in range(n):
                total_value = dfs(i, j)
                if total_value > 0 and total_value % k == 0:
                    num_of_islands += 1

        return num_of_islands


run_tests(
    Solution().countIslands,
    [
        {
            "input": [
                [
                    [0, 2, 1, 0, 0],
                    [0, 5, 0, 0, 5],
                    [0, 0, 1, 0, 0],
                    [0, 1, 4, 7, 0],
                    [0, 2, 0, 0, 8],
                ],
                5,
            ],
            "expected": 2,
        },
        {"input": [[[3, 0, 3, 0], [0, 3, 0, 3], [3, 0, 3, 0]], 3], "expected": 6},
        {
            "input": [
                [
                    [0, 12, 0, 0, 12, 0],
                    [0, 0, 5, 0, 0, 0],
                    [0, 0, 0, 0, 0, 12],
                    [0, 0, 17, 17, 15, 0],
                    [0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0],
                ],
                12,
            ],
            "expected": 3,
        },
    ],
)
