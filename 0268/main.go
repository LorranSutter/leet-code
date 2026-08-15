package main

import "leetcode/utils"

// 0 ms / 8.8 MB
func missingNumber(nums []int) int {
	hasZero := false
	sum, max := 0, 0
	for _, num := range nums {
		if num == 0 {
			hasZero = true
		} else if num > max {
			max = num
		}
		sum += num
	}

	lenNums := len(nums) + 1
	calcSum := (lenNums*lenNums - lenNums) / 2

	if !hasZero {
		return 0
	}
	if sum == calcSum {
		return max + 1
	}

	return calcSum - sum
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{3, 0, 1}, Got: missingNumber([]int{3, 0, 1}), Expected: 2},
		{Input: []int{0, 1}, Got: missingNumber([]int{0, 1}), Expected: 2},
		{Input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, Got: missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), Expected: 8},
		{Input: []int{2}, Got: missingNumber([]int{2}), Expected: 0},
	})
}
