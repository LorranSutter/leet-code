package main

import "fmt"

// 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0
// 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0
// 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0
// 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0
// 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1
//                |              |

// The idea is to find the longest sequence of ones using all flips initially
// Then, if we find a new 0 on the right, we can only flip it if we discard another 0 on the left

func longestOnes(nums []int, k int) int {
	left := 0
	longest := 0
	flips := k

	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == 1 {
			continue
		} else if flips > 0 {
			flips--
		} else {
			// Max flips reached, check if we got a longer sequence
			if longest < i-left {
				longest = i - left
			}

			// Move the left pointer to the next 0
			for left < i && nums[left] == 1 {
				left++
			}
			left++
		}
	}

	if longest < i-left {
		longest = i - left
	}

	return longest
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2) == 6)
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3) == 10)
	fmt.Println(longestOnes([]int{0, 0, 0, 1}, 4) == 4)
	fmt.Println(longestOnes([]int{0, 0, 1, 1}, 1) == 3)
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 1, 0, 0}, 0) == 3)
	fmt.Println(longestOnes([]int{0, 0, 0, 0}, 0) == 0)
}
