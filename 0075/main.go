package main

import (
	"leetcode/utils"
)

// 0 ms / 4.26 MB
func sortColors1(nums []int) []int {
	quicksort(&nums, 0, len(nums)-1)

	return nums
}

func quicksort(nums *[]int, lo int, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	p := partition(nums, lo, hi)

	quicksort(nums, lo, p-1)
	quicksort(nums, p+1, hi)
}

func partition(nums *[]int, lo int, hi int) int {
	pivot := (*nums)[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if (*nums)[j] <= pivot {
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			i++
		}
	}
	(*nums)[i], (*nums)[hi] = (*nums)[hi], (*nums)[i]
	return i

}

// 0 ms / 4.08 MB
func sortColors2(nums []int) []int {
	freqs := make([]int, 3)

	// Set frequencies for each color
	for _, num := range nums {
		freqs[num]++
	}

	currentSetColor := 0
	for i := range nums {
	StartLoop:
		if freqs[currentSetColor] == 0 {
			// If no color left, go to the next one
			currentSetColor++
			goto StartLoop
		}
		freqs[currentSetColor]--
		nums[i] = currentSetColor
	}

	return nums
}

func main() {
	utils.RunTests([]utils.TestCase[[]int]{
		{Input: []int{2, 0, 2, 1, 1, 0}, Got: sortColors1([]int{2, 0, 2, 1, 1, 0}), Expected: []int{0, 0, 1, 1, 2, 2}},
		{Input: []int{2, 0, 1}, Got: sortColors1([]int{2, 0, 1}), Expected: []int{0, 1, 2}},
		{Input: []int{2}, Got: sortColors1([]int{2}), Expected: []int{2}},
	})

	utils.RunTests([]utils.TestCase[[]int]{
		{Input: []int{2, 0, 2, 1, 1, 0}, Got: sortColors2([]int{2, 0, 2, 1, 1, 0}), Expected: []int{0, 0, 1, 1, 2, 2}},
		{Input: []int{2, 0, 1}, Got: sortColors2([]int{2, 0, 1}), Expected: []int{0, 1, 2}},
		{Input: []int{2}, Got: sortColors2([]int{2}), Expected: []int{2}},
	})
}
