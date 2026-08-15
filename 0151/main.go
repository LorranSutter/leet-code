package main

import (
	"strings"

	"leetcode/utils"
)

func reverseWords(s string) string {
	var word strings.Builder
	reversed := ""

	for _, letter := range s {
		if letter != ' ' {
			word.WriteRune(letter)
		} else if word.Len() > 0 {
			word.WriteString(" " + reversed)
			reversed = word.String()
			word.Reset()
		}
	}

	if word.Len() > 0 {
		word.WriteString(" " + reversed)
		reversed = word.String()
	}

	return reversed[:len(reversed)-1]
}

func main() {
	utils.RunTests([]utils.TestCase[string]{
		{Input: "the sky is blue", Got: reverseWords("the sky is blue"), Expected: "blue is sky the"},
		{Input: "  hello world  ", Got: reverseWords("  hello world  "), Expected: "world hello"},
		{Input: "a good   example", Got: reverseWords("a good   example"), Expected: "example good a"},
		{Input: "a", Got: reverseWords("a"), Expected: "a"},
		{Input: "a  ", Got: reverseWords("a  "), Expected: "a"},
		{Input: "  a", Got: reverseWords("  a"), Expected: "a"},
	})
}
