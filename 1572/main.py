import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        n = len(mat)
        main_diag = sum((mat[i][i] for i in range(n)))
        sec_diag = sum((mat[i][n - i - 1] for i in range(n)))

        if n % 2 != 0:
            sec_diag -= mat[n // 2][n // 2]

        return main_diag + sec_diag


run_tests(
    Solution().diagonalSum,
    [
        {"input": [[[1, 2, 3], [4, 5, 6], [7, 8, 9]]], "expected": 25},
        {
            "input": [[[1, 1, 1, 1], [1, 1, 1, 1], [1, 1, 1, 1], [1, 1, 1, 1]]],
            "expected": 8,
        },
        {"input": [[[5]]], "expected": 5},
    ],
)
