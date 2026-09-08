import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


# Boyer-Moore Voting Algorithm
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        candidate = None
        count = 0

        for num in nums:
            if count == 0:
                candidate = num

            if candidate == num:
                count += 1
            else:
                count -= 1

        return candidate


run_tests(
    Solution().majorityElement,
    [
        {"input": [[3, 2, 3]], "expected": 3},
        {"input": [[2, 2, 1, 1, 1, 2, 2]], "expected": 2},
    ],
)
