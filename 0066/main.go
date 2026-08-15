package main

import (
	"leetcode/utils"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		newDigit := digits[i] + 1
		if newDigit <= 9 {
			digits[i] = newDigit
			return digits
		}

		digits[i] = 0

		if i == 0 {
			return append([]int{1}, digits...)
		}
	}

	return digits
}

func main() {
	utils.RunTests([]utils.TestCase[[]int]{
		{Input: []int{1, 2, 3}, Got: plusOne([]int{1, 2, 3}), Expected: []int{1, 2, 4}},
		{Input: []int{4, 3, 2, 1}, Got: plusOne([]int{4, 3, 2, 1}), Expected: []int{4, 3, 2, 2}},
		{Input: []int{9}, Got: plusOne([]int{9}), Expected: []int{1, 0}},
		{Input: []int{8, 9, 9, 9}, Got: plusOne([]int{8, 9, 9, 9}), Expected: []int{9, 0, 0, 0}},
	})
}
