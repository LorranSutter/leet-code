package main

import (
	"fmt"
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
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
	fmt.Println(maximum69Number(6666))
	fmt.Println(maximum69Number(9))
	fmt.Println(maximum69Number(6))
}
