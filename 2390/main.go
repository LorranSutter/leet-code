package main

import "fmt"

func removeStars(s string) string {
	result := ""
	to_remove := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			to_remove++
		} else {
			if to_remove > 0 {
				to_remove--
			} else {
				result = string(s[i]) + result
			}
		}
	}

	return result
}

func main() {
	fmt.Println(removeStars("leet**cod*e") == "lecoe")
	fmt.Println(removeStars("erase*****") == "")
}
