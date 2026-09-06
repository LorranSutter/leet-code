import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def stableMountains(self, height: List[int], threshold: int) -> List[int]:
        stable = []
        for i in range(1, len(height)):
            if height[i - 1] > threshold:
                stable.append(i)

        return stable


run_tests(
    Solution().stableMountains,
    [
        {"input": [[1, 2, 3, 4, 5], 2], "expected": [3, 4]},
        {"input": [[10, 1, 10, 1, 10], 3], "expected": [1, 3]},
        {"input": [[10, 1, 10, 1, 10], 10], "expected": []},
    ],
)
