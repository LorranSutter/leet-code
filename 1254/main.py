import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests

from collections import deque


class Solution:
    def closedIsland(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])

        # -1 -> border
        #  0 -> island
        #  1 -> water
        #  2 -> visited island

        for i in range(m):
            if grid[i][0] == 0:
                grid[i][0] = -1
            if grid[i][n - 1] == 0:
                grid[i][n - 1] = -1

        for i in range(n):
            if grid[0][i] == 0:
                grid[0][i] = -1
            if grid[m - 1][i] == 0:
                grid[m - 1][i] = -1

        def dfs(i: int, j: int):
            queue = deque()
            queue.append((i, j))

            directions = [(1, 0), (0, 1), (-1, 0), (0, -1)]
            closed_island = True

            while queue:
                r, c = queue.popleft()

                if grid[r][c] == -1:
                    closed_island = False
                elif grid[r][c] == 0:
                    grid[r][c] = 2
                    for dr, dc in directions:
                        queue.append((r + dr, c + dc))

            return closed_island

        closed_islands = 0
        for i in range(1, m - 1):
            for j in range(1, n - 1):
                if grid[i][j] == 0 and dfs(i, j):
                    closed_islands += 1

        return closed_islands


run_tests(
    Solution().closedIsland,
    [
        {
            "input": [
                [
                    [1, 1, 1, 1, 1, 1, 1, 0],
                    [1, 0, 0, 0, 0, 1, 1, 0],
                    [1, 0, 1, 0, 1, 1, 1, 0],
                    [1, 0, 0, 0, 0, 1, 0, 1],
                    [1, 1, 1, 1, 1, 1, 1, 0],
                ]
            ],
            "expected": 2,
        },
        {"input": [[[0, 0, 1, 0, 0], [0, 1, 0, 1, 0], [0, 1, 1, 1, 0]]], "expected": 1},
        {
            "input": [
                [
                    [1, 1, 1, 1, 1, 1, 1],
                    [1, 0, 0, 0, 0, 0, 1],
                    [1, 0, 1, 1, 1, 0, 1],
                    [1, 0, 1, 0, 1, 0, 1],
                    [1, 0, 1, 1, 1, 0, 1],
                    [1, 0, 0, 0, 0, 0, 1],
                    [1, 1, 1, 1, 1, 1, 1],
                ]
            ],
            "expected": 2,
        },
    ],
)
