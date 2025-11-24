package main

import (
	"fmt"
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
	result := prefixesDivBy5([]int{0, 1, 1})
	fmt.Println(utils.DeepEqualSlices(result, []bool{true, false, false}))

	result = prefixesDivBy5([]int{1, 1, 1})
	fmt.Println(utils.DeepEqualSlices(result, []bool{false, false, false}))

	result = prefixesDivBy5([]int{0, 1, 1, 1, 1, 1})
	fmt.Println(utils.DeepEqualSlices(result, []bool{true, false, false, false, true, false}))
}
