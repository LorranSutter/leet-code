import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        s = set(nums)

        longest = 0
        for num in s:
            if num - 1 in s:
                continue

            next_num = num + 1
            length = 1
            while next_num in s:
                next_num += 1
                length += 1
            longest = max(longest, length)

        return longest


run_tests(
    Solution().longestConsecutive,
    [
        {"input": [[100, 4, 200, 1, 3, 2]], "expected": 4},
        {"input": [[0, 3, 7, 2, 5, 8, 4, 6, 0, 1]], "expected": 9},
        {"input": [[1, 0, 1, 2]], "expected": 3},
    ],
)
