package main

import "leetcode/utils"

// We scan the grid looking for 1s
// When found, we perform a DFS to find the entire island
// To avoid visiting the same cell again, we mark the cell as visited by setting it to 0
// After the DFS, we increment the count of islands

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	count := 0

	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				dfs(&grid, m, n, i, j)
				count++
			}
		}
	}

	return count
}

func dfs(grid *[][]byte, m, n, i, j int) {
	if (*grid)[i][j] == '0' {
		return
	}

	(*grid)[i][j] = '0'
	if i > 0 {
		dfs(grid, m, n, i-1, j)
	}
	if i < m-1 {
		dfs(grid, m, n, i+1, j)
	}
	if j > 0 {
		dfs(grid, m, n, i, j-1)
	}
	if j < n-1 {
		dfs(grid, m, n, i, j+1)
	}
}

func main() {
	grid1 := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	grid2 := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}

	utils.RunTests([]utils.TestCase[int]{
		{Input: grid1, Got: numIslands(grid1), Expected: 1},
		{Input: grid2, Got: numIslands(grid2), Expected: 3},
	})
}
