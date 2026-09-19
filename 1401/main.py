import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def checkOverlap(
        self,
        radius: int,
        xCenter: int,
        yCenter: int,
        x1: int,
        y1: int,
        x2: int,
        y2: int,
    ) -> bool:
        # Find the closest point to the circle within the rectangle
        closest_x = max(x1, min(x2, xCenter))
        closest_y = max(y1, min(y2, yCenter))

        # Calculate the distance squared from the circle's center to this closest point
        distance_x = xCenter - closest_x
        distance_y = yCenter - closest_y
        distance_squared = (distance_x * distance_x) + (distance_y * distance_y)

        return distance_squared <= (radius * radius)


run_tests(
    Solution().checkOverlap,
    [
        {"input": [1, 0, 0, 1, -1, 3, 1], "expected": True},
        {"input": [1, 0, 0, -1, 0, 0, 1], "expected": True},
    ],
)
