package main

import (
	"leetcode/utils"
)

func countOperations(num1 int, num2 int) int {
	count := 0
	for num1*num2 != 0 {
		if num1 < num2 {
			num2 = num2 - num1
		} else {
			num1 = num1 - num2
		}
		count++
	}

	return count
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{2, 3}, Got: countOperations(2, 3), Expected: 3},
		{Input: []any{10, 10}, Got: countOperations(10, 10), Expected: 1},
		{Input: []any{5, 2}, Got: countOperations(5, 2), Expected: 4},
	})
}
