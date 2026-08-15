package main

import "leetcode/utils"

// The idea is to keep track of the number of ones
// Every time we encounter a zero after a one, we add the count of ones to the maximum number of operations

func maxOperations(s string) int {
	maxOp, ones := 0, 0
	newZero := false

	for _, c := range s {
		if c == '1' {
			ones++
			newZero = false
		} else if !newZero {
			maxOp += ones
			newZero = true
		}
	}

	return maxOp
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: "1001101", Got: maxOperations("1001101"), Expected: 4},
		{Input: "00111", Got: maxOperations("00111"), Expected: 0},
		{Input: "10011010", Got: maxOperations("10011010"), Expected: 8},
	})
}
