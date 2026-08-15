package main

import (
	"leetcode/utils"
)

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	substring := ""
	for i := 0; i < len(s); i++ {
		substring += string(s[i])
		if words[substring] {
			substring = ""
		}
	}

	if substring == "" {
		return true
	}

	return false
}

func main() {
	// TODO Implement solution
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []any{"leetcode", []string{"leet", "code"}}, Got: wordBreak("leetcode", []string{"leet", "code"}), Expected: true},
		{Input: []any{"applepenapple", []string{"apple", "pen"}}, Got: wordBreak("applepenapple", []string{"apple", "pen"}), Expected: true},
		{Input: []any{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, Got: wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), Expected: false},
		{Input: []any{"aaaaaaa", []string{"aaaa", "aaa"}}, Got: wordBreak("aaaaaaa", []string{"aaaa", "aaa"}), Expected: true},
	})
}
