package main

import "fmt"

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
	fmt.Println(isSubsequence("ace", "abcde") == true)
	fmt.Println(isSubsequence("aec", "abcde") == false)
	fmt.Println(isSubsequence("abc", "ahbgdc") == true)
	fmt.Println(isSubsequence("axc", "ahbgdc") == false)
	fmt.Println(isSubsequence("aaa", "ahbgdc") == false)
	fmt.Println(isSubsequence("aaa", "ahbgaa") == true)
	fmt.Println(isSubsequence("", "ahbgdc") == true)
}
