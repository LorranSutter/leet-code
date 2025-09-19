package main

import (
	"fmt"
	"math"
)

func divideString(s string, k int, fill byte) []string {
	slen := len(s)
	numGroups := int(math.Ceil(float64(slen) / float64(k)))
	result := make([]string, numGroups)

	for i := range numGroups {
		if i*k+k >= slen {
			result[i] = s[i*k:]
			break
		}
		result[i] = s[i*k : i*k+k]
	}

	toFill := k - len(result[numGroups-1])
	if toFill > 0 {
		fill := string(fill)
		filling := ""
		for range toFill {
			filling += fill
		}

		result[numGroups-1] += filling
	}

	return result
}

func main() {
	fmt.Println(divideString("abcdefghi", 3, 'x'))
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}
