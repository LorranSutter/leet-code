package main

import "fmt"

// The idea here is just to keep track of the remainders until we find one that is 0.
// We can eliminate any multiple of 2 and 5, because these can't divide numbers ending in 1.

// We know that 111 = (11*10 + 1), 1111 = (111*10 + 1), and so on. We use that to iterate
// But we just have to keep the remainder of k because (a % k + b % k) % k == (a + b) % k
// e.g. 111 % 3 == (11*10 + 1) % 3 == ((11 % 3)*(10 % 3) + 1 % 3) % 3

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	remainder := 1
	count := 1
	for remainder%k != 0 {
		count++
		remainder = (remainder*10 + 1) % k
	}

	return count
}

func main() {
	fmt.Println(smallestRepunitDivByK(1) == 1)
	fmt.Println(smallestRepunitDivByK(2) == -1)
	fmt.Println(smallestRepunitDivByK(3) == 3)
	fmt.Println(smallestRepunitDivByK(4) == -1)
	fmt.Println(smallestRepunitDivByK(5) == -1)
	fmt.Println(smallestRepunitDivByK(5) == -1)
	fmt.Println(smallestRepunitDivByK(9) == 9)
	fmt.Println(smallestRepunitDivByK(11) == 2)
	fmt.Println(smallestRepunitDivByK(13) == 6)
	fmt.Println(smallestRepunitDivByK(17) == 16)
}
