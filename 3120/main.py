import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        word = set(word)

        specials = 0
        for i in range(26):
            if chr(97 + i) in word and chr(65 + i) in word:
                specials += 1

        return specials


run_tests(
    Solution().numberOfSpecialChars,
    [
        {"input": ["aaAbcBC"], "expected": 3},
        {"input": ["abc"], "expected": 0},
        {"input": ["abBCab"], "expected": 1},
    ],
)
