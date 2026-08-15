package main

import "leetcode/utils"

func countOdds(low int, high int) int {
	if high%2 == 0 && low%2 == 0 {
		return (high - low) / 2
	}

	return (high-low)/2 + 1
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{3, 7}, Got: countOdds(3, 7), Expected: 3},
		{Input: []any{8, 10}, Got: countOdds(8, 10), Expected: 1},
		{Input: []any{7, 10}, Got: countOdds(7, 10), Expected: 2},
		{Input: []any{2, 11}, Got: countOdds(2, 11), Expected: 5},
		{Input: []any{2, 2}, Got: countOdds(2, 2), Expected: 0},
		{Input: []any{7, 7}, Got: countOdds(7, 7), Expected: 1},
	})
}
