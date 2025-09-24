package main

import (
	"fmt"
	"slices"
)

var vowels = []rune{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}

// 719 ms / 10.31 MB
func sortVowels1(s string) string {
	runes := []rune(s)
	var vowelPositions []int
	var vowelSorted []rune

	for i, letter := range runes {
		if slices.Contains(vowels, letter) {
			vowelPositions = append(vowelPositions, i)
			if len(vowelSorted) == 0 {
				vowelSorted = append(vowelSorted, letter)
			} else {
				pos, _ := slices.BinarySearch(vowelSorted, letter)
				vowelSorted = slices.Insert(vowelSorted, pos, letter)
			}
		}
	}

	for i, pos := range vowelPositions {
		runes[pos] = rune(vowelSorted[i])
	}

	return string(runes)
}

// Time limit exceeded, 2212 / 2216
func sortVowels2(s string) string {
	if len(s) == 1 {
		return s
	}

	runes := []rune(s)

	for i := range runes {
		if slices.Contains(vowels, runes[i]) {
			for j := i + 1; j < len(runes); j++ {
				if slices.Contains(vowels, runes[j]) {
					if runes[i] > runes[j] {
						runes[i], runes[j] = runes[j], runes[i]
					}
				}
			}
		}
	}

	return string(runes)
}

func main() {
	fmt.Println(sortVowels1("lEetcOde") == "lEOtcede")
	fmt.Println(sortVowels1("lYmpH") == "lYmpH")
	fmt.Println(sortVowels1("ofaApIUmpshjOoIIAhn") == "AfAIpIImpshjOUaoohn")

	fmt.Println()

	fmt.Println(sortVowels2("lEetcOde") == "lEOtcede")
	fmt.Println(sortVowels2("lYmpH") == "lYmpH")
	fmt.Println(sortVowels2("ofaApIUmpshjOoIIAhn") == "AfAIpIImpshjOUaoohn")
}
