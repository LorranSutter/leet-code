package main

import (
	"leetcode/utils"
)

func minCost(colors string, neededTime []int) int {
	cost := 0
	current_color := colors[0]
	chunk_cost := neededTime[0]
	max_chunk_cost := neededTime[0]
	for i := 1; i < len(neededTime); i++ {
		if current_color != colors[i] {
			cost += (chunk_cost - max_chunk_cost)
			chunk_cost, max_chunk_cost = neededTime[i], neededTime[i]
			current_color = colors[i]
			continue
		}

		chunk_cost += neededTime[i]
		if neededTime[i] > max_chunk_cost {
			max_chunk_cost = neededTime[i]
		}
	}

	cost += (chunk_cost - max_chunk_cost)

	return cost
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{"abaac", []int{1, 2, 3, 4, 5}}, Got: minCost("abaac", []int{1, 2, 3, 4, 5}), Expected: 3},
		{Input: []any{"abc", []int{1, 2, 3}}, Got: minCost("abc", []int{1, 2, 3}), Expected: 0},
		{Input: []any{"aabaa", []int{1, 2, 3, 4, 1}}, Got: minCost("aabaa", []int{1, 2, 3, 4, 1}), Expected: 2},
		{Input: []any{"baaaabaa", []int{1, 1, 5, 1, 2, 6, 7, 1}}, Got: minCost("baaaabaa", []int{1, 1, 5, 1, 2, 6, 7, 1}), Expected: 5},
		{Input: []any{"bbbaaa", []int{4, 9, 3, 8, 8, 9}}, Got: minCost("bbbaaa", []int{4, 9, 3, 8, 8, 9}), Expected: 23},
	})
}
