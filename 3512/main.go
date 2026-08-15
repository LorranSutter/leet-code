package main

import "leetcode/utils"

func minOperations(nums []int, k int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum % k
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{[]int{3, 9, 7}, 5}, Got: minOperations([]int{3, 9, 7}, 5), Expected: 4},
		{Input: []any{[]int{4, 1, 3}, 4}, Got: minOperations([]int{4, 1, 3}, 4), Expected: 0},
		{Input: []any{[]int{3, 2}, 6}, Got: minOperations([]int{3, 2}, 6), Expected: 5},
	})
}
