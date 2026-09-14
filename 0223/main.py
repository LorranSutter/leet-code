import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def computeArea(
        self,
        ax1: int,
        ay1: int,
        ax2: int,
        ay2: int,
        bx1: int,
        by1: int,
        bx2: int,
        by2: int,
    ) -> int:
        areaA = abs(ax1 - ax2) * abs(ay1 - ay2)
        areaB = abs(bx1 - bx2) * abs(by1 - by2)
        intersection = 0

        # ax1 ---- bx1 ---- ax2 ---- bx2
        # Check if rectangles overlap
        x_axis = max(ax1, bx1) < min(ax2, bx2)
        y_axis = max(ay1, by1) < min(ay2, by2)

        if x_axis and y_axis:
            dx = abs(min(ax2, bx2) - max(ax1, bx1))
            dy = abs(min(ay2, by2) - max(ay1, by1))
            intersection = dx * dy

        return areaA + areaB - intersection


run_tests(
    Solution().computeArea,
    [
        {"input": [-3, 0, 3, 4, 0, -1, 9, 2], "expected": 45},
        {"input": [-2, -2, 2, 2, -2, -2, 2, 2], "expected": 16},
    ],
)
