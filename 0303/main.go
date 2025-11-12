package main

import "fmt"

type NumArray struct {
	prefixSum []int
}

// 0 ms / 9.24 MB
func Constructor(nums []int) NumArray {
	na := NumArray{make([]int, len(nums)+1)}

	na.prefixSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		na.prefixSum[i] = na.prefixSum[i-1] + nums[i-1]
	}

	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

// Old Solution 31 ms / 9.4 MB
// func Constructor(nums []int) NumArray {
// 	na := NumArray{nums}

// 	return na
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	sum := 0
// 	for i := left; i <= right; i++ {
// 		sum += this.nums[i]
// 	}

// 	return sum
// }

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.prefixSum)
	fmt.Println(obj.SumRange(0, 2) == 1)
	fmt.Println(obj.SumRange(2, 5) == -1)
	fmt.Println(obj.SumRange(0, 5) == -3)
}
