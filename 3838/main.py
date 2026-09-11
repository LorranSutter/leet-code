import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def mapWordWeights(self, words: List[str], weights: List[int]) -> str:
        result = ""
        for word in words:
            weight = 0
            for c in word:
                idx = ord(c) - ord("a")
                weight += weights[idx]
            result += chr(ord("z") - weight % 26)

        return result


run_tests(
    Solution().mapWordWeights,
    [
        {
            "input": [
                ["abcd", "def", "xyz"],
                [
                    5,
                    3,
                    12,
                    14,
                    1,
                    2,
                    3,
                    2,
                    10,
                    6,
                    6,
                    9,
                    7,
                    8,
                    7,
                    10,
                    8,
                    9,
                    6,
                    9,
                    9,
                    8,
                    3,
                    7,
                    7,
                    2,
                ],
            ],
            "expected": "rij",
        },
        {
            "input": [
                ["a", "b", "c"],
                [
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                    1,
                ],
            ],
            "expected": "yyy",
        },
        {
            "input": [
                ["abcd"],
                [
                    7,
                    5,
                    3,
                    4,
                    3,
                    5,
                    4,
                    9,
                    4,
                    2,
                    2,
                    7,
                    10,
                    2,
                    5,
                    10,
                    6,
                    1,
                    2,
                    2,
                    4,
                    1,
                    3,
                    4,
                    4,
                    5,
                ],
            ],
            "expected": "g",
        },
    ],
)
