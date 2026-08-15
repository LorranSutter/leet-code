package main

import "leetcode/utils"

func smallestNumber(n int) int {
	smallest := 1
	for smallest < n {
		smallest = smallest<<1 | 1
	}

	return smallest
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: 5, Got: smallestNumber(5), Expected: 7},
		{Input: 10, Got: smallestNumber(10), Expected: 15},
		{Input: 3, Got: smallestNumber(3), Expected: 3},
		{Input: 1000, Got: smallestNumber(1000), Expected: 1023},
	})
}
