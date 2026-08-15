package main

import "leetcode/utils"

func containsDuplicate(nums []int) bool {
	dups := make(map[int]bool)

	for _, n := range nums {
		if dups[n] {
			return true
		}
		dups[n] = true
	}

	return false
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []int{1, 2, 3, 1}, Got: containsDuplicate([]int{1, 2, 3, 1}), Expected: true},
		{Input: []int{1, 2, 3, 4}, Got: containsDuplicate([]int{1, 2, 3, 4}), Expected: false},
		{Input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, Got: containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), Expected: true},
	})
}
