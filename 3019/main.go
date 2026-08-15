package main

import (
	"leetcode/utils"
	"unicode"
)

// 0 ms / 4.05 MB
func countKeyChanges1(s string) int {
	count := 0
	previous := unicode.ToUpper(rune(s[0]))

	for i := 1; i < len(s); i++ {
		if unicode.ToUpper(rune(s[i])) != previous {
			count++
		}
		previous = unicode.ToUpper(rune(s[i]))
	}

	return count
}

// 0 ms / 4.11 MB
func countKeyChanges2(s string) int {
	count := 0

	for i := 1; i < len(s); i++ {
		if unicode.ToUpper(rune(s[i])) != unicode.ToUpper(rune(s[i-1])) {
			count++
		}
	}

	return count
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: "aAbBcC", Got: countKeyChanges1("aAbBcC"), Expected: 2},
		{Input: "AaAaAaaA", Got: countKeyChanges1("AaAaAaaA"), Expected: 0},
	})

	utils.RunTests([]utils.TestCase[int]{
		{Input: "aAbBcC", Got: countKeyChanges2("aAbBcC"), Expected: 2},
		{Input: "AaAaAaaA", Got: countKeyChanges2("AaAaAaaA"), Expected: 0},
	})
}
