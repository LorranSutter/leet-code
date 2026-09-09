import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def sumAndMultiply(self, n: int) -> int:
        x, s = "", 0

        while n > 0:
            digit = n % 10
            if digit != 0:
                x += str(digit)
                s += digit
            n = n // 10

        if x == "":
            return 0
        return int(x[::-1]) * s


run_tests(
    Solution().sumAndMultiply,
    [
        {"input": [10203004], "expected": 12340},
        {"input": [1000], "expected": 1},
        {"input": [1000000000], "expected": 1},
        {"input": [505050505], "expected": 1388875},
        {"input": [0], "expected": 0},
    ],
)
