package main

import "leetcode/utils"

func maxDifference(s string) int {
	freqs := make(map[rune]int, 26)

	for i := 0; i < len(s); i++ {
		freqs[rune(s[i])]++
	}

	max_odd := 0
	min_even := 100

	for _, freq := range freqs {
		if freq%2 == 0 && freq < min_even {
			min_even = freq
		} else {
			if freq%2 == 1 && freq > max_odd {
				max_odd = freq
			}
		}
	}

	return max_odd - min_even
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: "aaaaabbc", Got: maxDifference("aaaaabbc"), Expected: 3},
		{Input: "abcabcab", Got: maxDifference("abcabcab"), Expected: 1},
		{Input: "ililm", Got: maxDifference("ililm"), Expected: -1},
	})
}
