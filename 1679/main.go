package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	fmt.Println(nums)

	start, end := 0, len(nums)-1
	result := 0
	for start < end {
		if nums[start] >= k {
			break
		}
		if nums[end] >= k {
			end--
			continue
		}
		sum := nums[start] + nums[end]
		if sum == k {
			result++
			start++
			end--
		} else if sum > k {
			end--
		} else {
			start++
		}
	}

	return result
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5) == 2)
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6) == 1)
	fmt.Println(maxOperations([]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3) == 4)
	fmt.Println(maxOperations([]int{3, 1, 5, 1, 1, 1, 1, 1, 2, 2, 3, 2, 2}, 1) == 0)
	fmt.Println(maxOperations([]int{2, 2, 2, 3, 1, 1, 4, 1}, 3) == 3)
}
