import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        s = s.split()
        if len(pattern) != len(s):
            return False

        mapping = dict()
        mapped_s = set()
        for i in range(len(pattern)):
            if pattern[i] not in mapping:
                if s[i] in mapped_s:
                    return False
                mapping[pattern[i]] = s[i]
                mapped_s.add(s[i])
            elif mapping[pattern[i]] != s[i]:
                return False

        return True


run_tests(
    Solution().wordPattern,
    [
        {"input": ["abba", "dog cat cat dog"], "expected": True},
        {"input": ["abba", "dog cat cat fish"], "expected": False},
        {"input": ["aaaa", "dog cat cat"], "expected": False},
        {"input": ["abba", "dog dog dog dog"], "expected": False},
    ],
)
