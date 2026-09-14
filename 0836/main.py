import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def isRectangleOverlap(self, rec1: List[int], rec2: List[int]) -> bool:
        #        0   1   2   3
        # rec1 [x1, y1, x2, y2]
        # rec2 [x1, y1, x2, y2]
        #
        #             ----------(x2, y2)
        #            |             |
        #            |             |
        #     -----------(x2, y1)  |
        #    |       |      |      |
        #    |    (x1,y1)---|------
        #    |              |
        # (x1,y1)-----------

        # Check if intervals on x and y overlaps
        x_axis = max(rec1[0], rec2[0]) < min(rec1[2], rec2[2])
        y_axis = max(rec1[1], rec2[1]) < min(rec1[3], rec2[3])

        return x_axis and y_axis


run_tests(
    Solution().isRectangleOverlap,
    [
        {"input": [[0, 0, 2, 2], [1, 1, 3, 3]], "expected": True},
        {"input": [[0, 0, 1, 1], [1, 0, 2, 1]], "expected": False},
        {"input": [[0, 0, 1, 1], [2, 2, 3, 3]], "expected": False},
    ],
)
