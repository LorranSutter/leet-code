package main

import (
	"leetcode/utils"
	"math"
)

var LOG_3 float64 = math.Log(3)

func roundToFiveDecimals(num float64) float64 {
	return math.Round(num*1000000000) / 1000000000
}

func isPowerOfThree1(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}

func isPowerOfThree2(n int) bool {
	if n <= 0 {
		return false
	}

	exp := math.Log(float64(n)) / LOG_3

	return roundToFiveDecimals(exp) == math.Trunc(exp)

}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: 1, Got: isPowerOfThree1(1), Expected: true},
		{Input: 3, Got: isPowerOfThree1(3), Expected: true},
		{Input: 9, Got: isPowerOfThree1(9), Expected: true},
		{Input: 27, Got: isPowerOfThree1(27), Expected: true},
		{Input: 81, Got: isPowerOfThree1(81), Expected: true},
		{Input: 0, Got: isPowerOfThree1(0), Expected: false},
		{Input: -1, Got: isPowerOfThree1(-1), Expected: false},
	})

	utils.RunTests([]utils.TestCase[bool]{
		{Input: 1, Got: isPowerOfThree2(1), Expected: true},
		{Input: 3, Got: isPowerOfThree2(3), Expected: true},
		{Input: 9, Got: isPowerOfThree2(9), Expected: true},
		{Input: 27, Got: isPowerOfThree2(27), Expected: true},
		{Input: 81, Got: isPowerOfThree2(81), Expected: true},
		{Input: 0, Got: isPowerOfThree2(0), Expected: false},
		{Input: -1, Got: isPowerOfThree2(-1), Expected: false},
		{Input: 43046722, Got: isPowerOfThree2(43046722), Expected: false},
	})
}
