package main

import "leetcode/utils"

// 25 ms / 6.7 MB
func maxVowels1(s string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	max := 0

	for i := range k {
		if vowels[s[i]] {
			max++
		}
	}

	if max == k {
		return max
	}

	count_vowels := max
	for i := 0; i < len(s)-k; i++ {
		if vowels[s[i]] {
			count_vowels--
		}
		if vowels[s[i+k]] {
			count_vowels++
		}

		if count_vowels > max {
			if count_vowels == k {
				return count_vowels
			}
			max = count_vowels
		}
	}

	return max
}

// 1 ms / 6.67 MB
func maxVowels2(s string, k int) int {
	max := 0

	for i := range k {
		if isVowel(s[i]) {
			max++
		}
	}

	if max == k {
		return max
	}

	count_vowels := max
	for i := 0; i < len(s)-k; i++ {
		if isVowel(s[i]) {
			count_vowels--
		}
		if isVowel(s[i+k]) {
			count_vowels++
		}

		if count_vowels > max {
			if count_vowels == k {
				return count_vowels
			}
			max = count_vowels
		}
	}

	return max
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{"abciiidef", 3}, Got: maxVowels1("abciiidef", 3), Expected: 3},
		{Input: []any{"aeiou", 2}, Got: maxVowels1("aeiou", 2), Expected: 2},
		{Input: []any{"leetcode", 3}, Got: maxVowels1("leetcode", 3), Expected: 2},
		{Input: []any{"weallloveyou", 7}, Got: maxVowels1("weallloveyou", 7), Expected: 4},
	})

	utils.RunTests([]utils.TestCase[int]{
		{Input: []any{"abciiidef", 3}, Got: maxVowels2("abciiidef", 3), Expected: 3},
		{Input: []any{"aeiou", 2}, Got: maxVowels2("aeiou", 2), Expected: 2},
		{Input: []any{"leetcode", 3}, Got: maxVowels2("leetcode", 3), Expected: 2},
		{Input: []any{"weallloveyou", 7}, Got: maxVowels2("weallloveyou", 7), Expected: 4},
	})
}
