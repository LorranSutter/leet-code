package main

import (
	"leetcode/utils"
)

func largestAltitude(gain []int) int {
	highest, current := 0, 0

	for _, g := range gain {
		current += g
		if current > highest {
			highest = current
		}
	}

	return highest
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{-5, 1, 5, 0, -7}, Got: largestAltitude([]int{-5, 1, 5, 0, -7}), Expected: 1},
		{Input: []int{-4, -3, -2, -1, 4, 3, 2}, Got: largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}), Expected: 0},
	})
}
