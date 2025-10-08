from typing import List


class Solution:
    def alternatingSum(self, nums: List[int]) -> int:
        res = 0
        for i in range(len(nums)):
            res += nums[i] if i % 2 == 0 else -nums[i]
        return res
