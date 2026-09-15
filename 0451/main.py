import sys
from pathlib import Path
from collections import Counter

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def frequencySort(self, s: str) -> str:
        s = Counter(s)
        return "".join(c * f for c, f in s.most_common())


run_tests(
    Solution().frequencySort,
    [
        {"input": ["tree"], "expected": "eetr"},
        {"input": ["cccaaa"], "expected": "cccaaa"},
        {"input": ["Aabb"], "expected": "bbAa"},
    ],
)
