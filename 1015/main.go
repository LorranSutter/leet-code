package main

import "leetcode/utils"

// The idea here is just to keep track of the remainders until we find one that is 0.
// We can eliminate any multiple of 2 and 5, because these can't divide numbers ending in 1.

// We know that 111 = (11*10 + 1), 1111 = (111*10 + 1), and so on. We use that to iterate
// But we just have to keep the remainder of k because (a % k + b % k) % k == (a + b) % k
// e.g. 111 % 3 == (11*10 + 1) % 3 == ((11 % 3)*(10 % 3) + 1 % 3) % 3

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	remainder := 1
	count := 1
	for remainder%k != 0 {
		count++
		remainder = (remainder*10 + 1) % k
	}

	return count
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: 1, Got: smallestRepunitDivByK(1), Expected: 1},
		{Input: 2, Got: smallestRepunitDivByK(2), Expected: -1},
		{Input: 3, Got: smallestRepunitDivByK(3), Expected: 3},
		{Input: 4, Got: smallestRepunitDivByK(4), Expected: -1},
		{Input: 5, Got: smallestRepunitDivByK(5), Expected: -1},
		{Input: 5, Got: smallestRepunitDivByK(5), Expected: -1},
		{Input: 9, Got: smallestRepunitDivByK(9), Expected: 9},
		{Input: 11, Got: smallestRepunitDivByK(11), Expected: 2},
		{Input: 13, Got: smallestRepunitDivByK(13), Expected: 6},
		{Input: 17, Got: smallestRepunitDivByK(17), Expected: 16},
	})
}
