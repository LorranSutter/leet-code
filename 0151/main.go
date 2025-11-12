package main

import (
	"fmt"
	"strings"
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
	fmt.Println(reverseWords("the sky is blue") == "blue is sky the")
	fmt.Println(reverseWords("  hello world  ") == "world hello")
	fmt.Println(reverseWords("a good   example") == "example good a")
	fmt.Println(reverseWords("a") == "a")
	fmt.Println(reverseWords("a  ") == "a")
	fmt.Println(reverseWords("  a") == "a")
}
