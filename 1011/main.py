import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        def calculate_days(capacity: int) -> int:
            total_days = 0
            acc = 0
            for w in weights:
                if w + acc > capacity:
                    total_days += 1
                    acc = 0
                acc += w

            if acc > 0:
                total_days += 1

            return total_days

        low = max(weights)
        high = sum(weights)
        while low < high:
            capacity = low + (high - low) // 2
            if calculate_days(capacity) <= days:
                high = capacity
            else:
                low = capacity + 1

        return low


run_tests(
    Solution().shipWithinDays,
    [
        {"input": [[1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 5], "expected": 15},
        {"input": [[3, 2, 2, 4, 1, 4], 3], "expected": 6},
        {"input": [[1, 2, 3, 1, 1], 4], "expected": 3},
        {"input": [[10, 50, 100, 100, 50, 100, 100, 100], 5], "expected": 160},
        {"input": [[5, 5, 5, 5], 10], "expected": 5},
    ],
)
