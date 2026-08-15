package main

import (
	"leetcode/utils"
	"math"
)

// Area = (1/2) * |x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2)|
func area(p [3][2]int) float64 {
	return 0.5 * math.Abs(float64(p[0][0]*(p[1][1]-p[2][1])+p[1][0]*(p[2][1]-p[0][1])+p[2][0]*(p[0][1]-p[1][1])))
}

func combinations(p [][]int, size int, accumulated []int) {

}

func largestTriangleArea(points [][]int) float64 {
	// TODO Implement solution
	return 0
}

func main() {
	utils.RunTests([]utils.TestCase[float64]{
		{Input: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}, Got: largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}), Expected: 2.0},
		{Input: [][]int{{1, 0}, {0, 0}, {0, 1}}, Got: largestTriangleArea([][]int{{1, 0}, {0, 0}, {0, 1}}), Expected: 0.5},
	})
}
