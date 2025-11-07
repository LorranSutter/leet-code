package main

import "fmt"

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
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}) == true)
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}) == true)
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) == false)
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}) == true)
}
