import sys
from typing import List
from pathlib import Path
from functools import cache

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def hasValidPath(self, grid: List[List[str]]) -> bool:
        m, n = len(grid), len(grid[0])

        if grid[0][0] != "(" or grid[m - 1][n - 1] != ")":
            return False

        @cache
        def dfs(x: int, y: int, count: int) -> bool:
            if x == m - 1 and y == n - 1:
                return count == 1

            if grid[x][y] == "(":
                if x < m - 1 and dfs(x + 1, y, count + 1):
                    return True
                if y < n - 1 and dfs(x, y + 1, count + 1):
                    return True
            else:
                if count == 0:
                    return False
                if x < m - 1 and dfs(x + 1, y, count - 1):
                    return True
                if y < n - 1 and dfs(x, y + 1, count - 1):
                    return True

            return False

        return dfs(0, 0, 0)


run_tests(
    Solution().hasValidPath,
    [
        {
            "input": [
                [["(", "(", "("], [")", "(", ")"], ["(", "(", ")"], ["(", "(", ")"]]
            ],
            "expected": True,
        },
        {"input": [[[")", ")"], ["(", "("]]], "expected": False},
    ],
)
