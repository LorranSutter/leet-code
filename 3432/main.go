package main

import "fmt"

// Just make the prefix sum of the array first
// Iterate checking the difference between the sums of the left and right subarrays
// If the difference is even, we increment the count of partitions

func countPartitions(nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	for i := range nums {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	partitions, left, right := 0, 0, 0
	for i := 1; i < len(nums); i++ {
		left = prefixSum[i] - prefixSum[0]
		right = prefixSum[len(nums)] - prefixSum[i]

		if (right-left)%2 == 0 {
			partitions += 1
		}
	}

	return partitions
}

func main() {
	fmt.Println(countPartitions([]int{10, 10, 3, 7, 6}) == 4)
	fmt.Println(countPartitions([]int{1, 2, 2}) == 0)
	fmt.Println(countPartitions([]int{2, 4, 6, 8}) == 3)
	fmt.Println(countPartitions([]int{2, 4}) == 1)
}
