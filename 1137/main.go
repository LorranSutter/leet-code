package main

import "fmt"

func tribonacci1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return tribonacci1(n-3) + tribonacci1(n-2) + tribonacci1(n-1)
}

func tribonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}

func main() {
	fmt.Println("Solution 01")
	fmt.Println(tribonacci1(4) == 4)
	fmt.Println(tribonacci1(7) == 24)
	fmt.Println(tribonacci1(25) == 1389537)

	fmt.Println()

	fmt.Println("Solution 02")
	fmt.Println(tribonacci2(4) == 4)
	fmt.Println(tribonacci2(7) == 24)
	fmt.Println(tribonacci2(25) == 1389537)
	fmt.Println(tribonacci2(37) == 2082876103)
}
