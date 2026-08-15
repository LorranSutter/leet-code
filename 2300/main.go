package main

import (
	"leetcode/utils"
	"sort"
)

// 28 ms / 14.56 MB
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	num_success := make([]int, len(spells))

	for i, spell := range spells {
		num_success[i] = search(potions, spell, success)
	}

	return num_success
}

func search(potions []int, spell int, success int64) int {
	left, mid := 0, 0
	right := len(potions) - 1

	for left <= right {
		mid = left + (right-left)/2

		if int64(potions[mid]*spell) < success {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if int64(potions[mid]*spell) > success {
		return len(potions) - mid
	} else if int64(potions[mid]*spell) < success {
		return len(potions) - mid - 1
	}

	for mid > 0 && potions[mid-1] == potions[mid] {
		mid--
	}

	return len(potions) - mid
}

func main() {
	utils.RunTests([]utils.TestCase[[]int]{
		{
			Input:    []any{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, int64(7)},
			Got:      successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7),
			Expected: []int{4, 0, 3},
		},
		{
			Input:    []any{[]int{5, 1, 3}, []int{1, 2, 3, 3, 4, 5}, int64(7)},
			Got:      successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 3, 4, 5}, 7),
			Expected: []int{5, 0, 4},
		},
		{
			Input:    []any{[]int{3, 1, 2}, []int{8, 5, 8}, int64(16)},
			Got:      successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16),
			Expected: []int{2, 0, 2},
		},
		{
			Input: []any{
				[]int{40, 11, 24, 28, 40, 22, 26, 38, 28, 10, 31, 16, 10, 37, 13, 21, 9, 22, 21, 18, 34, 2, 40, 40, 6, 16, 9, 14, 14, 15, 37, 15, 32, 4, 27, 20, 24, 12, 26, 39, 32, 39, 20, 19, 22, 33, 2, 22, 9, 18, 12, 5},
				[]int{31, 40, 29, 19, 27, 16, 25, 8, 33, 25, 36, 21, 7, 27, 40, 24, 18, 26, 32, 25, 22, 21, 38, 22, 37, 34, 15, 36, 21, 22, 37, 14, 31, 20, 36, 27, 28, 32, 21, 26, 33, 37, 27, 39, 19, 36, 20, 23, 25, 39, 40},
				int64(600),
			},
			Got: successfulPairs(
				[]int{40, 11, 24, 28, 40, 22, 26, 38, 28, 10, 31, 16, 10, 37, 13, 21, 9, 22, 21, 18, 34, 2, 40, 40, 6, 16, 9, 14, 14, 15, 37, 15, 32, 4, 27, 20, 24, 12, 26, 39, 32, 39, 20, 19, 22, 33, 2, 22, 9, 18, 12, 5},
				[]int{31, 40, 29, 19, 27, 16, 25, 8, 33, 25, 36, 21, 7, 27, 40, 24, 18, 26, 32, 25, 22, 21, 38, 22, 37, 34, 15, 36, 21, 22, 37, 14, 31, 20, 36, 27, 28, 32, 21, 26, 33, 37, 27, 39, 19, 36, 20, 23, 25, 39, 40},
				600,
			),
			Expected: []int{48, 0, 32, 37, 48, 22, 33, 47, 37, 0, 43, 6, 0, 46, 0, 21, 0, 22, 21, 14, 46, 0, 48, 48, 0, 6, 0, 0, 0, 3, 46, 3, 45, 0, 34, 20, 32, 0, 33, 47, 45, 47, 20, 18, 22, 45, 0, 22, 0, 14, 0, 0},
		},
	})
}
