package main

import (
	"leetcode/utils"
)

func areaOfMaxDiagonal(dimensions [][]int) int {
	var maxDiag int
	var maxArea int

	for _, dims := range dimensions {
		diag := dims[0]*dims[0] + dims[1]*dims[1]

		if diag == maxDiag {
			newArea := dims[0] * dims[1]

			if newArea > maxArea {
				maxArea = newArea
			}
		} else if diag > maxDiag {
			maxDiag = diag
			maxArea = dims[0] * dims[1]
		}

	}

	return maxArea
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: [][]int{{4, 3}, {3, 4}}, Got: areaOfMaxDiagonal([][]int{{4, 3}, {3, 4}}), Expected: 12},
	})
}
