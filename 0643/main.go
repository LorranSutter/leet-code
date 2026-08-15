package main

import "leetcode/utils"

func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i := range k {
		max += nums[i]
	}

	current_sum := max
	for i := 0; i < len(nums)-k; i++ {
		current_sum += (-nums[i] + nums[i+k])
		if current_sum > max {
			max = current_sum
		}
	}

	return float64(max) / float64(k)
}

func main() {
	utils.RunTests([]utils.TestCase[float64]{
		{Input: []any{[]int{1, 12, -5, -6, 50, 3}, 4}, Got: findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4), Expected: 12.75},
		{Input: []any{[]int{1, 12, -5, -6, 50, 3}, 2}, Got: findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 2), Expected: 26.5},
		{Input: []any{[]int{1, 12, -5, -6, 50, 3}, 1}, Got: findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 1), Expected: 50},
		{Input: []any{[]int{9, 7, 3, 5, 6, 2, 0, 8, 1, 9}, 6}, Got: findMaxAverage([]int{9, 7, 3, 5, 6, 2, 0, 8, 1, 9}, 6), Expected: 5.333333333333333},
		{Input: []any{[]int{5}, 1}, Got: findMaxAverage([]int{5}, 1), Expected: 5},
	})
}
