package main

import "fmt"

// 0 ms / 4.56 MB
func threeConsecutiveOdds(arr []int) bool {
	countOdds := 0
	for _, num := range arr {
		if num%2 != 0 {
			countOdds++
			if countOdds >= 3 {
				return true
			}
		} else {
			countOdds = 0
		}
	}
	return false
}

func main() {
	fmt.Println(threeConsecutiveOdds([]int{2, 6, 4, 1}) == false)
	fmt.Println(threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}) == true)
	fmt.Println(threeConsecutiveOdds([]int{1}) == false)
}
