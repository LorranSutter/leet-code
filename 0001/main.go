package main

import (
	"leetcode/utils"
)

func twoSum(nums []int, target int) []int {
	p0 := 0
	p1 := len(nums) - 1
	diff := target - nums[p0]
	for {
		if p0 == p1 {
			p0++
			p1 = len(nums) - 1
			diff = target - nums[p0]
			continue
		}

		if nums[p1] == diff {
			break
		}

		p1--
	}

	return []int{p0, p1}
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []int{2, 7, 11, 15}, Got: utils.DeepEqualSlices(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}), Expected: true},
		{Input: []int{3, 2, 4}, Got: utils.DeepEqualSlices(twoSum([]int{3, 2, 4}, 6), []int{1, 2}), Expected: true},
		{Input: []int{3, 3}, Got: utils.DeepEqualSlices(twoSum([]int{3, 3}, 6), []int{0, 1}), Expected: true},
		{Input: []int{15, 11, 7, 2}, Got: utils.DeepEqualSlices(twoSum([]int{15, 11, 7, 2}, 9), []int{2, 3}), Expected: true},
	})
}
