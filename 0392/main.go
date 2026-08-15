package main

import "leetcode/utils"

func isSubsequence(s string, t string) bool {
	sLen := len(s)
	if sLen == 0 {
		return true
	}

	p := 0
	for i := range t {
		if s[p] == t[i] {
			p++
			if p == sLen {
				return true
			}
		}
	}

	return false
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []string{"ace", "abcde"}, Got: isSubsequence("ace", "abcde"), Expected: true},
		{Input: []string{"aec", "abcde"}, Got: isSubsequence("aec", "abcde"), Expected: false},
		{Input: []string{"abc", "ahbgdc"}, Got: isSubsequence("abc", "ahbgdc"), Expected: true},
		{Input: []string{"axc", "ahbgdc"}, Got: isSubsequence("axc", "ahbgdc"), Expected: false},
		{Input: []string{"aaa", "ahbgdc"}, Got: isSubsequence("aaa", "ahbgdc"), Expected: false},
		{Input: []string{"aaa", "ahbgaa"}, Got: isSubsequence("aaa", "ahbgaa"), Expected: true},
		{Input: []string{"", "ahbgdc"}, Got: isSubsequence("", "ahbgdc"), Expected: true},
	})
}
