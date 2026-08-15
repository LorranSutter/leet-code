import sys
from pathlib import Path
from typing import List

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        fib_sum = nums[0] + nums[1]
        max_subarray, current_max = 2, 2

        for i in range(2, len(nums)):
            print(fib_sum, [nums[i-1],nums[i]], i)
            if fib_sum == nums[i]:
                current_max += 1
            else:
                if current_max > max_subarray:
                    max_subarray = current_max
                current_max = 2
            fib_sum = nums[i - 1] + nums[i]

        if current_max > max_subarray:
            max_subarray = current_max
        print(max_subarray)
        return max_subarray


run_tests(Solution().longestSubarray, [
    {"input": [[1, 1, 1, 1, 2, 3, 5, 1]], "expected": 5},
    {"input": [[5, 2, 7, 9, 16]], "expected": 5},
    {"input": [[1000000000, 1000000000, 1000000000]], "expected": 2},
    {"input": [[3, 1, 4, 5, 3, 1, 4, 3, 1, 4]], "expected": 4},
])
