package main

import (
	"leetcode/utils"
)

// 2 ms / 3.97 MB
func mergeAlternately1(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	p1, p2 := 0, 0
	letters := make([]byte, len1+len2)

	for i := range letters {
		if i%2 == 0 {
			if p1 < len1 {
				letters[i] = word1[p1]
				p1++
			} else {
				letters[i] = word2[p2]
				p2++
			}
		} else {
			if p2 < len2 {
				letters[i] = word2[p2]
				p2++
			} else {
				letters[i] = word1[p1]
				p1++
			}
		}
	}

	return string(letters)
}

// 2 ms / 3.97 MB
func mergeAlternately2(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	var letters []byte

	for i := 0; i < len1 || i < len2; {
		if i < len1 {
			letters = append(letters, word1[i])
		}
		if i < len2 {
			letters = append(letters, word2[i])
		}
		i++
	}

	return string(letters)
}

func main() {
	utils.RunTests([]utils.TestCase[string]{
		{Input: []any{"abc", "pqr"}, Got: mergeAlternately1("abc", "pqr"), Expected: "apbqcr"},
		{Input: []any{"ab", "pqrs"}, Got: mergeAlternately1("ab", "pqrs"), Expected: "apbqrs"},
		{Input: []any{"abcd", "pq"}, Got: mergeAlternately1("abcd", "pq"), Expected: "apbqcd"},
	})

	utils.RunTests([]utils.TestCase[string]{
		{Input: []any{"abc", "pqr"}, Got: mergeAlternately2("abc", "pqr"), Expected: "apbqcr"},
		{Input: []any{"ab", "pqrs"}, Got: mergeAlternately2("ab", "pqrs"), Expected: "apbqrs"},
		{Input: []any{"abcd", "pq"}, Got: mergeAlternately2("abcd", "pq"), Expected: "apbqcd"},
	})
}
