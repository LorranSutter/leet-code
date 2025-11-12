package main

import "fmt"

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	i := 0
	j := len(s) - 1
	b := []byte(s)

	for i < j {
		if vowels[s[i]] && vowels[s[j]] {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
		if !vowels[s[i]] {
			i++
		}
		if !vowels[s[j]] {
			j--
		}
	}

	return string(b)
}

func main() {
	fmt.Println(reverseVowels("IceCreAm") == "AceCreIm")
	fmt.Println(reverseVowels("leetcode") == "leotcede")
	fmt.Println(reverseVowels("aeiou") == "uoiea")
	fmt.Println(reverseVowels("aeiJou") == "uoiJea")
}
