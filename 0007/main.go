package main

import (
	"fmt"
	"strconv"
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
	fmt.Println("Solution 1")
	fmt.Println(reverse1(-2147483642) == 0)
	fmt.Println(reverse1(2147483642) == 0)
	fmt.Println(reverse1(123) == 321)
	fmt.Println(reverse1(-123) == -321)
	fmt.Println(reverse1(120) == 21)

	fmt.Println()

	fmt.Println("Solution 2")
	fmt.Println(reverse2(-2147483642) == 0)
	fmt.Println(reverse2(2147483642) == 0)
	fmt.Println(reverse2(123) == 321)
	fmt.Println(reverse2(-123) == -321)
	fmt.Println(reverse2(120) == 21)
}
