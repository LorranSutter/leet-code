package main

import (
	"leetcode/utils"
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
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{"hello world", "ad"}, Got: canBeTypedWords("hello world", "ad"), Expected: 1},
		{Input: []any{"leet code", "lt"}, Got: canBeTypedWords("leet code", "lt"), Expected: 1},
		{Input: []any{"leet code", "e"}, Got: canBeTypedWords("leet code", "e"), Expected: 0},
		{Input: []any{"leet code", "ab"}, Got: canBeTypedWords("leet code", "ab"), Expected: 2},
	})
}
