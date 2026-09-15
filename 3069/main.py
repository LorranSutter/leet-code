import sys
from typing import List
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    def resultArray(self, nums: List[int]) -> List[int]:
        arr1 = [nums[0]]
        arr2 = [nums[1]]

        for element in nums[2:]:
            if arr1[-1] > arr2[-1]:
                arr1.append(element)
            else:
                arr2.append(element)

        return arr1 + arr2


run_tests(
    Solution().resultArray,
    [
        {"input": [[2, 1, 3]], "expected": [2, 3, 1]},
        {"input": [[5, 4, 3, 8]], "expected": [5, 3, 4, 8]},
    ],
)
