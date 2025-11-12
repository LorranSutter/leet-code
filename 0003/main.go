package main

import (
	"fmt"
	"strings"
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
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("abcdabcbb") == 4)
	fmt.Println(lengthOfLongestSubstring("abcabcdbabc") == 4)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("dvdf") == 3)
	fmt.Println(lengthOfLongestSubstring("") == 0)
	fmt.Println(lengthOfLongestSubstring("a") == 1)
}
