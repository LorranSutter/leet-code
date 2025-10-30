package main

import "fmt"

// 0 ms / 9.43 MB
func moveZeroes1(nums []int) {
	p1, p2 := 0, 1

	for p2 < len(nums) {
		if nums[p1] == 0 {
			if nums[p2] != 0 {
				nums[p1], nums[p2] = nums[p2], nums[p1]
				p1++
				p2++
			} else {
				p2++
			}
		} else {
			p1++
			if p1 >= p2 {
				p2 = p1 + 1
			}
		}
	}
}

// 48 ms / 8.82 MB
func moveZeroes2(nums []int) {
	fmt.Println(nums)

	for i := range nums {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}

	fmt.Println(nums)

}

func main() {
	moveZeroes1([]int{0, 1, 0, 3, 12})
	moveZeroes1([]int{1, 0, 0, 3, 12})
	moveZeroes1([]int{0, 0, 1, 3, 12})
	moveZeroes1([]int{1, 3, 12, 0, 0})
	moveZeroes1([]int{0, 0, 1, 3, 12, 0, 0})
	moveZeroes1([]int{0, 0, 1, 0, 3, 0, 0, 12, 13, 34, 55, 0, 0})
	moveZeroes1([]int{0})
	moveZeroes1([]int{1})

	moveZeroes2([]int{0, 1, 0, 3, 12})
	moveZeroes2([]int{1, 0, 0, 3, 12})
	moveZeroes2([]int{0, 0, 1, 3, 12})
	moveZeroes2([]int{1, 3, 12, 0, 0})
	moveZeroes2([]int{0, 0, 1, 3, 12, 0, 0})
	moveZeroes2([]int{0, 0, 1, 0, 3, 0, 0, 12, 13, 34, 55, 0, 0})
	moveZeroes2([]int{0})
	moveZeroes2([]int{1})
}
