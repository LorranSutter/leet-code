package main

import (
	"leetcode/utils"
)

func totalMoney(n int) int {
	total := 0
	i := 0
	for {
		if n/7 > 0 {
			total += 7 * (i + 1 + 7 + i) / 2
			n -= 7
			i++
		} else {
			break
		}
	}

	n = n % 7
	if n > 0 {
		total += n * (i + 1 + n + i) / 2
	}

	return total
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: 4, Got: totalMoney(4), Expected: 10},
		{Input: 10, Got: totalMoney(10), Expected: 37},
		{Input: 20, Got: totalMoney(20), Expected: 96},
		{Input: 21, Got: totalMoney(21), Expected: 105},
		{Input: 22, Got: totalMoney(22), Expected: 109},
	})
}
