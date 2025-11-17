package main

import "fmt"

// Keep track of the last 1's position and check the diff between current and last 1's position

func kLengthApart(nums []int, k int) bool {
	lastOne := -k - 1
	for i, num := range nums {
		if num == 1 {
			if i-lastOne-1 < k {
				return false
			}
			lastOne = i
		}
	}
	return true
}

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2) == true)
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2) == false)
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 0) == true)
	fmt.Println(kLengthApart([]int{0, 0, 0, 0, 0, 0}, 5) == true)
	fmt.Println(kLengthApart([]int{1, 1, 0, 1, 0, 0}, 1) == false)
	fmt.Println(kLengthApart([]int{0, 1, 0, 1, 0, 1}, 1) == true)
}
