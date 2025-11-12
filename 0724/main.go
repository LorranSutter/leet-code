package main

import "fmt"

// 3 ms / 8.34 MB
// Simple prefix sum
func pivotIndex1(nums []int) int {
	prefix := make([]int, len(nums)+1)
	for i := range nums {
		prefix[i+1] = prefix[i] + nums[i]
	}

	total_sum := prefix[len(nums)]
	for i := range nums {
		if prefix[i] == total_sum-prefix[i+1] {
			return i
		}
	}

	return -1
}

// 1 ms / 7.78 MB
func pivotIndex2(nums []int) int {
	right_sum := 0
	for i := range nums {
		right_sum += nums[i]
	}

	left_sum := 0
	for i := range nums {
		right_sum -= nums[i]
		if left_sum == right_sum {
			return i
		}
		left_sum += nums[i]
	}

	return -1
}

func main() {
	fmt.Println("Solution 01")
	fmt.Println(pivotIndex1([]int{1, 7, 3, 6, 5, 6}) == 3)
	fmt.Println(pivotIndex1([]int{1, 2, 3}) == -1)
	fmt.Println(pivotIndex1([]int{2, 1, -1}) == 0)

	fmt.Println()

	fmt.Println("Solution 02")
	fmt.Println(pivotIndex2([]int{1, 7, 3, 6, 5, 6}) == 3)
	fmt.Println(pivotIndex2([]int{1, 2, 3}) == -1)
	fmt.Println(pivotIndex2([]int{2, 1, -1}) == 0)
}
