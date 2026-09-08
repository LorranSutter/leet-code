import sys
from typing import List, Tuple, Set
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests, equal_matrix_unordered


class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        m, n = len(heights), len(heights[0])
        directions = ((-1, 0), (0, 1), (1, 0), (0, -1))

        def dfs(cell: Tuple[int], last_value: int, valid: Set[Tuple[int]]):
            if cell in valid:
                return
            if cell[0] < 0 or cell[0] >= m or cell[1] < 0 or cell[1] >= n:
                return
            if heights[cell[0]][cell[1]] < last_value:
                return

            valid.add(cell)
            for dr, dc in directions:
                dfs((cell[0] + dr, cell[1] + dc), heights[cell[0]][cell[1]], valid)

        valid_pacific = set()
        valid_atlantic = set()
        for i in range(m):
            dfs((i, 0), 0, valid_pacific)
            dfs((i, n - 1), 0, valid_atlantic)
        for i in range(n):
            dfs((0, i), 0, valid_pacific)
            dfs((m - 1, i), 0, valid_atlantic)

        return list(valid_pacific & valid_atlantic)


run_tests(
    Solution().pacificAtlantic,
    [
        {
            "input": [
                [
                    [1, 2, 2, 3, 5],
                    [3, 2, 3, 4, 4],
                    [2, 4, 5, 3, 1],
                    [6, 7, 1, 4, 5],
                    [5, 1, 1, 2, 4],
                ]
            ],
            "expected": [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]],
        },
        {"input": [[[1]]], "expected": [[0, 0]]},
        {"input": [[[1, 2], [2, 2]]], "expected": [[0, 1], [1, 0], [1, 1]]},
        {"input": [[[1, 2], [1, 2]]], "expected": [[0, 0], [0, 1], [1, 0], [1, 1]]},
        {
            "input": [[[1, 2, 3], [8, 9, 4], [7, 6, 5]]],
            "expected": [[0, 2], [1, 0], [1, 1], [1, 2], [2, 0], [2, 1], [2, 2]],
        },
    ],
    is_equal=equal_matrix_unordered,
)
