package main

import (
	"strings"

	"leetcode/utils"
)

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}

	maxLen := 0
	currentLen := 1
	for i, j := 0, 1; j < sLen; j++ {
		indexRune := strings.IndexRune(s[i:j], rune(s[j]))
		if indexRune != -1 {
			if currentLen > maxLen {
				maxLen = currentLen
			}
			i += indexRune + 1
			currentLen = len(s[i:j])
		}
		currentLen++
	}

	if currentLen > maxLen {
		return currentLen
	}
	return maxLen
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: "abcabcbb", Got: lengthOfLongestSubstring("abcabcbb"), Expected: 3},
		{Input: "abcdabcbb", Got: lengthOfLongestSubstring("abcdabcbb"), Expected: 4},
		{Input: "abcabcdbabc", Got: lengthOfLongestSubstring("abcabcdbabc"), Expected: 4},
		{Input: "bbbbb", Got: lengthOfLongestSubstring("bbbbb"), Expected: 1},
		{Input: "pwwkew", Got: lengthOfLongestSubstring("pwwkew"), Expected: 3},
		{Input: "dvdf", Got: lengthOfLongestSubstring("dvdf"), Expected: 3},
		{Input: "", Got: lengthOfLongestSubstring(""), Expected: 0},
		{Input: "a", Got: lengthOfLongestSubstring("a"), Expected: 1},
	})
}
