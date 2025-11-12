package main

import "fmt"

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	newMin := min(height[start], height[end])
	newArea := newMin * (end - start)
	maxArea := newArea

	for start < end {
		if height[start] == newMin {
			newArea = height[start] * (end - start)
			start++
		} else {
			newArea = height[end] * (end - start)
			end--
		}
		newMin = min(height[start], height[end])

		if newArea > maxArea {
			maxArea = newArea
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	fmt.Println(maxArea([]int{1, 1}) == 1)
	fmt.Println(maxArea([]int{1, 2, 1}) == 2)
	fmt.Println(maxArea([]int{1, 2, 4, 3}) == 4)
	fmt.Println(maxArea([]int{7, 10, 6, 2, 5, 4, 8, 3, 7}) == 56)
	fmt.Println(maxArea([]int{5, 4, 3, 2, 1}) == 6)
}
