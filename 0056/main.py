import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals = sorted(intervals, key=lambda x: (x[0], x[1]))
        merged = [intervals[0]]

        for interval in intervals[1:]:
            if interval[0] > merged[-1][-1]:
                merged.append(interval)
            else:
                merged[-1][-1] = max(interval[-1], merged[-1][-1])

        return merged


run_tests(
    Solution().merge,
    [
        {
            "input": [[[1, 3], [2, 6], [8, 10], [15, 18]]],
            "expected": [[1, 6], [8, 10], [15, 18]],
        },
        {"input": [[[1, 4], [4, 5]]], "expected": [[1, 5]]},
        {"input": [[[4, 7], [1, 4]]], "expected": [[1, 7]]},
        {
            "input": [
                [
                    [1, 3],
                    [2, 6],
                    [8, 10],
                    [8, 9],
                    [9, 11],
                    [15, 18],
                    [2, 4],
                    [16, 17],
                ]
            ],
            "expected": [[1, 6], [8, 11], [15, 18]],
        },
        {"input": [[[1, 4], [0, 0]]], "expected": [[0, 0], [1, 4]]},
        {"input": [[[1, 4], [4, 5]]], "expected": [[1, 5]]},
    ],
)
