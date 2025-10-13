package main

import (
	"fmt"
	"leetcode/utils"
	"maps"
)

func removeAnagrams(words []string) []string {
	var result []string

	result = append(result, words[0])
	currentCounts := make(map[rune]int)
	for _, letter := range words[0] {
		currentCounts[letter]++
	}

	for i := 1; i < len(words); i++ {
		newCounts := make(map[rune]int)
		for _, letter := range words[i] {
			newCounts[letter]++
		}

		if !maps.Equal(currentCounts, newCounts) {
			currentCounts = newCounts
			result = append(result, words[i])
		}
	}

	return result
}

func main() {
	fmt.Println(utils.EqualSlices(removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"}), []string{"abba", "cd"}))
	fmt.Println(utils.EqualSlices(removeAnagrams([]string{"a", "b", "c", "d", "e"}), []string{"a", "b", "c", "d", "e"}))
}
