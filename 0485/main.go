package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	maxCount, count := 0, 0

	for i := range nums {
		if nums[i] == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return maxCount
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}) == 3)
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}) == 2)
}
