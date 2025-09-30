package main

import "fmt"

func twoSum(nums []int, target int) []int {
	p0 := 0
	p1 := len(nums) - 1
	diff := target - nums[p0]
	for {
		if p0 == p1 {
			p0++
			p1 = len(nums) - 1
			diff = target - nums[p0]
			continue
		}

		if nums[p1] == diff {
			break
		}

		p1--
	}

	return []int{p0, p1}
}

func equal(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(equal(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}))
	fmt.Println(equal(twoSum([]int{3, 2, 4}, 6), []int{1, 2}))
	fmt.Println(equal(twoSum([]int{3, 3}, 6), []int{0, 1}))
	fmt.Println(equal(twoSum([]int{15, 11, 7, 2}, 9), []int{2, 3}))
}
