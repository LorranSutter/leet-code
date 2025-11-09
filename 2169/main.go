package main

import "fmt"

func countOperations(num1 int, num2 int) int {
	count := 0
	for num1*num2 != 0 {
		if num1 < num2 {
			num2 = num2 - num1
		} else {
			num1 = num1 - num2
		}
		count++
	}

	return count
}

func main() {
	fmt.Println(countOperations(2, 3) == 3)
	fmt.Println(countOperations(10, 10) == 1)
	fmt.Println(countOperations(5, 2) == 4)
}
