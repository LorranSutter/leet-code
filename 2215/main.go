package main

import (
	"fmt"
	"leetcode/utils"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	num1_map := map[int]struct{}{}
	num2_map := map[int]struct{}{}
	diff1 := map[int]struct{}{}
	diff2 := map[int]struct{}{}
	diffs := [][]int{}
	diffs = append(diffs, []int{})
	diffs = append(diffs, []int{})

	for _, num := range nums1 {
		num1_map[num] = struct{}{}
	}
	for _, num := range nums2 {
		num2_map[num] = struct{}{}
	}

	for _, num := range nums1 {
		if _, ok := num2_map[num]; !ok {
			if _, ok := diff1[num]; !ok {
				diff1[num] = struct{}{}
				diffs[0] = append(diffs[0], num)
			}
		}
	}
	for _, num := range nums2 {
		if _, ok := num1_map[num]; !ok {
			if _, ok := diff2[num]; !ok {
				diff2[num] = struct{}{}
				diffs[1] = append(diffs[1], num)
			}
		}
	}

	return diffs
}

func main() {
	diff := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
	fmt.Println(utils.EqualSlices(diff[0], []int{1, 3}))
	fmt.Println(utils.EqualSlices(diff[1], []int{4, 6}))

	diff = findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})
	fmt.Println(utils.EqualSlices(diff[0], []int{3}))
	fmt.Println(utils.EqualSlices(diff[1], []int{}))
}
