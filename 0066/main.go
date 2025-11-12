package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		newDigit := digits[i] + 1
		if newDigit <= 9 {
			digits[i] = newDigit
			return digits
		}

		digits[i] = 0

		if i == 0 {
			return append([]int{1}, digits...)
		}
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{8, 9, 9, 9}))
}
