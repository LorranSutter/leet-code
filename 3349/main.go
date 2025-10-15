package main

import "fmt"

func hasIncreasingSubarrays(nums []int, k int) bool {
	if k == 1 {
		return true
	}

	size_subarray := 1
	for i := 1; i < len(nums)-k; i++ {
		if nums[i] > nums[i-1] && nums[i+k] > nums[i+k-1] {
			size_subarray++
			if size_subarray == k {
				return true
			}
		} else if size_subarray != 1 {
			size_subarray = 1
		}
	}

	return false
}

func main() {
	fmt.Println(hasIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3) == true)
	fmt.Println(hasIncreasingSubarrays([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 5) == false)
	fmt.Println(hasIncreasingSubarrays([]int{0, 4, 16, 20, -6}, 2) == true)
	fmt.Println(hasIncreasingSubarrays([]int{-15, 19}, 1) == true)
	fmt.Println(hasIncreasingSubarrays([]int{5, 8, -2, -1}, 2) == true)
}
