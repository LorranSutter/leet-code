package main

import (
	"leetcode/utils"
	"strconv"
)

func maximum69Number(num int) int {
	numStr := strconv.Itoa(num)

	for i := range numStr {
		if numStr[i] == '6' {
			numStr = numStr[:i] + "9" + numStr[i+1:]
			newNum, _ := strconv.Atoi(numStr)
			return newNum
		}
	}

	return num
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: 9669, Got: maximum69Number(9669), Expected: 9969},
		{Input: 9996, Got: maximum69Number(9996), Expected: 9999},
		{Input: 9999, Got: maximum69Number(9999), Expected: 9999},
		{Input: 6666, Got: maximum69Number(6666), Expected: 9666},
		{Input: 9, Got: maximum69Number(9), Expected: 9},
		{Input: 6, Got: maximum69Number(6), Expected: 9},
	})
}
