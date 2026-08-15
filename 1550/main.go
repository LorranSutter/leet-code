package main

import (
	"leetcode/utils"
)

// 0 ms / 4.56 MB
func threeConsecutiveOdds(arr []int) bool {
	countOdds := 0
	for _, num := range arr {
		if num%2 != 0 {
			countOdds++
			if countOdds >= 3 {
				return true
			}
		} else {
			countOdds = 0
		}
	}
	return false
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []int{2, 6, 4, 1}, Got: threeConsecutiveOdds([]int{2, 6, 4, 1}), Expected: false},
		{Input: []int{1, 2, 34, 3, 4, 5, 7, 23, 12}, Got: threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}), Expected: true},
		{Input: []int{1}, Got: threeConsecutiveOdds([]int{1}), Expected: false},
	})
}
