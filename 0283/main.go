package main

import "leetcode/utils"

// 0 ms / 9.43 MB
func moveZeroes1(nums []int) []int {
	p1, p2 := 0, 1

	for p2 < len(nums) {
		if nums[p1] == 0 {
			if nums[p2] != 0 {
				nums[p1], nums[p2] = nums[p2], nums[p1]
				p1++
				p2++
			} else {
				p2++
			}
		} else {
			p1++
			if p1 >= p2 {
				p2 = p1 + 1
			}
		}
	}

	return nums
}

// 48 ms / 8.82 MB
func moveZeroes2(nums []int) []int {
	for i := range nums {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}

	return nums
}

func main() {
	utils.RunTests([]utils.TestCase[[]int]{
		{Input: []int{0, 1, 0, 3, 12}, Got: moveZeroes1([]int{0, 1, 0, 3, 12}), Expected: []int{1, 3, 12, 0, 0}},
		{Input: []int{1, 0, 0, 3, 12}, Got: moveZeroes1([]int{1, 0, 0, 3, 12}), Expected: []int{1, 3, 12, 0, 0}},
		{Input: []int{0, 0, 1, 3, 12}, Got: moveZeroes1([]int{0, 0, 1, 3, 12}), Expected: []int{1, 3, 12, 0, 0}},
		{Input: []int{1, 3, 12, 0, 0}, Got: moveZeroes1([]int{1, 3, 12, 0, 0}), Expected: []int{1, 3, 12, 0, 0}},
		{Input: []int{0, 0, 1, 3, 12, 0, 0}, Got: moveZeroes1([]int{0, 0, 1, 3, 12, 0, 0}), Expected: []int{1, 3, 12, 0, 0, 0, 0}},
		{Input: []int{0, 0, 1, 0, 3, 0, 0, 12, 13, 34, 55, 0, 0}, Got: moveZeroes1([]int{0, 0, 1, 0, 3, 0, 0, 12, 13, 34, 55, 0, 0}), Expected: []int{1, 3, 12, 13, 34, 55, 0, 0, 0, 0, 0, 0, 0}},
		{Input: []int{0}, Got: moveZeroes1([]int{0}), Expected: []int{0}},
		{Input: []int{1}, Got: moveZeroes1([]int{1}), Expected: []int{1}},
	})

	utils.RunTests([]utils.TestCase[[]int]{
		{Input: []int{0, 1, 0, 3, 12}, Got: moveZeroes2([]int{0, 1, 0, 3, 12}), Expected: []int{1, 3, 12, 0, 0}},
		{Input: []int{1, 0, 0, 3, 12}, Got: moveZeroes2([]int{1, 0, 0, 3, 12}), Expected: []int{1, 3, 12, 0, 0}},
		{Input: []int{0, 0, 1, 3, 12}, Got: moveZeroes2([]int{0, 0, 1, 3, 12}), Expected: []int{1, 3, 12, 0, 0}},
		{Input: []int{1, 3, 12, 0, 0}, Got: moveZeroes2([]int{1, 3, 12, 0, 0}), Expected: []int{1, 3, 12, 0, 0}},
		{Input: []int{0, 0, 1, 3, 12, 0, 0}, Got: moveZeroes2([]int{0, 0, 1, 3, 12, 0, 0}), Expected: []int{1, 3, 12, 0, 0, 0, 0}},
		{Input: []int{0, 0, 1, 0, 3, 0, 0, 12, 13, 34, 55, 0, 0}, Got: moveZeroes2([]int{0, 0, 1, 0, 3, 0, 0, 12, 13, 34, 55, 0, 0}), Expected: []int{1, 3, 12, 13, 34, 55, 0, 0, 0, 0, 0, 0, 0}},
		{Input: []int{0}, Got: moveZeroes2([]int{0}), Expected: []int{0}},
		{Input: []int{1}, Got: moveZeroes2([]int{1}), Expected: []int{1}},
	})
}
