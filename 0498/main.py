import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests

"""
Solution:

- We can enumerate the diagonals. For the example below, we have

    1 2 3 -> 0 1 2
    4 5 6 -> 1 2 3
    7 8 9 -> 2 3 4

    diag 0 -> 1
    diag 1 -> 2 4
    diag 2 -> 3 5 7
    diag 3 -> 6 8
    diag 4 -> 9

- The idea is to generate the diagonals, then concatenate them.
- We start grouping the diagonals by their index
- By default, the diagonal is built from top to down. This is correct
  for the diagonals in odd indexes. So, we have to reverse the
  diagonal for even indexes
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
