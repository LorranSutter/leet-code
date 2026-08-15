package main

import (
	"leetcode/utils"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0

	for _, candy := range candies {
		if candy > maxCandy {
			maxCandy = candy
		}
	}

	greatest := make([]bool, len(candies))

	for i := range candies {
		if candies[i]+extraCandies >= maxCandy {
			greatest[i] = true
		}
	}

	return greatest
}

func main() {
	utils.RunTests([]utils.TestCase[[]bool]{
		{Input: []any{[]int{2, 3, 5, 1, 3}, 3}, Got: kidsWithCandies([]int{2, 3, 5, 1, 3}, 3), Expected: []bool{true, true, true, false, true}},
		{Input: []any{[]int{4, 2, 1, 1, 2}, 1}, Got: kidsWithCandies([]int{4, 2, 1, 1, 2}, 1), Expected: []bool{true, false, false, false, false}},
		{Input: []any{[]int{12, 1, 12}, 10}, Got: kidsWithCandies([]int{12, 1, 12}, 10), Expected: []bool{true, false, true}},
	})
}
