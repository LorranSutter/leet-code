package main

import (
	"leetcode/utils"
)

// 0 ms / 4.83 MB
func searchInsert1(nums []int, target int) int {
	left, mid := 0, 0
	right := len(nums) - 1

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if nums[mid] > target {
		return mid
	}
	return mid + 1
}

// 0 ms / 4.72 MB
func searchInsert2(nums []int, target int) int {
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{[]int{1, 3, 5, 6}, 5}, Got: searchInsert1([]int{1, 3, 5, 6}, 5), Expected: 2},
		{Input: []any{[]int{1, 3, 5, 6}, 2}, Got: searchInsert1([]int{1, 3, 5, 6}, 2), Expected: 1},
		{Input: []any{[]int{1, 3, 5, 6}, 7}, Got: searchInsert1([]int{1, 3, 5, 6}, 7), Expected: 4},
		{Input: []any{[]int{1, 3, 5, 6}, 0}, Got: searchInsert1([]int{1, 3, 5, 6}, 0), Expected: 0},
	})

	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{[]int{1, 3, 5, 6}, 5}, Got: searchInsert2([]int{1, 3, 5, 6}, 5), Expected: 2},
		{Input: []any{[]int{1, 3, 5, 6}, 2}, Got: searchInsert2([]int{1, 3, 5, 6}, 2), Expected: 1},
		{Input: []any{[]int{1, 3, 5, 6}, 7}, Got: searchInsert2([]int{1, 3, 5, 6}, 7), Expected: 4},
		{Input: []any{[]int{1, 3, 5, 6}, 0}, Got: searchInsert2([]int{1, 3, 5, 6}, 0), Expected: 0},
	})
}
