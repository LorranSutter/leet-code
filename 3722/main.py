import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def lexSmallest(self, s: str) -> str:
        smallest = s

        for i in range(len(s)):
            s_revers_front = s[:i][::-1] + s[i:]
            if s_revers_front < smallest:
                smallest = s_revers_front
            s_revers_back = s[:-i] + s[-i:][::-1]
            if s_revers_back < smallest:
                smallest = s_revers_back

        return smallest


run_tests(Solution().lexSmallest, [
    {"input": ["dcab"], "expected": "acdb"},
    {"input": ["abba"], "expected": "aabb"},
    {"input": ["zxy"], "expected": "xzy"},
    {"input": ["lk"], "expected": "kl"},
])
