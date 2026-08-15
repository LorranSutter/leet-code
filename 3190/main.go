package main

import "leetcode/utils"

func minimumOperations(nums []int) int {
	count := 0
	for i := range nums {
		if nums[i]%3 != 0 {
			count++
		}
	}

	return count
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{1, 2, 3, 4}, Got: minimumOperations([]int{1, 2, 3, 4}), Expected: 3},
		{Input: []int{3, 6, 9}, Got: minimumOperations([]int{3, 6, 9}), Expected: 0},
	})
}
