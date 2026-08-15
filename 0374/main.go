package main

import "leetcode/utils"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var picked int

func guess(num int) int {
	if num > picked {
		return -1
	} else if num < picked {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		result := guess(mid)
		switch result {
		case 0:
			return mid
		case -1:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return 0
}

func main() {
	picked = 6
	tc1 := utils.TestCase[int]{Input: 10, Got: guessNumber(10), Expected: picked}

	picked = 1
	tc2 := utils.TestCase[int]{Input: 1, Got: guessNumber(1), Expected: picked}

	picked = 1
	tc3 := utils.TestCase[int]{Input: 2, Got: guessNumber(2), Expected: picked}

	picked = 123542
	tc4 := utils.TestCase[int]{Input: 34256347645634, Got: guessNumber(34256347645634), Expected: picked}

	utils.RunTests([]utils.TestCase[int]{tc1, tc2, tc3, tc4})
}
