package main

import "fmt"

func largestAltitude(gain []int) int {
	highest, current := 0, 0

	for _, g := range gain {
		current += g
		if current > highest {
			highest = current
		}
	}

	return highest
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}) == 1)
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}) == 0)
}
