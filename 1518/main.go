package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	total := numBottles
	numNewBottles := numBottles

	for numNewBottles > numExchange {
		fmt.Println("numBottles before",numBottles)
		numNewBottles = numBottles / numExchange
		numBottles = numNewBottles + numBottles % numExchange
		fmt.Println(" numBottles after",numBottles)
		total += numNewBottles
		fmt.Println("      total after",total)
	}

	fmt.Println(total)
	return total
}

func main() {
	fmt.Println(numWaterBottles(9, 3) == 13)
	fmt.Println(numWaterBottles(15, 4) == 19)
	// fmt.Println(numWaterBottles(15, 2) == 23)
}
