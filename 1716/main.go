package main

import "fmt"

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
	fmt.Println(totalMoney(4) == 10)
	fmt.Println(totalMoney(10) == 37)
	fmt.Println(totalMoney(20) == 96)
	fmt.Println(totalMoney(21) == 105)
	fmt.Println(totalMoney(22) == 109)
}
