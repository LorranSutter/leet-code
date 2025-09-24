package main

import (
	"fmt"
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
	fmt.Println(maxFreqSum("successes") == 6)
	fmt.Println(maxFreqSum("aeiaeia") == 3)
	fmt.Println(maxFreqSum("c") == 1)
	fmt.Println(maxFreqSum("i") == 1)
}
