package main

import "fmt"

func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)
	if nLen > hLen {
		return -1
	}

	k := 0
	for i := 0; i < hLen; i++ {
		k = 0
		for j := i; j < i+nLen && j < hLen; j++ {
			if haystack[j] != needle[k] {
				break
			}
			k++
		}

		if k == nLen {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad") == 0)
	fmt.Println(strStr("leetcode", "leeto") == -1)
	fmt.Println(strStr("mississippi", "issip") == 4)
	fmt.Println(strStr("mississippi", "issipi") == -1)
	fmt.Println(strStr("aaa", "aaaa") == -1)
}
