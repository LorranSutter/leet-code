package main

import "fmt"

func minimumOperations(nums []int) int {
	count := 0
	for i := range nums {
		if nums[i]%3 != 0 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4}) == 3)
	fmt.Println(minimumOperations([]int{3, 6, 9}) == 0)
}
