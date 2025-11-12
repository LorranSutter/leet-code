package main

import "fmt"

func singleNumber(nums []int) int {
	num := 0
	for i := range nums {
		num ^= nums[i]
	}

	return num
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}) == 1)
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}) == 4)
	fmt.Println(singleNumber([]int{1}) == 1)
}
