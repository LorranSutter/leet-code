package main

import "leetcode/utils"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	lenFlowerbed := len(flowerbed)

	if lenFlowerbed == 1 {
		return flowerbed[0] == 0
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		n--
	}

	for i := 1; n > 0 && i < lenFlowerbed-1; {
		if flowerbed[i] == 0 {
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
				i++
			}
		}
		i++
	}

	if n == 1 && flowerbed[lenFlowerbed-2] == 0 && flowerbed[lenFlowerbed-1] == 0 {
		return true
	}

	return n == 0
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: []any{[]int{1, 0, 0, 0, 1}, 1}, Got: canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1), Expected: true},
		{Input: []any{[]int{1, 0, 0, 0, 1}, 2}, Got: canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2), Expected: false},
		{Input: []any{[]int{0, 0, 1, 0, 0}, 1}, Got: canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1), Expected: true},
		{Input: []any{[]int{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0}, 4}, Got: canPlaceFlowers([]int{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0}, 4), Expected: true},
		{Input: []any{[]int{1, 0, 0, 0, 0, 0, 1}, 2}, Got: canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1}, 2), Expected: true},
		{Input: []any{[]int{0}, 1}, Got: canPlaceFlowers([]int{0}, 1), Expected: true},
		{Input: []any{[]int{0}, 0}, Got: canPlaceFlowers([]int{0}, 0), Expected: true},
		{Input: []any{[]int{1, 0, 0, 0, 0, 1}, 2}, Got: canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2), Expected: false},
		{Input: []any{[]int{1, 0, 1, 0, 1, 0, 1}, 1}, Got: canPlaceFlowers([]int{1, 0, 1, 0, 1, 0, 1}, 1), Expected: false},
	})
}
