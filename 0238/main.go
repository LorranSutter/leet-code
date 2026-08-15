package main

import (
	"leetcode/utils"
)

// 4 ms / 10.29 MB
func productExceptSelf1(nums []int) []int {
	product, zeroes := 1, 0
	result := make([]int, len(nums))
	for i := range nums {
		if nums[i] != 0 {
			product *= nums[i]
		} else {
			zeroes++
			if zeroes > 1 {
				return result
			}
		}
	}

	if zeroes > 0 {
		for i := range nums {
			if nums[i] == 0 {
				result[i] = product
				return result
			}
		}
	}

	for i := range nums {
		result[i] = product / nums[i]
	}

	return result
}

// 0 ms / 8.71 MB
func productExceptSelf2(nums []int) []int {
	product, zeroes := 1, 0
	for i := range nums {
		if nums[i] != 0 {
			product *= nums[i]
		} else {
			zeroes++
		}
	}

	if zeroes > 1 {
		for i := range nums {
			nums[i] = 0
		}
	} else if zeroes > 0 {
		for i := range nums {
			if nums[i] == 0 {
				nums[i] = product
			} else {
				nums[i] = 0
			}
		}
	} else {
		for i := range nums {
			nums[i] = product / nums[i]
		}
	}

	return nums
}

// 0 ms / 10.06 MB
func productExceptSelf3(nums []int) []int {
	prefix := make([]int, len(nums))
	left := 1
	for i := range nums {
		prefix[i] = left
		left *= nums[i]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		val := prefix[i] * right
		right *= nums[i]
		nums[i] = val
	}

	return nums
}

func main() {
	utils.RunTests([]utils.TestCase[[]int]{
		{Input: []int{1, 2, 3, 4}, Got: productExceptSelf1([]int{1, 2, 3, 4}), Expected: []int{24, 12, 8, 6}},
		{Input: []int{-1, 1, 0, -3, 3}, Got: productExceptSelf1([]int{-1, 1, 0, -3, 3}), Expected: []int{0, 0, 9, 0, 0}},
		{Input: []int{0, 1, 0, -3, 3}, Got: productExceptSelf1([]int{0, 1, 0, -3, 3}), Expected: []int{0, 0, 0, 0, 0}},
	})

	utils.RunTests([]utils.TestCase[[]int]{
		{Input: []int{1, 2, 3, 4}, Got: productExceptSelf2([]int{1, 2, 3, 4}), Expected: []int{24, 12, 8, 6}},
		{Input: []int{-1, 1, 0, -3, 3}, Got: productExceptSelf2([]int{-1, 1, 0, -3, 3}), Expected: []int{0, 0, 9, 0, 0}},
		{Input: []int{0, 1, 0, -3, 3}, Got: productExceptSelf2([]int{0, 1, 0, -3, 3}), Expected: []int{0, 0, 0, 0, 0}},
	})

	utils.RunTests([]utils.TestCase[[]int]{
		{Input: []int{1, 2, 3, 4}, Got: productExceptSelf3([]int{1, 2, 3, 4}), Expected: []int{24, 12, 8, 6}},
		{Input: []int{-1, 1, 0, -3, 3}, Got: productExceptSelf3([]int{-1, 1, 0, -3, 3}), Expected: []int{0, 0, 9, 0, 0}},
		{Input: []int{0, 1, 0, -3, 3}, Got: productExceptSelf3([]int{0, 1, 0, -3, 3}), Expected: []int{0, 0, 0, 0, 0}},
	})
}
