package main

import "fmt"

func countValidSelections(nums []int) int {
	lenPrefix := len(nums) + 1
	sum_left := make([]int, lenPrefix)
	sum_right := make([]int, lenPrefix)

	for i, j := 0, len(nums); i < len(nums); i, j = i+1, j-1 {
		sum_left[i+1] = sum_left[i] + nums[i]
		sum_right[j-1] = sum_right[j] + nums[j-1]
	}

	total := 0
	for i := 0; i < lenPrefix-1; i++ {
		if nums[i] != 0 {
			continue
		}

		if sum_left[i] == sum_right[i] {
			total += 2
		} else if sum_left[i]-1 == sum_right[i] {
			total += 1
		} else if sum_left[i] == sum_right[i]-1 {
			total += 1
		}
	}

	return total
}

func main() {
	fmt.Println(countValidSelections([]int{1, 0, 2, 0, 3}) == 2)
	fmt.Println(countValidSelections([]int{1, 0, 2, 0, 4}) == 1)
	fmt.Println(countValidSelections([]int{1, 0, 2, 0, 2}) == 1)
	fmt.Println(countValidSelections([]int{2, 3, 4, 0, 4, 1, 0}) == 0)
	fmt.Println(countValidSelections([]int{0}) == 2)
	fmt.Println(countValidSelections([]int{16, 13, 10, 0, 0, 0, 10, 6, 7, 8, 7}) == 3)
}
