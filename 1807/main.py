import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def evaluate(self, s: str, knowledge: list[list[str]]) -> str:
        knowledge = {k[0]: k[1] for k in knowledge}

        result = ""
        i = 0
        while i < len(s):
            if s[i] == "(":
                i += 1
                key = ""
                while s[i] != ")":
                    key += s[i]
                    i += 1

                if key in knowledge.keys():
                    result += knowledge[key]
                else:
                    result += "?"
            else:
                result += s[i]

            i += 1

        return result


run_tests(
    Solution().evaluate,
    [
        {
            "input": ["(name)is(age)yearsold", [["name", "bob"], ["age", "two"]]],
            "expected": "bobistwoyearsold",
        },
        {"input": ["hi(name)", [["a", "b"]]], "expected": "hi?"},
        {"input": ["(a)(a)(a)aaa", [["a", "yes"]]], "expected": "yesyesyesaaa"},
    ],
)
