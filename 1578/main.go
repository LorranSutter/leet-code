package main

import "fmt"

func minCost(colors string, neededTime []int) int {
	cost := 0
	current_color := colors[0]
	chunk_cost := neededTime[0]
	max_chunk_cost := neededTime[0]
	for i := 1; i < len(neededTime); i++ {
		if current_color != colors[i] {
			cost += (chunk_cost - max_chunk_cost)
			chunk_cost, max_chunk_cost = neededTime[i], neededTime[i]
			current_color = colors[i]
			continue
		}

		chunk_cost += neededTime[i]
		if neededTime[i] > max_chunk_cost {
			max_chunk_cost = neededTime[i]
		}
	}

	cost += (chunk_cost - max_chunk_cost)

	return cost
}

func main() {
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}) == 3)
	fmt.Println(minCost("abc", []int{1, 2, 3}) == 0)
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}) == 2)
	fmt.Println(minCost("baaaabaa", []int{1, 1, 5, 1, 2, 6, 7, 1}) == 5)
	fmt.Println(minCost("bbbaaa", []int{4, 9, 3, 8, 8, 9}) == 23)
}
