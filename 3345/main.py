import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def smallestNumber(self, n: int, t: int) -> int:
        def prod(n: int) -> int:
            result = 1
            for digit in str(n):
                result *= int(digit)
            return result

        while prod(n) % t != 0:
            n += 1

        return n


run_tests(
    Solution().smallestNumber,
    [
        {"input": [10, 2], "expected": 10},
        {"input": [15, 3], "expected": 16},
    ],
)
