import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        m = len(cardPoints) - k
        total = sum(cardPoints)
        sub_total = sum(cardPoints[:m])
        score = total - sub_total

        for i in range(m, len(cardPoints)):
            sub_total += cardPoints[i]
            sub_total -= cardPoints[i - m]

            score = max(score, total - sub_total)

        return score


run_tests(
    Solution().maxScore,
    [
        {"input": [[1, 2, 3, 4, 5, 6, 1], 3], "expected": 12},
        {"input": [[2, 2, 2], 2], "expected": 4},
        {"input": [[9, 7, 7, 9, 7, 7, 9], 7], "expected": 55},
        {"input": [[11, 49, 100, 20, 86, 29, 72], 4], "expected": 232},
    ],
)
