import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def checkIfExist(self, arr: List[int]) -> bool:
        seen = set()
        for num in arr:
            if 2 * num in seen or num / 2 in seen:
                return True
            seen.add(num)
        return False


run_tests(
    Solution().checkIfExist,
    [
        {"input": [[10, 2, 5, 3]], "expected": True},
        {"input": [[3, 1, 7, 11]], "expected": False},
    ],
)
