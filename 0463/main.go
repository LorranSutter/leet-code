package main

import "fmt"

// The idea here is to iterate through each cell in the grid and count the number of edges that are not connected to the water
// We start with the maximum possible perimeter (4 times the number of cells)
// For a water cell, we subtract 4 from the perimeter
// For an island cell, we subtract 1 for each adjacent island cell (up, down, left, right)
func islandPerimeter1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	perimeter := m * n * 4

	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				perimeter -= 4
				continue
			}
			if i > 0 && grid[i-1][j] == 1 {
				perimeter--
			}
			if i < m-1 && grid[i+1][j] == 1 {
				perimeter--
			}
			if j > 0 && grid[i][j-1] == 1 {
				perimeter--
			}
			if j < n-1 && grid[i][j+1] == 1 {
				perimeter--
			}
		}
	}
	return perimeter
}

// An optimized version of the above solution
// Instead of checking all four directions for each island cell, we only check the top and left cells
// This is because when we reach the adjacent cells later in the iteration, we will account for the shared edges
func islandPerimeter2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	perimeter := m * n * 4

	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				perimeter -= 4
				continue
			}
			if i > 0 && grid[i-1][j] == 1 {
				perimeter -= 2
			}
			if j > 0 && grid[i][j-1] == 1 {
				perimeter -= 2
			}
		}
	}
	return perimeter
}

func main() {
	fmt.Println("Solution 01")
	fmt.Println(islandPerimeter1([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}) == 16)
	fmt.Println(islandPerimeter1([][]int{{1}}) == 4)
	fmt.Println(islandPerimeter1([][]int{{1, 0}}) == 4)
	fmt.Println(islandPerimeter1([][]int{{1}, {1}, {1}, {0}, {1}, {1}}) == 14)

	fmt.Println()

	fmt.Println("Solution 02")
	fmt.Println(islandPerimeter2([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}) == 16)
	fmt.Println(islandPerimeter2([][]int{{1}}) == 4)
	fmt.Println(islandPerimeter2([][]int{{1, 0}}) == 4)
	fmt.Println(islandPerimeter2([][]int{{1}, {1}, {1}, {0}, {1}, {1}}) == 14)
}
