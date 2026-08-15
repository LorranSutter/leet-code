package main

import (
	"leetcode/utils"
)

func removeStars(s string) string {
	result := ""
	to_remove := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			to_remove++
		} else {
			if to_remove > 0 {
				to_remove--
			} else {
				result = string(s[i]) + result
			}
		}
	}

	return result
}

func main() {
	utils.RunTests([]utils.TestCase[string]{
		{Input: "leet**cod*e", Got: removeStars("leet**cod*e"), Expected: "lecoe"},
		{Input: "erase*****", Got: removeStars("erase*****"), Expected: ""},
	})
}
