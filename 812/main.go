package main

import (
	"fmt"
	"math"
)

// Area = (1/2) * |x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2)|
func area(p [3][2]int) float64 {
	return 0.5 * math.Abs(float64(p[0][0]*(p[1][1]-p[2][1])+p[1][0]*(p[2][1]-p[0][1])+p[2][0]*(p[0][1]-p[1][1])))
}

func combinations(p [][]int, size int, accumulated []int) {

}

func largestTriangleArea(points [][]int) float64 {

	return 0
}

func main() {
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}) == 2.0)
	fmt.Println(largestTriangleArea([][]int{{1, 0}, {0, 0}, {0, 1}}) == 0.5)
}
