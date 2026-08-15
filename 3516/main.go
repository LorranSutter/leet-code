package main

import "leetcode/utils"

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
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{2, 7, 4}, Got: findClosest(2, 7, 4), Expected: 1},
		{Input: []int{2, 5, 6}, Got: findClosest(2, 5, 6), Expected: 2},
		{Input: []int{1, 5, 3}, Got: findClosest(1, 5, 3), Expected: 0},
	})
}
