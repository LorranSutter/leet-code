package main

import (
	"leetcode/utils"
)

func singleNumber(nums []int) int {
	num := 0
	for i := range nums {
		num ^= nums[i]
	}

	return num
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{2, 2, 1}, Got: singleNumber([]int{2, 2, 1}), Expected: 1},
		{Input: []int{4, 1, 2, 1, 2}, Got: singleNumber([]int{4, 1, 2, 1, 2}), Expected: 4},
		{Input: []int{1}, Got: singleNumber([]int{1}), Expected: 1},
	})
}
