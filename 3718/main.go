package main

import (
	"fmt"
)

// 0 ms / 5.01 MB
func missingMultiple(nums []int, k int) int {
	setNums := make(map[int]bool, len(nums))
	for _, num := range nums {
		setNums[num] = true
	}

	mul := k
	for {
		if !setNums[mul] {
			return mul
		}
		mul += k
	}
}

func main() {
	fmt.Println(missingMultiple([]int{8, 2, 3, 4, 6}, 2) == 10)
	fmt.Println(missingMultiple([]int{1, 4, 7, 10, 15}, 5) == 5)
}
