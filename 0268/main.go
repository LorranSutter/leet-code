package main

import "fmt"

// 0 ms / 8.8 MB
func missingNumber(nums []int) int {
	hasZero := false
	sum, max := 0, 0
	for _, num := range nums {
		if num == 0 {
			hasZero = true
		} else if num > max {
			max = num
		}
		sum += num
	}

	lenNums := len(nums) + 1
	calcSum := (lenNums*lenNums - lenNums) / 2

	if !hasZero {
		return 0
	}
	if sum == calcSum {
		return max + 1
	}

	return calcSum - sum
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}) == 2)
	fmt.Println(missingNumber([]int{0, 1}) == 2)
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) == 8)
	fmt.Println(missingNumber([]int{2}) == 0)
}
