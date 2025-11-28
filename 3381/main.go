package main

import (
	"fmt"
	"math"
)

// TLE
// The idea is to use prefix sums to calculate the sum of all subarrays
// whose lengths are multiples of k, and keep track of the maximum sum found.
// However, this approach has a time complexity of O(n^2) which is not efficient for large inputs.
func maxSubarraySum1(nums []int, k int) int64 {
	maxSum := int64(-1 << 63)
	size := (len(nums) / k) * k
	prefixSum := make([]int64, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + int64(nums[i-1])
	}

	for size >= k {
		for i := 0; i <= len(nums)-size; i++ {
			subarraySum := prefixSum[i+size] - prefixSum[i]
			if subarraySum > maxSum {
				maxSum = subarraySum
			}
		}
		size -= k
	}
	return maxSum
}

// 25 ms / 17.57 MB
// Optimized solution using Kadane's algorithm variant
// The idea is to iterate through the array in k different starting points
// and apply a modified Kadane's algorithm to find the maximum subarray sum
// for subarrays whose lengths are multiples of k.
func maxSubarraySum2(nums []int, k int) int64 {
	maxSum := int64(-1 << 63)
	prefixSum := make([]int64, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + int64(nums[i-1])
		if i%k == 0 {
		}
	}

	for start := range k {
		current := int64(0)
		best := int64(-1 << 63)

		for i := start; i+k-1 < len(nums); i += k {
			subarraySum := prefixSum[i+k] - prefixSum[i]

			current = int64(math.Max(float64(subarraySum), float64(current+subarraySum)))
			best = int64(math.Max(float64(best), float64(current)))
		}
		maxSum = int64(math.Max(float64(maxSum), float64(best)))
	}
	return maxSum
}

func main() {
	fmt.Println("Solution 01")
	fmt.Println(maxSubarraySum1([]int{1, 2}, 1) == 3)
	fmt.Println(maxSubarraySum1([]int{-1, -2, -3, -4, -5}, 4) == -10)
	fmt.Println(maxSubarraySum1([]int{-5, 1, 2, -3, 4}, 2) == 4)
	fmt.Println(maxSubarraySum1([]int{-10, -1}, 1) == -1)

	fmt.Println()

	fmt.Println("Solution 02")
	fmt.Println(maxSubarraySum2([]int{1, 2}, 1) == 3)
	fmt.Println(maxSubarraySum2([]int{-1, -2, -3, -4, -5}, 4) == -10)
	fmt.Println(maxSubarraySum2([]int{-5, 1, 2, -3, 4}, 2) == 4)
	fmt.Println(maxSubarraySum2([]int{-10, -1}, 1) == -1)
}
