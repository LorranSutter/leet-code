package main

import "fmt"

// The idea is to keep track of the number of ones
// Every time we encounter a zero after a one, we add the count of ones to the maximum number of operations

func maxOperations(s string) int {
	maxOp, ones := 0, 0
	newZero := false

	for _, c := range s {
		if c == '1' {
			ones++
			newZero = false
		} else if !newZero {
			maxOp += ones
			newZero = true
		}
	}

	return maxOp
}

func main() {
	fmt.Println(maxOperations("1001101") == 4)
	fmt.Println(maxOperations("00111") == 0)
	fmt.Println(maxOperations("10011010") == 8)
}
