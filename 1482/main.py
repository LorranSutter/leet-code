import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        def make_bouquets(day: int) -> int:
            made = 0
            current_bouquet = k
            for bloom in bloomDay:
                if bloom <= day:
                    current_bouquet -= 1
                    if current_bouquet == 0:
                        current_bouquet = k
                        made += 1
                else:
                    current_bouquet = k

            return made

        low, high = 1, max(bloomDay)
        result = -1
        while low <= high:
            mid = (high + low) // 2

            if make_bouquets(mid) >= m:
                result = mid
                high = mid - 1
            else:
                low = mid + 1

        return result


run_tests(
    Solution().minDays,
    [
        {"input": [[1, 10, 3, 10, 2], 3, 1], "expected": 3},
        {"input": [[1, 10, 3, 10, 2], 3, 2], "expected": -1},
        {"input": [[7, 7, 7, 7, 12, 7, 7], 2, 3], "expected": 12},
    ],
)
