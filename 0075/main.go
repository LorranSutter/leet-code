package main

import "fmt"

// 0 ms / 4.26 MB
func sortColors1(nums []int) {
	quicksort(&nums, 0, len(nums)-1)

	fmt.Println(nums)
}

func quicksort(nums *[]int, lo int, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	p := partition(nums, lo, hi)

	quicksort(nums, lo, p-1)
	quicksort(nums, p+1, hi)
}

func partition(nums *[]int, lo int, hi int) int {
	pivot := (*nums)[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if (*nums)[j] <= pivot {
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			i++
		}
	}
	(*nums)[i], (*nums)[hi] = (*nums)[hi], (*nums)[i]
	return i

}

// 0 ms / 4.08 MB
func sortColors2(nums []int) {
	freqs := make([]int, 3)

	// Set frequencies for each color
	for _, num := range nums {
		freqs[num]++
	}

	currentSetColor := 0
	for i := range nums {
	StartLoop:
		if freqs[currentSetColor] == 0 {
			// If no color left, go to the next one
			currentSetColor++
			goto StartLoop
		}
		freqs[currentSetColor]--
		nums[i] = currentSetColor
	}

	fmt.Println(nums)
}

func main() {
	fmt.Println("Solution 1")
	sortColors1([]int{2, 0, 2, 1, 1, 0})
	sortColors1([]int{2, 0, 1})
	sortColors1([]int{2})

	fmt.Println()

	fmt.Println("Solution 2")
	sortColors2([]int{2, 0, 2, 1, 1, 0})
	sortColors2([]int{2, 0, 1})
	sortColors2([]int{2})
}
