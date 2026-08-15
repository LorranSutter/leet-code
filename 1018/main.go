package main

import (
	"leetcode/utils"
)

// The idea here is to keep track only of the modulus 5 of the current number
// To simplify the addition of a new bit, we can use bit shifting and add the new bit

func prefixesDivBy5(nums []int) []bool {
	answer := make([]bool, len(nums))
	num := 0

	for i := range nums {
		// Shifting by 1 is the same as multiplying by 2
		// num*2 == num << 1
		num = (num<<1 | nums[i]) % 5
		answer[i] = num == 0
	}

	return answer
}

func main() {
	utils.RunTests([]utils.TestCase[[]bool]{
		{Input: []int{0, 1, 1}, Got: prefixesDivBy5([]int{0, 1, 1}), Expected: []bool{true, false, false}},
		{Input: []int{1, 1, 1}, Got: prefixesDivBy5([]int{1, 1, 1}), Expected: []bool{false, false, false}},
		{Input: []int{0, 1, 1, 1, 1, 1}, Got: prefixesDivBy5([]int{0, 1, 1, 1, 1, 1}), Expected: []bool{true, false, false, false, true, false}},
	})
}
