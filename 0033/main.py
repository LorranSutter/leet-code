import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def search(self, nums: list[int], target: int) -> int:
        start, end = 0, len(nums) - 1
        while start <= end:
            mid = (start + end) // 2
            if nums[mid] == target:
                return mid
            if nums[start] <= nums[mid]:
                if nums[start] <= target < nums[mid]:
                    end = mid - 1
                else:
                    start = mid + 1
            else:
                if nums[mid] < target <= nums[end]:
                    start = mid + 1
                else:
                    end = mid - 1
        return -1


run_tests(
    Solution().search,
    [
        {"input": [[4, 5, 6, 7, 0, 1, 2], 0], "expected": 4},
        {"input": [[4, 5, 6, 7, 0, 1, 2], 3], "expected": -1},
        {"input": [[1], 0], "expected": -1},
    ],
)
