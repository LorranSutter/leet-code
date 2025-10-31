package main

import (
	"fmt"
	"leetcode/utils"
)

func getSneakyNumbers(nums []int) []int {
	mapping := make(map[int]bool)
	sneakies := []int{}
	for _, num := range nums {
		if mapping[num] {
			sneakies = append(sneakies, num)
			if len(sneakies) > 1 {
				break
			}
		} else {
			mapping[num] = true
		}
	}

	return sneakies
}

func main() {
	fmt.Println(utils.EqualSlices(getSneakyNumbers([]int{0, 1, 1, 0}), []int{0, 1}))
	fmt.Println(utils.EqualSlices(getSneakyNumbers([]int{63, 1, 1, 63}), []int{63, 1}))
	fmt.Println(utils.EqualSlices(getSneakyNumbers([]int{0, 3, 2, 1, 3, 2}), []int{2, 3}))
	fmt.Println(utils.EqualSlices(getSneakyNumbers([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}), []int{4, 5}))
	fmt.Println(utils.EqualSlices(getSneakyNumbers([]int{5, 35, 41, 59, 56, 30, 58, 43, 39, 52, 25, 49, 50, 23, 63, 66, 22, 3, 51, 47, 2, 19, 11, 6, 12, 29, 33, 26, 4, 17, 15, 38, 55, 36, 28, 61, 34, 42, 37, 40, 21, 0, 48, 65, 7, 44, 63, 31, 57, 54, 9, 13, 10, 18, 46, 1, 1, 24, 62, 8, 45, 16, 53, 60, 20, 14, 32, 27, 64}), []int{1, 63}))
}
