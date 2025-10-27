package main

import "fmt"

func numberOfBeams(bank []string) int {
	lasers, prev, current := 0, 0, 0

	for _, floor := range bank {
		current = 0
		for _, device := range floor {
			if device == '1' {
				current++
			}
		}

		if prev == 0 {
			prev = current
			continue
		}

		if prev > 0 && current > 0 {
			lasers += prev * current
			prev, current = current, 0
			continue
		}
	}

	return lasers
}

func main() {
	fmt.Println(numberOfBeams([]string{"011001", "000000", "010100", "001000"}) == 8)
	fmt.Println(numberOfBeams([]string{"000", "111", "000"}) == 0)
	fmt.Println(numberOfBeams([]string{"010", "111", "010"}) == 6)
	fmt.Println(numberOfBeams([]string{"010", "001", "010"}) == 2)
}
