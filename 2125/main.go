package main

import (
	"leetcode/utils"
)

func numberOfBeams(bank []string) int {
	lasers, prev, current := 0, 0, 0

	for _, floor := range bank {
		current = 0
		for _, device := range floor {
			if device == '1' {
				current++
			}
		}

		if current > 0 {
			lasers += prev * current
			prev = current
		}
	}

	return lasers
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []string{"011001", "000000", "010100", "001000"}, Got: numberOfBeams([]string{"011001", "000000", "010100", "001000"}), Expected: 8},
		{Input: []string{"000", "111", "000"}, Got: numberOfBeams([]string{"000", "111", "000"}), Expected: 0},
		{Input: []string{"010", "111", "010"}, Got: numberOfBeams([]string{"010", "111", "010"}), Expected: 6},
		{Input: []string{"010", "001", "010"}, Got: numberOfBeams([]string{"010", "001", "010"}), Expected: 2},
	})
}
