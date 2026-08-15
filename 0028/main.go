package main

import (
	"leetcode/utils"
)

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
	utils.RunTests([]utils.TestCase[int]{
		{Input: []string{"sadbutsad", "sad"}, Got: strStr("sadbutsad", "sad"), Expected: 0},
		{Input: []string{"leetcode", "leeto"}, Got: strStr("leetcode", "leeto"), Expected: -1},
		{Input: []string{"mississippi", "issip"}, Got: strStr("mississippi", "issip"), Expected: 4},
		{Input: []string{"mississippi", "issipi"}, Got: strStr("mississippi", "issipi"), Expected: -1},
		{Input: []string{"aaa", "aaaa"}, Got: strStr("aaa", "aaaa"), Expected: -1},
	})
}
