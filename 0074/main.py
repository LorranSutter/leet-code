import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        # Search by row: find the last row whose first value is <= target.
        low, high = 0, len(matrix) - 1
        while low <= high:
            mid = (high + low) // 2

            if matrix[mid][0] == target:
                return True
            if matrix[mid][0] > target:
                high = mid - 1
            else:
                low = mid + 1

        # When the loop ends, `high` is the only row that could hold target
        # (or -1 if target is smaller than every row's first value).
        if high < 0:
            return False

        # Search by column
        low, high, row = 0, len(matrix[0]) - 1, high
        while low <= high:
            mid = (high + low) // 2

            if matrix[row][mid] == target:
                return True
            if matrix[row][mid] > target:
                high = mid - 1
            else:
                low = mid + 1

        return False


run_tests(
    Solution().searchMatrix,
    [
        {
            "input": [[[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]], 3],
            "expected": True,
        },
        {
            "input": [[[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]], 13],
            "expected": False,
        },
        {"input": [[[1]], 2], "expected": False},
        {
            "input": [[[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 50]], 11],
            "expected": True,
        },
    ],
)
