import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def transpose(self, matrix: List[List[int]]) -> List[List[int]]:
        m, n = len(matrix), len(matrix[0])
        new_matrix = [[0 for _ in range(m)] for _ in range(n)]

        for i in range(n):
            for j in range(m):
                new_matrix[i][j] = matrix[j][i]

        return new_matrix


run_tests(
    Solution().transpose,
    [
        {
            "input": [[[1, 2, 3], [4, 5, 6], [7, 8, 9]]],
            "expected": [[1, 4, 7], [2, 5, 8], [3, 6, 9]],
        },
        {"input": [[[1, 2, 3], [4, 5, 6]]], "expected": [[1, 4], [2, 5], [3, 6]]},
    ],
)
