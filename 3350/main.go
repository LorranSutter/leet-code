package main

import "leetcode/utils"

func maxIncreasingSubarrays(nums []int) int {
	for k := len(nums) / 2; k > 1; k-- {
		size_subarray := 1
		for i := 1; i < len(nums)-k; i++ {
			if nums[i] > nums[i-1] && nums[i+k] > nums[i+k-1] {
				size_subarray++
				if size_subarray == k {
					return k
				}
			} else if size_subarray != 1 {
				size_subarray = 1
			}
		}
	}

	return 1
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, Got: maxIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}), Expected: 3},
		{Input: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, Got: maxIncreasingSubarrays([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}), Expected: 2},
		{Input: []int{0, 4, 16, 20, -6}, Got: maxIncreasingSubarrays([]int{0, 4, 16, 20, -6}), Expected: 2},
		{Input: []int{-15, 19}, Got: maxIncreasingSubarrays([]int{-15, 19}), Expected: 1},
		{Input: []int{5, 8, -2, -1}, Got: maxIncreasingSubarrays([]int{5, 8, -2, -1}), Expected: 2},
	})
}
