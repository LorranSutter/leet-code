package main

import "fmt"

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosest(x int, y int, z int) int {
	x_to_z := AbsInt(z - x)
	y_to_z := AbsInt(z - y)

	if x_to_z == y_to_z {
		return 0
	}
	if x_to_z < y_to_z {
		return 1
	}

	return 2
}

func main() {
	fmt.Println(findClosest(2, 7, 4))
	fmt.Println(findClosest(2, 5, 6))
	fmt.Println(findClosest(1, 5, 3))
}
