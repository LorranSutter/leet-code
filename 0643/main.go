package main

import "fmt"

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
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4) == 12.75)
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 2) == 26.5)
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 1) == 50)
	fmt.Println(findMaxAverage([]int{9, 7, 3, 5, 6, 2, 0, 8, 1, 9}, 6) == 5.333333333333333)
	fmt.Println(findMaxAverage([]int{5}, 1) == 5)
}
