package main

import "fmt"

// 170 ms / 9.93 MB
func equalPairs1(grid [][]int) int {
	n := len(grid)
	row_hash := map[string]int{}

	row_str := ""
	for _, row := range grid {
		for _, num := range row {
			row_str += fmt.Sprint(num, ",")
		}
		row_hash[row_str]++
		row_str = ""
	}

	count := 0
	col_str := ""
	for i := range n {
		for j := range n {
			col_str += fmt.Sprint(grid[j][i], ",")
		}
		count += row_hash[col_str]
		col_str = ""
	}

	return count
}

// 2 ms / 9.93 MB
func equalPairs2(grid [][]int) int {
	n := len(grid)
	row_hash := make(map[[200]int]int)

	for _, row := range grid {
		var local_row [200]int
		copy(local_row[:n], row)
		row_hash[local_row]++
	}

	count := 0
	for i := range n {
		var local_col [200]int
		for j := range n {
			local_col[j] = grid[j][i]
		}
		count += row_hash[local_col]
	}

	return count
}

func main() {
	fmt.Println("Solution 01")
	fmt.Println(equalPairs1([][]int{{13, 13}, {13, 13}}) == 4)
	fmt.Println(equalPairs1([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}) == 1)
	fmt.Println(equalPairs1([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}) == 3)

	fmt.Println()

	fmt.Println("Solution 02")
	fmt.Println(equalPairs2([][]int{{13, 13}, {13, 13}}) == 4)
	fmt.Println(equalPairs2([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}) == 1)
	fmt.Println(equalPairs2([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}) == 3)
}
