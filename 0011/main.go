package main

import (
	"leetcode/utils"
)

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
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, Got: maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), Expected: 49},
		{Input: []int{1, 1}, Got: maxArea([]int{1, 1}), Expected: 1},
		{Input: []int{1, 2, 1}, Got: maxArea([]int{1, 2, 1}), Expected: 2},
		{Input: []int{1, 2, 4, 3}, Got: maxArea([]int{1, 2, 4, 3}), Expected: 4},
		{Input: []int{7, 10, 6, 2, 5, 4, 8, 3, 7}, Got: maxArea([]int{7, 10, 6, 2, 5, 4, 8, 3, 7}), Expected: 56},
		{Input: []int{5, 4, 3, 2, 1}, Got: maxArea([]int{5, 4, 3, 2, 1}), Expected: 6},
	})
}
