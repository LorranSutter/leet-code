package main

import "fmt"

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
	fmt.Println(largeGroupPositions("aaa"))
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
	fmt.Println(largeGroupPositions("bababbaaab"))
}
