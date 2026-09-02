import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        
        x1 = x
        x2 = 0
        while x1 > 0:
            x2 = x2*10 + x1 % 10
            x1 = x1 // 10

        return x == x2

run_tests(Solution().isPalindrome, [
    {"input": [121], "expected": True},
    {"input": [-121], "expected": False},
    {"input": [2147483647], "expected": False},
    {"input": [0], "expected": True},
])
