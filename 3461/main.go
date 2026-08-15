package main

import "leetcode/utils"

// 0 ms / 4.3 MB
func hasSameDigits(s string) bool {
	maxLen := len(s)
	nums := make([]int, maxLen)

	for i := range s {
		nums[i] = int(s[i] - '0')
	}

	for {
		for i := 0; i < maxLen-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}

		maxLen--
		if maxLen == 2 {
			return nums[0] == nums[1]
		}
	}
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: "3902", Got: hasSameDigits("3902"), Expected: true},
		{Input: "34789", Got: hasSameDigits("34789"), Expected: false},
		{Input: "4867716181911413979960155821878", Got: hasSameDigits("4867716181911413979960155821878"), Expected: true},
		{Input: "472371665226339635424494604455723348326857778972516313101962164", Got: hasSameDigits("472371665226339635424494604455723348326857778972516313101962164"), Expected: true},
	})
}
