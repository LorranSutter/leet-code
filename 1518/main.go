package main

import "leetcode/utils"

func numWaterBottles(numBottles int, numExchange int) int {
	// TODO Implement solution
	// Off-by-one bug
	total := numBottles
	numNewBottles := numBottles

	for numNewBottles > numExchange {
		numNewBottles = numBottles / numExchange
		numBottles = numNewBottles + numBottles%numExchange
		total += numNewBottles
	}

	return total
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{9, 3}, Got: numWaterBottles(9, 3), Expected: 13},
		{Input: []any{15, 4}, Got: numWaterBottles(15, 4), Expected: 19},
	})
}
