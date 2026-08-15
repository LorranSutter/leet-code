package main

import (
	"leetcode/utils"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	start, end := 0, len(nums)-1
	result := 0
	for start < end {
		if nums[start] >= k {
			break
		}
		if nums[end] >= k {
			end--
			continue
		}
		sum := nums[start] + nums[end]
		if sum == k {
			result++
			start++
			end--
		} else if sum > k {
			end--
		} else {
			start++
		}
	}

	return result
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{[]int{1, 2, 3, 4}, 5}, Got: maxOperations([]int{1, 2, 3, 4}, 5), Expected: 2},
		{Input: []any{[]int{3, 1, 3, 4, 3}, 6}, Got: maxOperations([]int{3, 1, 3, 4, 3}, 6), Expected: 1},
		{Input: []any{[]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3}, Got: maxOperations([]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3), Expected: 4},
		{Input: []any{[]int{3, 1, 5, 1, 1, 1, 1, 1, 2, 2, 3, 2, 2}, 1}, Got: maxOperations([]int{3, 1, 5, 1, 1, 1, 1, 1, 2, 2, 3, 2, 2}, 1), Expected: 0},
		{Input: []any{[]int{2, 2, 2, 3, 1, 1, 4, 1}, 3}, Got: maxOperations([]int{2, 2, 2, 3, 1, 1, 4, 1}, 3), Expected: 3},
	})
}
