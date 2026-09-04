import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def distanceBetweenBusStops(
        self, distance: List[int], start: int, destination: int
    ) -> int:
        if destination < start:
            start, destination = destination, start

        clockwise_dist = sum(distance[start:destination])
        counter_clockwise_dist = sum(distance) - clockwise_dist

        return min(clockwise_dist, counter_clockwise_dist)


run_tests(
    Solution().distanceBetweenBusStops,
    [
        {"input": [[1, 2, 3, 4], 0, 1], "expected": 1},
        {"input": [[1, 2, 3, 4], 0, 2], "expected": 3},
        {"input": [[1, 2, 3, 4], 0, 3], "expected": 4},
        {"input": [[7, 10, 1, 12, 11, 14, 5, 0], 7, 2], "expected": 17},
    ],
)
