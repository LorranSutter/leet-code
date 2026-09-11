import sys
from typing import List
from pathlib import Path
from itertools import permutations

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        nums = set(permutations(digits, 3))

        count = 0
        for num in nums:
            if num[0] == 0:
                continue
            if num[-1] % 2 != 0:
                continue
            count += 1

        return count


run_tests(
    Solution().totalNumbers,
    [
        {"input": [[1, 2, 3, 4]], "expected": 12},
        {"input": [[0, 2, 2]], "expected": 2},
        {"input": [[6, 6, 6]], "expected": 1},
        {"input": [[1, 3, 5]], "expected": 0},
    ],
)
