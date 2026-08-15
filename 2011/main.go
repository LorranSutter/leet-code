package main

import (
	"leetcode/utils"
)

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if op[1] == '+' {
			x++
		} else {
			x--
		}
	}

	return x
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []string{"--X", "X++", "X++"}, Got: finalValueAfterOperations([]string{"--X", "X++", "X++"}), Expected: 1},
		{Input: []string{"++X", "++X", "X++"}, Got: finalValueAfterOperations([]string{"++X", "++X", "X++"}), Expected: 3},
		{Input: []string{"X++", "++X", "--X", "X--"}, Got: finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"}), Expected: 0},
	})
}
