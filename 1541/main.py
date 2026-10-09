import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def minInsertions(self, s: str) -> int:
        open_count, insertions, i = 0, 0, 0
        while i < len(s):
            if s[i] == "(":
                open_count += 1
            else:
                if open_count > 0:
                    open_count -= 1
                else:
                    # "(" missing
                    insertions += 1

                if i + 1 < len(s):
                    # ")" missing
                    if s[i + 1] != ")":
                        insertions += 1
                    # "))" found
                    else:
                        i += 1
                else:
                    # ")" missing, end of string
                    insertions += 1
            i += 1

        return insertions + open_count * 2


run_tests(
    Solution().minInsertions,
    [
        {"input": ["(()))"], "expected": 1},
        {"input": ["())"], "expected": 0},
        {"input": ["))())("], "expected": 3},
        {"input": [")"], "expected": 2},
        {"input": ["("], "expected": 2},
        {"input": ["(()))(()))()())))"], "expected": 4},
    ],
)
