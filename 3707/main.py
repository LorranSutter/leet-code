import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def scoreBalance(self, s: str) -> bool:
        s = [ord(letter) - 96 for letter in s]
        l = len(s)

        sum_init, sum_end = s[0], s[-1]
        i, j = 0, l - 1
        while True:
            if j - i == 1:
                return sum_init == sum_end
            if sum_init > sum_end:
                j -= 1
                sum_end += s[j]
            else:
                i += 1
                sum_init += s[i]


run_tests(Solution().scoreBalance, [
    {"input": ["adcb"], "expected": True},
    {"input": ["bace"], "expected": False},
    {"input": ["baece"], "expected": True},
])
