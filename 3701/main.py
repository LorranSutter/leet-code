import sys
from pathlib import Path
from typing import List

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests


class Solution:
    # 3 ms / 18.12 MB
    def alternatingSum1(self, nums: List[int]) -> int:
        res = 0
        for i in range(len(nums)):
            res += nums[i] if i % 2 == 0 else -nums[i]
        return res

    # 3 ms / 17.89 ms
    def alternatingSum2(self, nums: List[int]) -> int:
        res = sum((nums[i] for i in range(0, len(nums), 2)))
        res += sum((-nums[i] for i in range(1, len(nums), 2)))
        return res

    # 4 ms / 17.82 ms
    def alternatingSum3(self, nums: List[int]) -> int:
        res = 0
        for i, num in enumerate(nums):
            res += (-1) ** i * num
        return res


test_cases = [
    {"input": [[]], "expected": 0},
    {"input": [[5]], "expected": 5},
    {"input": [[1, 2, 3, 4]], "expected": -2},
    {"input": [[3, 7, 1]], "expected": -3},
]

run_tests(Solution().alternatingSum1, test_cases)
run_tests(Solution().alternatingSum2, test_cases)
run_tests(Solution().alternatingSum3, test_cases)
