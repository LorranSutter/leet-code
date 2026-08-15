package main

import "leetcode/utils"

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
	utils.RunTests([]utils.TestCase[int]{
		{Input: 4, Got: tribonacci1(4), Expected: 4},
		{Input: 7, Got: tribonacci1(7), Expected: 24},
		{Input: 25, Got: tribonacci1(25), Expected: 1389537},
	})

	utils.RunTests([]utils.TestCase[int]{
		{Input: 4, Got: tribonacci2(4), Expected: 4},
		{Input: 7, Got: tribonacci2(7), Expected: 24},
		{Input: 25, Got: tribonacci2(25), Expected: 1389537},
		{Input: 37, Got: tribonacci2(37), Expected: 2082876103},
	})
}
