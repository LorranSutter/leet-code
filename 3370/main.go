package main

import "fmt"

func smallestNumber(n int) int {
	smallest := 1
	for smallest < n {
		smallest = smallest<<1 | 1
	}

	return smallest
}

func main() {
	fmt.Println(smallestNumber(5) == 7)
	fmt.Println(smallestNumber(10) == 15)
	fmt.Println(smallestNumber(3) == 3)
	fmt.Println(smallestNumber(1000) == 3)
}
