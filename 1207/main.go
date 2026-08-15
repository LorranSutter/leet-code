package main

import "leetcode/utils"

func uniqueOccurrences(arr []int) bool {
	freqs := map[int]int{}
	for _, a := range arr {
		freqs[a]++
	}

	count_freq := make([]int, 1001)
	for _, freq := range freqs {
		if count_freq[freq] > 0 {
			return false
		}
		count_freq[freq] = 1
	}

	return true
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []int{1, 2, 2, 1, 1, 3}, Got: uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}), Expected: true},
		{Input: []int{1, 2}, Got: uniqueOccurrences([]int{1, 2}), Expected: false},
		{Input: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, Got: uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}), Expected: true},
		{Input: []int{3, 5, -2, -3, -6, -6}, Got: uniqueOccurrences([]int{3, 5, -2, -3, -6, -6}), Expected: false},
	})
}
