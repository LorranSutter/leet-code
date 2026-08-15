package main

import (
	"leetcode/utils"
	"slices"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func maxFreqSum(s string) int {
	maxVowel := 0
	maxConsonant := 0
	mapVowels := make(map[rune]int)
	mapConsonants := make(map[rune]int)

	for _, letter := range s {
		if slices.Contains(vowels, letter) {
			mapVowels[letter]++
			if mapVowels[letter] > maxVowel {
				maxVowel = mapVowels[letter]
			}
		} else {
			mapConsonants[letter]++
			if mapConsonants[letter] > maxConsonant {
				maxConsonant = mapConsonants[letter]
			}
		}
	}

	return maxVowel + maxConsonant
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: "successes", Got: maxFreqSum("successes"), Expected: 6},
		{Input: "aeiaeia", Got: maxFreqSum("aeiaeia"), Expected: 3},
		{Input: "c", Got: maxFreqSum("c"), Expected: 1},
		{Input: "i", Got: maxFreqSum("i"), Expected: 1},
	})
}
