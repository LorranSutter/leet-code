import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def minimumFlips(self, n: int) -> int:
        s = str(bin(n))[2:]
        flip_count = 0
        for i in range(len(s)//2):
            if s[i] != s[-i-1]:
                flip_count += 1
        return flip_count*2

run_tests(Solution().minimumFlips, [
    {"input": [7], "expected": 0},
    {"input": [10], "expected": 4},
    {"input": [11], "expected": 2},
    {"input": [17], "expected": 0},
    {"input": [19], "expected": 2},
])