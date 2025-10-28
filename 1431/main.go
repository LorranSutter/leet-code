package main

import (
	"fmt"
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
	fmt.Println(utils.EqualSlices(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3), []bool{true, true, true, false, true}))
	fmt.Println(utils.EqualSlices(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1), []bool{true, false, false, false, false}))
	fmt.Println(utils.EqualSlices(kidsWithCandies([]int{12, 1, 12}, 10), []bool{true, false, true}))
}
