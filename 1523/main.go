package main

import "fmt"

func countOdds(low int, high int) int {
	if high%2 == 0 && low%2 == 0 {
		return (high - low) / 2
	}

	return (high-low)/2 + 1
}

func main() {
	fmt.Println(countOdds(3, 7) == 3)
	fmt.Println(countOdds(8, 10) == 1)
	fmt.Println(countOdds(7, 10) == 2)
	fmt.Println(countOdds(2, 11) == 5)
	fmt.Println(countOdds(2, 2) == 0)
	fmt.Println(countOdds(7, 7) == 1)
}
