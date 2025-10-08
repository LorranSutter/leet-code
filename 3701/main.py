from typing import List


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
