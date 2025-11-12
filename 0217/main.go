package main

import "fmt"

func containsDuplicate(nums []int) bool {
	dups := make(map[int]bool)

	for _, n := range nums {
		if dups[n] {
			return true
		}
		dups[n] = true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}) == true)
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}) == false)
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) == true)
}
