package main

import "leetcode/utils"

// 3 ms / 8.34 MB
// Simple prefix sum
func pivotIndex1(nums []int) int {
	prefix := make([]int, len(nums)+1)
	for i := range nums {
		prefix[i+1] = prefix[i] + nums[i]
	}

	total_sum := prefix[len(nums)]
	for i := range nums {
		if prefix[i] == total_sum-prefix[i+1] {
			return i
		}
	}

	return -1
}

// 1 ms / 7.78 MB
func pivotIndex2(nums []int) int {
	right_sum := 0
	for i := range nums {
		right_sum += nums[i]
	}

	left_sum := 0
	for i := range nums {
		right_sum -= nums[i]
		if left_sum == right_sum {
			return i
		}
		left_sum += nums[i]
	}

	return -1
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{1, 7, 3, 6, 5, 6}, Got: pivotIndex1([]int{1, 7, 3, 6, 5, 6}), Expected: 3},
		{Input: []int{1, 2, 3}, Got: pivotIndex1([]int{1, 2, 3}), Expected: -1},
		{Input: []int{2, 1, -1}, Got: pivotIndex1([]int{2, 1, -1}), Expected: 0},
	})

	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{1, 7, 3, 6, 5, 6}, Got: pivotIndex2([]int{1, 7, 3, 6, 5, 6}), Expected: 3},
		{Input: []int{1, 2, 3}, Got: pivotIndex2([]int{1, 2, 3}), Expected: -1},
		{Input: []int{2, 1, -1}, Got: pivotIndex2([]int{2, 1, -1}), Expected: 0},
	})
}
