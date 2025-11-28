package main

import "fmt"

// The idea here is find the area of each island and return the biggest one
// We scan the grid looking for 1s
// When found, we perform a DFS to find the entire island and count its area
// To avoid visiting the same cell again, we mark the cell as visited by setting it to 0

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxArea, islandArea := 0, 0

	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				dfs(&grid, m, n, i, j, &islandArea)
				if islandArea > maxArea {
					maxArea = islandArea
				}
				islandArea = 0
			}
		}
	}

	return maxArea
}

func dfs(grid *[][]int, m, n, i, j int, count *int) {
	if (*grid)[i][j] == 0 {
		return
	}

	(*count)++
	(*grid)[i][j] = 0

	if i > 0 {
		dfs(grid, m, n, i-1, j, count)
	}
	if i < m-1 {
		dfs(grid, m, n, i+1, j, count)
	}
	if j > 0 {
		dfs(grid, m, n, i, j-1, count)
	}
	if j < n-1 {
		dfs(grid, m, n, i, j+1, count)
	}
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}) == 6)
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 0, 0, 0, 0, 0, 0}}) == 0)
	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}) == 4)
}
