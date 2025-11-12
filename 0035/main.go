package main

import "fmt"

// 0 ms / 4.83 MB
func searchInsert1(nums []int, target int) int {
	left, mid := 0, 0
	right := len(nums) - 1

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if nums[mid] > target {
		return mid
	}
	return mid + 1
}

// 0 ms / 4.72 MB
func searchInsert2(nums []int, target int) int {
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	fmt.Println("Solution 1")
	fmt.Println(searchInsert1([]int{1, 3, 5, 6}, 5) == 2)
	fmt.Println(searchInsert1([]int{1, 3, 5, 6}, 2) == 1)
	fmt.Println(searchInsert1([]int{1, 3, 5, 6}, 7) == 4)
	fmt.Println(searchInsert1([]int{1, 3, 5, 6}, 0) == 0)

	fmt.Println()

	fmt.Println("Solution 2")
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 5) == 2)
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 2) == 1)
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 7) == 4)
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 0) == 0)
}
