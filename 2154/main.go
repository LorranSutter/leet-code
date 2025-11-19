package main

import (
	"fmt"
	"sort"
)

// 1 ms / 5.02 MB
func findFinalValue1(nums []int, original int) int {
	sort.Ints(nums)

	for _, num := range nums {
		if num > original {
			break
		}
		if num == original {
			original *= 2
		}
	}
	return original
}

// 0 ms / 5.64 MB
func findFinalValue2(nums []int, original int) int {
	mapNums := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		mapNums[num] = struct{}{}
	}

	for {
		if _, ok := mapNums[original]; !ok {
			break
		}
		original *= 2
	}

	return original
}

func main() {
	fmt.Println("Solution 01")
	fmt.Println(findFinalValue1([]int{5, 3, 6, 1, 12}, 3) == 24)
	fmt.Println(findFinalValue1([]int{2, 7, 9}, 4) == 4)

	fmt.Println()

	fmt.Println("Solution 02")
	fmt.Println(findFinalValue2([]int{5, 3, 6, 1, 12}, 3) == 24)
	fmt.Println(findFinalValue2([]int{2, 7, 9}, 4) == 4)
}
