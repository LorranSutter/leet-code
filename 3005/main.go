package main

import "leetcode/utils"

const MAX_NUM int = 100

func maxFrequencyElements(nums []int) int {
	freqs := make([]int, MAX_NUM+1)
	maxFrequency := 0
	numElemMaxFreq := 0

	for _, num := range nums {
		freqs[num]++
		if freqs[num] > maxFrequency {
			maxFrequency = freqs[num]
			numElemMaxFreq = 1
		} else if freqs[num] == maxFrequency {
			numElemMaxFreq++
		}
	}

	return numElemMaxFreq * maxFrequency
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{1, 2, 2, 3, 1, 4}, Got: maxFrequencyElements([]int{1, 2, 2, 3, 1, 4}), Expected: 4},
		{Input: []int{1, 2, 3, 4, 5}, Got: maxFrequencyElements([]int{1, 2, 3, 4, 5}), Expected: 5},
		{
			Input:    []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 7, 8, 8, 9, 10, 10, 11},
			Got:      maxFrequencyElements([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 7, 8, 8, 9, 10, 10, 11}),
			Expected: 15,
		},
	})
}
