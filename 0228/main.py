import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        if len(nums) == 0:
            return []

        ranges = []
        new_range = [nums[0], nums[0]]
        for num in nums[1:]:
            if num == new_range[-1] + 1:
                new_range[-1] = num
            else:
                ranges.append(new_range)
                new_range = [num, num]

        ranges.append(new_range)
        for i in range(len(ranges)):
            if ranges[i][0] == ranges[i][1]:
                ranges[i] = str(ranges[i][0])
            else:
                ranges[i] = f"{ranges[i][0]}->{ranges[i][1]}"

        return ranges


run_tests(
    Solution().summaryRanges,
    [
        {"input": [[0, 1, 2, 4, 5, 7]], "expected": ["0->2", "4->5", "7"]},
        {"input": [[0, 2, 3, 4, 6, 8, 9]], "expected": ["0", "2->4", "6", "8->9"]},
        {"input": [[]], "expected": []},
        {"input": [[1]], "expected": ["1"]},
    ],
)
