package main

import "fmt"

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
	fmt.Println("Solution 01")
	fmt.Println(mergeAlternately1("abc", "pqr") == "apbqcr")
	fmt.Println(mergeAlternately1("ab", "pqrs") == "apbqrs")
	fmt.Println(mergeAlternately1("abcd", "pq") == "apbqcd")

	fmt.Println()

	fmt.Println("Solution 02")
	fmt.Println(mergeAlternately2("abc", "pqr") == "apbqcr")
	fmt.Println(mergeAlternately2("ab", "pqrs") == "apbqrs")
	fmt.Println(mergeAlternately2("abcd", "pq") == "apbqcd")
}
