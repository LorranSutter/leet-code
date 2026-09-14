import sys
from typing import List
from pathlib import Path
from functools import cmp_to_key

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def largestNumber(self, nums: List[int]) -> str:
        # Greedy approach comparing each possible concatenation
        # Check if num1 + num2 is greater than num2 + num1
        # E.g: [2,10] -> 2 + 10 > 10 + 2
        def custom_compare(num1: str, num2: str) -> int:
            n1n2, n2n1 = int(num1 + num2), int(num2 + num1)
            if n1n2 > n2n1:
                return -1
            elif n1n2 < n2n1:
                return 1
            return 0

        nums = sorted(map(str, nums), key=cmp_to_key(custom_compare))

        # If the largest number is 0, it means everything else is 0
        if nums[0] == "0":
            return "0"

        return "".join(nums)


run_tests(
    Solution().largestNumber,
    [
        {"input": [[10, 2]], "expected": "210"},
        {"input": [[3, 30, 34, 5, 9]], "expected": "9534330"},
        {"input": [[999, 9998, 9990]], "expected": "99999989990"},
    ],
)
