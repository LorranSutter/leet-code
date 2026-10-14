import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def maxSumOfSquares(self, num: int, s: int) -> str:
        if s > 9 * num:
            return ""

        result = ""
        for _ in range(num):
            digit = min(9, s)
            result += str(digit)
            s -= digit

        return result


run_tests(
    Solution().maxSumOfSquares,
    [
        {"input": [2, 3], "expected": "30"},
        {"input": [2, 17], "expected": "98"},
        {"input": [1, 10], "expected": ""},
    ],
)
