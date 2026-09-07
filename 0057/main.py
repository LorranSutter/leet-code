import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def insert(
        self, intervals: List[List[int]], newInterval: List[int]
    ) -> List[List[int]]:
        if len(intervals) == 0:
            return [newInterval]

        inserted = False
        for i in range(len(intervals)):
            if newInterval[0] <= intervals[i][0]:
                intervals.insert(i, newInterval)
                inserted = True
                break

        if not inserted:
            intervals.append(newInterval)

        result = [intervals[0]]
        for interval in intervals[1:]:
            if interval[0] <= result[-1][1]:
                result[-1][1] = max(interval[1], result[-1][1])
            else:
                result.append(interval)

        return result


run_tests(
    Solution().insert,
    [
        {"input": [[[1, 3], [6, 9]], [2, 5]], "expected": [[1, 5], [6, 9]]},
        {
            "input": [[[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]],
            "expected": [[1, 2], [3, 10], [12, 16]],
        },
        {"input": [[[1, 5]], [6, 8]], "expected": [[1, 5], [6, 8]]},
    ],
)
