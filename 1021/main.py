import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        count_open = 0
        res = ""
        for char in s:
            if char == "(":
                if count_open != 0:
                    res += char
                count_open += 1
            else:
                if count_open > 1:
                    res += char
                count_open -= 1
        return res


run_tests(
    Solution().removeOuterParentheses,
    [
        {"input": ["(()())(())"], "expected": "()()()"},
        {"input": ["(()())(())(()(()))"], "expected": "()()()()(())"},
        {"input": ["()()"], "expected": ""},
    ],
)
