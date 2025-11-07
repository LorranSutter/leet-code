package main

import "fmt"

func maxIncreasingSubarrays(nums []int) int {
	for k := len(nums) / 2; k > 1; k-- {
		size_subarray := 1
		for i := 1; i < len(nums)-k; i++ {
			if nums[i] > nums[i-1] && nums[i+k] > nums[i+k-1] {
				size_subarray++
				if size_subarray == k {
					return k
				}
			} else if size_subarray != 1 {
				size_subarray = 1
			}
		}
	}

	return 1
}

func main() {
	fmt.Println(maxIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}) == 3)
	fmt.Println(maxIncreasingSubarrays([]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}) == 2)
	fmt.Println(maxIncreasingSubarrays([]int{0, 4, 16, 20, -6}) == 2)
	fmt.Println(maxIncreasingSubarrays([]int{-15, 19}) == 1)
	fmt.Println(maxIncreasingSubarrays([]int{5, 8, -2, -1}) == 2)
}
