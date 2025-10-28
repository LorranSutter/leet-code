package main

import "fmt"

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
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2) == false)
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0}, 4) == true)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1}, 2) == true)
	fmt.Println(canPlaceFlowers([]int{0}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{0}, 0) == true)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2) == false)
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 1, 0, 1}, 1) == false)
}
