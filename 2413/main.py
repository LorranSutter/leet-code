import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def smallestEvenMultiple(self, n: int) -> int:
        if n % 2 == 0:
            return n
        return n * 2


run_tests(
    Solution().smallestEvenMultiple,
    [
        {"input": [5], "expected": 10},
        {"input": [6], "expected": 6},
    ],
)
