package main

import "leetcode/utils"

func findMaxConsecutiveOnes(nums []int) int {
	maxCount, count := 0, 0

	for i := range nums {
		if nums[i] == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return maxCount
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{1, 1, 0, 1, 1, 1}, Got: findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}), Expected: 3},
		{Input: []int{1, 0, 1, 1, 0, 1}, Got: findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}), Expected: 2},
	})
}
