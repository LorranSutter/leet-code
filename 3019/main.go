package main

import (
	"fmt"
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
	fmt.Println("Solution 1")
	fmt.Println(countKeyChanges1("aAbBcC") == 2)
	fmt.Println(countKeyChanges1("AaAaAaaA") == 0)

	fmt.Println()

	fmt.Println("Solution 2")
	fmt.Println(countKeyChanges2("aAbBcC") == 2)
	fmt.Println(countKeyChanges2("AaAaAaaA") == 0)
}
