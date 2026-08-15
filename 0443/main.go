package main

import (
	"leetcode/utils"
	"strconv"
)

func compress(chars []byte) int {
	current := chars[0]
	count := 1
	char_id := 0
	for i := 1; i < len(chars); i++ {
		if chars[i] == current {
			count++
		} else {
			chars[char_id] = current
			current = chars[i]
			char_id++

			if count > 1 {
				count_str := strconv.Itoa(count)
				for _, digit := range count_str {
					chars[char_id] = byte(digit)
					char_id++
				}
				count = 1
			}
		}
	}

	chars[char_id] = current
	char_id++
	if count > 1 {
		count_str := strconv.Itoa(count)
		for _, digit := range count_str {
			chars[char_id] = byte(digit)
			char_id++
		}
	}

	return char_id
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, Got: compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}), Expected: 6},
		{Input: []byte{'a'}, Got: compress([]byte{'a'}), Expected: 1},
		{Input: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, Got: compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}), Expected: 4},
		{Input: []byte{'a', 'a', 'b', 'b', 'b', 'a', 'b', 'b', 'b', 'a', 'a', 'b', 'b', 'a'}, Got: compress([]byte{'a', 'a', 'b', 'b', 'b', 'a', 'b', 'b', 'b', 'a', 'a', 'b', 'b', 'a'}), Expected: 12},
	})
}
