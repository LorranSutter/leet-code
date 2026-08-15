package main

import "leetcode/utils"

func countValidSelections(nums []int) int {
	lenPrefix := len(nums) + 1
	sum_left := make([]int, lenPrefix)
	sum_right := make([]int, lenPrefix)

	for i, j := 0, len(nums); i < len(nums); i, j = i+1, j-1 {
		sum_left[i+1] = sum_left[i] + nums[i]
		sum_right[j-1] = sum_right[j] + nums[j-1]
	}

	total := 0
	for i := 0; i < lenPrefix-1; i++ {
		if nums[i] != 0 {
			continue
		}

		if sum_left[i] == sum_right[i] {
			total += 2
		} else if sum_left[i]-1 == sum_right[i] {
			total += 1
		} else if sum_left[i] == sum_right[i]-1 {
			total += 1
		}
	}

	return total
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{1, 0, 2, 0, 3}, Got: countValidSelections([]int{1, 0, 2, 0, 3}), Expected: 2},
		{Input: []int{1, 0, 2, 0, 4}, Got: countValidSelections([]int{1, 0, 2, 0, 4}), Expected: 1},
		{Input: []int{1, 0, 2, 0, 2}, Got: countValidSelections([]int{1, 0, 2, 0, 2}), Expected: 1},
		{Input: []int{2, 3, 4, 0, 4, 1, 0}, Got: countValidSelections([]int{2, 3, 4, 0, 4, 1, 0}), Expected: 0},
		{Input: []int{0}, Got: countValidSelections([]int{0}), Expected: 2},
		{Input: []int{16, 13, 10, 0, 0, 0, 10, 6, 7, 8, 7}, Got: countValidSelections([]int{16, 13, 10, 0, 0, 0, 10, 6, 7, 8, 7}), Expected: 3},
	})
}
