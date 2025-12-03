package main

import "fmt"

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
	fmt.Println(guessNumber(10) == picked)
	picked = 1
	fmt.Println(guessNumber(1) == picked)
	picked = 1
	fmt.Println(guessNumber(2) == picked)
	picked = 123542
	fmt.Println(guessNumber(34256347645634) == picked)
}
