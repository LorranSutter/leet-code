import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def createGrid(self, m: int, n: int) -> list[str]:
        grid = []
        for _ in range(m - 1):
            grid.append("." + "#" * (n - 1))
        grid.append("." * n)

        return grid


run_tests(
    Solution().createGrid,
    [
        {"input": [2, 3], "expected": [".##", "..."]},
        {"input": [3, 3], "expected": [".##", ".##", "..."]},
        {"input": [1, 4], "expected": ["...."]},
        {"input": [3, 1], "expected": [".", ".", "."]},
    ],
)
