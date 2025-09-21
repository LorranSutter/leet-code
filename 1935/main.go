package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	if brokenLetters == "" {
		return len(strings.Split(text, " "))
	}

	brokenMap := make(map[rune]bool)
	for _, letter := range brokenLetters {
		brokenMap[letter] = true
	}

	totalWords := 0
	invalidWord := 0
	for _, word := range strings.Split(text, " ") {
		totalWords++
		for _, letter := range word {
			if brokenMap[letter] {
				invalidWord++
				break
			}
		}
	}

	return totalWords - invalidWord
}

func main() {
	fmt.Println(canBeTypedWords("hello world", "ad"))
	fmt.Println(canBeTypedWords("leet code", "lt"))
	fmt.Println(canBeTypedWords("leet code", "e"))
	fmt.Println(canBeTypedWords("leet code", "ab"))
}
