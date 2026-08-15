package main

import "leetcode/utils"

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
	utils.RunTests([]utils.TestCase[string]{
		{Input: "IceCreAm", Got: reverseVowels("IceCreAm"), Expected: "AceCreIm"},
		{Input: "leetcode", Got: reverseVowels("leetcode"), Expected: "leotcede"},
		{Input: "aeiou", Got: reverseVowels("aeiou"), Expected: "uoiea"},
		{Input: "aeiJou", Got: reverseVowels("aeiJou"), Expected: "uoiJea"},
	})
}
