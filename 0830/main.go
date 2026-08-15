package main

import "leetcode/utils"

const LARGE_SIZE int = 3

func largeGroupPositions(s string) [][]int {
	slen := len(s)
	var result [][]int

	i, j := 0, 0
	for i < slen {
		for j < slen && s[i] == s[j] {
			j++
		}
		if j-i >= LARGE_SIZE {
			result = append(result, []int{i, j - 1})
		}
		i = j
	}

	return result
}

func main() {
	utils.RunTests([]utils.TestCase[[][]int]{
		{Input: "aaa", Got: largeGroupPositions("aaa"), Expected: [][]int{{0, 2}}},
		{Input: "abbxxxxzzy", Got: largeGroupPositions("abbxxxxzzy"), Expected: [][]int{{3, 6}}},
		{Input: "abcdddeeeeaabbbcd", Got: largeGroupPositions("abcdddeeeeaabbbcd"), Expected: [][]int{{3, 5}, {6, 9}, {12, 14}}},
		{Input: "bababbaaab", Got: largeGroupPositions("bababbaaab"), Expected: [][]int{{6, 8}}},
	})
}
