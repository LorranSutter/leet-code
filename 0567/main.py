import sys
from pathlib import Path
from collections import Counter

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        n1 = len(s1)
        n2 = len(s2)

        if n2 < n1:
            return False

        c1 = Counter(s1)
        c2 = Counter(s2[:n1])

        if c1 == c2:
            return True

        for i in range(1, n2 - n1 + 1):
            c2 = Counter(s2[i : i + n1])
            if c1 == c2:
                return True

        return False


run_tests(
    Solution().checkInclusion,
    [
        {"input": ["ab", "eidbaooo"], "expected": True},
        {"input": ["ab", "eidboaoo"], "expected": False},
        {"input": ["adc", "dcda"], "expected": True},
    ],
)
