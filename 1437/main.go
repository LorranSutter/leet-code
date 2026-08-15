package main

import "leetcode/utils"

// Keep track of the last 1's position and check the diff between current and last 1's position

func kLengthApart(nums []int, k int) bool {
	lastOne := -k - 1
	for i, num := range nums {
		if num == 1 {
			if i-lastOne-1 < k {
				return false
			}
			lastOne = i
		}
	}
	return true
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []any{[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2}, Got: kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2), Expected: true},
		{Input: []any{[]int{1, 0, 0, 1, 0, 1}, 2}, Got: kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2), Expected: false},
		{Input: []any{[]int{1, 0, 0, 1, 0, 1}, 0}, Got: kLengthApart([]int{1, 0, 0, 1, 0, 1}, 0), Expected: true},
		{Input: []any{[]int{0, 0, 0, 0, 0, 0}, 5}, Got: kLengthApart([]int{0, 0, 0, 0, 0, 0}, 5), Expected: true},
		{Input: []any{[]int{1, 1, 0, 1, 0, 0}, 1}, Got: kLengthApart([]int{1, 1, 0, 1, 0, 0}, 1), Expected: false},
		{Input: []any{[]int{0, 1, 0, 1, 0, 1}, 1}, Got: kLengthApart([]int{0, 1, 0, 1, 0, 1}, 1), Expected: true},
	})
}
