package main

import (
	"leetcode/utils"
)

// 0 ms / 5.01 MB
func missingMultiple(nums []int, k int) int {
	setNums := make(map[int]bool, len(nums))
	for _, num := range nums {
		setNums[num] = true
	}

	mul := k
	for {
		if !setNums[mul] {
			return mul
		}
		mul += k
	}
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{[]int{8, 2, 3, 4, 6}, 2}, Got: missingMultiple([]int{8, 2, 3, 4, 6}, 2), Expected: 10},
		{Input: []any{[]int{1, 4, 7, 10, 15}, 5}, Got: missingMultiple([]int{1, 4, 7, 10, 15}, 5), Expected: 5},
	})
}
