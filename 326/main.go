package main

import (
	"fmt"
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

	fmt.Println(n, exp, math.Trunc(roundToFiveDecimals(exp)))

	return roundToFiveDecimals(exp) == math.Trunc(exp)

}

func main() {
	fmt.Println(isPowerOfThree1(1) == true)
	fmt.Println(isPowerOfThree1(3) == true)
	fmt.Println(isPowerOfThree1(9) == true)
	fmt.Println(isPowerOfThree1(27) == true)
	fmt.Println(isPowerOfThree1(81) == true)
	fmt.Println(isPowerOfThree1(0) == false)
	fmt.Println(isPowerOfThree1(-1) == false)

	fmt.Println()

	fmt.Println(isPowerOfThree2(1) == true)
	fmt.Println(isPowerOfThree2(3) == true)
	fmt.Println(isPowerOfThree2(9) == true)
	fmt.Println(isPowerOfThree2(27) == true)
	fmt.Println(isPowerOfThree2(81) == true)
	fmt.Println(isPowerOfThree2(0) == false)
	fmt.Println(isPowerOfThree2(-1) == false)
	fmt.Println(isPowerOfThree2(43046722))
}
