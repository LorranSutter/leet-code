package main

import "leetcode/utils"

// If the last bit is 1, then it's not a one-bit character
// Otherwise, we just have to check if the next bit is 1 and skip the next bit
// If the last iteration it is the last bit, then it's a one-bit character

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	if bits[n-1] == 1 {
		return false
	}

	i := 0
	for ; i < n-1; i++ {
		if bits[i] == 1 {
			i++
		}
	}

	return i == n-1
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []int{1, 0, 0}, Got: isOneBitCharacter([]int{1, 0, 0}), Expected: true},
		{Input: []int{1, 1, 1, 0}, Got: isOneBitCharacter([]int{1, 1, 1, 0}), Expected: false},
		{Input: []int{1, 1, 1, 1}, Got: isOneBitCharacter([]int{1, 1, 1, 1}), Expected: false},
		{Input: []int{1, 0, 1, 0}, Got: isOneBitCharacter([]int{1, 0, 1, 0}), Expected: false},
	})
}
