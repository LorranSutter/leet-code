package main

import (
	"fmt"
	"leetcode/utils"
	"math"
)

// TLE
func countBits1(n int) []int {
	bits := make([]int, n+1)

	for i := 0; i <= n; i++ {
		bits[i] = count(i)
	}

	return bits
}

func count(n int) int {
	bits := 0
	for i := range n {
		if n&(1<<i) != 0 {
			bits++
		}
	}

	return bits
}

// 24 ms / 11.15 MB
func countBits2(n int) []int {
	// 2^0
	// 0 -> 0
	// 1 -> 1

	// 2^1
	// 2 -> 10 -> 2 + 0
	// 3 -> 11 -> 2 + 1

	// 2^2
	// 4 -> 100 -> 4 + 0
	// 5 -> 101 -> 4 + 1
	// 6 -> 110 -> 4 + 2
	// 7 -> 111 -> 4 + 3

	// 2^3
	// 8  -> 1000 -> 8 + 0
	// 9  -> 1001 -> 8 + 1
	// 10 -> 1010 -> 8 + 2
	// 11 -> 1011 -> 8 + 3
	// 12 -> 1100 -> 8 + 4
	// 13 -> 1101 -> 8 + 5
	// 14 -> 1110 -> 8 + 6
	// 15 -> 1111 -> 8 + 7

	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	bits := make([]int, n+1)
	bits_dp := map[int]int{}
	bits_dp[0] = 0

	for i := 1; i <= n; i++ {
		base := int(math.Pow(2, float64(i-1)))
		for j := 0; j < 2<<i && base+j <= n; j++ {
			bits[base+j] = 1 + bits_dp[j]
			bits_dp[base+j] = bits[base+j]
		}
	}

	return bits
}

func countBits3(n int) []int {
	bits := make([]int, n+1)
	for i := range n + 1 {
		bits[i] = bits[i>>1] + (i & 1)
	}
	return bits
}

func main() {
	fmt.Println("Solution 01")
	fmt.Println(utils.DeepEqualSlices(countBits1(2), []int{0, 1, 1}))
	fmt.Println(utils.DeepEqualSlices(countBits1(5), []int{0, 1, 1, 2, 1, 2}))
	fmt.Println(utils.DeepEqualSlices(countBits1(10), []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}))
	fmt.Println(utils.DeepEqualSlices(countBits1(17), []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2}))

	fmt.Println()

	fmt.Println("Solution 02")
	fmt.Println(utils.DeepEqualSlices(countBits2(2), []int{0, 1, 1}))
	fmt.Println(utils.DeepEqualSlices(countBits2(5), []int{0, 1, 1, 2, 1, 2}))
	fmt.Println(utils.DeepEqualSlices(countBits2(10), []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}))
	fmt.Println(utils.DeepEqualSlices(countBits2(17), []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2}))

	fmt.Println()

	fmt.Println("Solution 03")
	fmt.Println(utils.DeepEqualSlices(countBits3(2), []int{0, 1, 1}))
	fmt.Println(utils.DeepEqualSlices(countBits3(5), []int{0, 1, 1, 2, 1, 2}))
	fmt.Println(utils.DeepEqualSlices(countBits3(10), []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}))
	fmt.Println(utils.DeepEqualSlices(countBits3(17), []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2}))
}
