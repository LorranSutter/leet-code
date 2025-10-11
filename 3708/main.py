from typing import List


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


s = Solution()
print(s.longestSubarray([1, 1, 1, 1, 2, 3, 5, 1]) == 5)
print(s.longestSubarray([5, 2, 7, 9, 16]) == 5)
print(s.longestSubarray([1000000000, 1000000000, 1000000000]) == 2)
print(s.longestSubarray([3, 1, 4, 5, 3, 1, 4, 3, 1, 4]) == 4)
