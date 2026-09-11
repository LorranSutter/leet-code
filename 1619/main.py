import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def trimMean(self, arr: List[int]) -> float:
        n = len(arr)
        amount_to_remove = n // 20

        arr.sort()
        arr = arr[amount_to_remove:-amount_to_remove]

        return sum(arr) / (n - 2 * amount_to_remove)


run_tests(
    Solution().trimMean,
    [
        {
            "input": [[1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3]],
            "expected": 2.000000,
        },
        {
            "input": [[6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0]],
            "expected": 4.000000,
        },
        {
            "input": [
                [
                    6,
                    0,
                    7,
                    0,
                    7,
                    5,
                    7,
                    8,
                    3,
                    4,
                    0,
                    7,
                    8,
                    1,
                    6,
                    8,
                    1,
                    1,
                    2,
                    4,
                    8,
                    1,
                    9,
                    5,
                    4,
                    3,
                    8,
                    5,
                    10,
                    8,
                    6,
                    6,
                    1,
                    0,
                    6,
                    10,
                    8,
                    2,
                    3,
                    4,
                ]
            ],
            "expected": 4.777777777777778,
        },
    ],
)
