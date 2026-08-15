package main

import (
	"leetcode/utils"
)

// The idea here is to iterate through all possible values of a, b and c
// We start storing all perfect squares up to n*n in a map for O(1) lookup
// Then for each pair (a, b) we check if a^2 + b^2 is a perfect square
// If it is, we found a valid triplet (a, b, c) and (b, a, c), so we increment the count by 2

func countTriples(n int) int {
	squares := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		squares[i*i] = i
	}

	triplets := 0
	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			c := a*a + b*b
			if _, ok := squares[c]; ok {
				triplets += 2
			}
		}
	}

	return triplets
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: 5, Got: countTriples(5), Expected: 2},
		{Input: 10, Got: countTriples(10), Expected: 4},
		{Input: 11, Got: countTriples(11), Expected: 4},
	})
}
