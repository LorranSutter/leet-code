import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"

        num1 = list(map(int, num1))[::-1]
        num2 = list(map(int, num2))[::-1]
        result = [0] * (len(num1) + len(num2))

        for i in range(len(num1)):
            for j in range(len(num2)):
                partial = num1[i] * num2[j] + result[i + j]
                result[i + j] = partial % 10
                result[i + j + 1] += partial // 10

        if result[-1] == 0:
            result = result[:-1]

        return "".join(map(str, result))[::-1]


run_tests(
    Solution().multiply,
    [
        {"input": ["2", "3"], "expected": "6"},
        {"input": ["123", "456"], "expected": "56088"},
        {"input": ["9133", "0"], "expected": "0"},
    ],
)
