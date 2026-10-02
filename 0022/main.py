import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        result = []

        def backtrack(left: int, right: int, current: str):
            nonlocal result

            if left > n or right > n:
                return
            if left < right:
                return
            if left == n and right == n:
                result.append(current)
                return

            backtrack(left + 1, right, current + "(")
            backtrack(left, right + 1, current + ")")

        backtrack(0, 0, "")
        return result


run_tests(
    Solution().generateParenthesis,
    [
        {"input": [3], "expected": ["((()))", "(()())", "(())()", "()(())", "()()()"]},
        {"input": [1], "expected": ["()"]},
    ],
)
