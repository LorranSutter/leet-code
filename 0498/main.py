import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests

"""
Solution:

- The key insight is that every cell on the same anti-diagonal shares the same `i + j` value, so instead
  of walking the matrix diagonal by diagonal we can make a single pass over it and drop each cell
  straight into `diagonals[i + j]`. For the example below, `i + j` gives us

    1 2 3 -> 0 1 2
    4 5 6 -> 1 2 3
    7 8 9 -> 2 3 4

    diag 0 -> 1
    diag 1 -> 2 4
    diag 2 -> 3 5 7
    diag 3 -> 6 8
    diag 4 -> 9

- The idea is to generate the diagonals first, then concatenate them.
- By default, each diagonal comes out built top to bottom, which is already the direction the output
  wants for odd-indexed diagonals. So we only reverse the even-indexed diagonals before flattening
  everything into the result, which gives the zig-zag order the problem asks for.
"""


class Solution:
    def findDiagonalOrder(self, mat: List[List[int]]) -> List[int]:
        m, n = len(mat), len(mat[0])
        max_diagonal_num = m + n - 1

        # Group diagonals by index
        diagonals = [[] for _ in range(max_diagonal_num)]

        for i in range(m):
            for j in range(n):
                diagonals[i + j].append(mat[i][j])

        result = []
        for i in range(max_diagonal_num):
            if i % 2 == 0:
                diagonals[i].reverse()
            result.extend(diagonals[i])

        return result


run_tests(
    Solution().findDiagonalOrder,
    [
        {
            "input": [[[1, 2, 3], [4, 5, 6], [7, 8, 9]]],
            "expected": [1, 2, 4, 7, 5, 3, 6, 8, 9],
        },
        {"input": [[[1, 2], [3, 4]]], "expected": [1, 2, 3, 4]},
        {"input": [[[7], [9], [6]]], "expected": [7, 9, 6]},
        {"input": [[[1, 2, 3, 4, 5, 6]]], "expected": [1, 2, 3, 4, 5, 6]},
        {
            "input": [
                [
                    [7, 8, -9],
                    [10, -11, 12],
                    [-13, 14, -15],
                    [16, 17, -18],
                    [19, -20, 21],
                ]
            ],
            "expected": [7, 8, 10, -13, -11, -9, 12, 14, 16, 19, 17, -15, -18, -20, 21],
        },
    ],
)
