package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	fmt.Println(string(chars))

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

	fmt.Println(string(chars))
	fmt.Println(char_id)

	return char_id
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}) == 6)
	fmt.Println(compress([]byte{'a'}) == 1)
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}) == 4)
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'b', 'a', 'b', 'b', 'b', 'a', 'a', 'b', 'b', 'a'}) == 12)
}
