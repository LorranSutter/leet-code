import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def isHappy(self, n: int) -> bool:
        squares = [i**2 for i in range(10)]
        seen = set()

        while n != 1:
            n = sum(squares[int(i)] for i in str(n))
            if n in seen:
                return False
            seen.add(n)

        return True


run_tests(
    Solution().isHappy,
    [
        {"input": [19], "expected": True},
        {"input": [2], "expected": False},
        {"input": [1111111], "expected": True},
    ],
)
