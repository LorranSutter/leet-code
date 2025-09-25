package main

import "fmt"

func sortColors(nums []int) {
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

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
	sortColors([]int{2, 0, 1})
}
