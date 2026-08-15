package main

import (
	"strconv"

	"leetcode/utils"
)

const MAX_INT_32 int = 2147483647

// 2 ms / 4 MB
func reverse1(x int) int {
	var result int

	var negative bool
	if x < 0 {
		negative = true
		x = -x
	}

	for x > 0 {
		digit := x % 10
		if MAX_INT_32-digit < result*10 {
			return 0
		}
		result = result*10 + digit
		x /= 10
	}

	if negative {
		return -result
	}
	return result
}

// 0 ms / 4.01 MB
func reverse2(x int) int {
	var negative bool
	if x < 0 {
		negative = true
		x = -x
	}

	xStr := []rune(strconv.Itoa(x))

	for i, j := 0, len(xStr)-1; i < j; i, j = i+1, j-1 {
		xStr[i], xStr[j] = xStr[j], xStr[i]
	}

	x, _ = strconv.Atoi(string(xStr))

	if x >= MAX_INT_32 {
		return 0
	}
	if negative {
		return -x
	}
	return x
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: -2147483642, Got: reverse1(-2147483642), Expected: 0},
		{Input: 2147483642, Got: reverse1(2147483642), Expected: 0},
		{Input: 123, Got: reverse1(123), Expected: 321},
		{Input: -123, Got: reverse1(-123), Expected: -321},
		{Input: 120, Got: reverse1(120), Expected: 21},
	})

	utils.RunTests([]utils.TestCase[int]{
		{Input: -2147483642, Got: reverse2(-2147483642), Expected: 0},
		{Input: 2147483642, Got: reverse2(2147483642), Expected: 0},
		{Input: 123, Got: reverse2(123), Expected: 321},
		{Input: -123, Got: reverse2(-123), Expected: -321},
		{Input: 120, Got: reverse2(120), Expected: 21},
	})
}
