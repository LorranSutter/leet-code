import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def maxSubsequence(self, nums: List[int], k: int) -> List[int]:
        # Create a list to keep track of the original index after sorted
        # [2, 1, 3, 3] -> [(3, 2), (3, 3), (2, 0), (1, 1)]
        nums = [(nums[i], i) for i in range(len(nums))]
        nums.sort(reverse=True, key=lambda x: x[0])

        # Get the kth elements
        result = nums[:k]

        # Sort the result by the original index
        result.sort(key=lambda x: x[1])
        result = [r[0] for r in result]

        return result


run_tests(
    Solution().maxSubsequence,
    [
        {"input": [[2, 1, 3, 3], 2], "expected": [3, 3]},
        {"input": [[-1, -2, 3, 4], 3], "expected": [-1, 3, 4]},
        {"input": [[3, 4, 3, 3], 2], "expected": [3, 4]},
    ],
)
