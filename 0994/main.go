package main

import "fmt"

// The idea here is to perform a variation of the BFS algorithm to solve the problem
// Instead of starting in a single cell, we start from all the rotten oranges
// We mark all surrounding fresh oranges to rotten and to be visited, and we remove the already visited rotten oranges
// We do the same process if there is still rotten oranges to be visited, or no more fresh oranges

const (
	FRESH  = 1
	ROTTEN = 2
)

func orangesRotting(grid [][]int) int {
	minutes, fresh := 0, 0
	toVisit := [][2]int{}
	m, n := len(grid), len(grid[0])

	// Compute the initial number of fresh oranges and position of rotten oranges to visit
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case FRESH:
				fresh++
			case ROTTEN:
				toVisit = append(toVisit, [2]int{i, j})
			}
		}
	}

	for fresh > 0 && len(toVisit) > 0 {
		for _, rotten := range toVisit {
			x, y := rotten[0], rotten[1]

			if x > 0 && grid[x-1][y] == FRESH {
				grid[x-1][y] = ROTTEN
				toVisit = append(toVisit, [2]int{x - 1, y})
				fresh--
			}
			if x < m-1 && grid[x+1][y] == FRESH {
				grid[x+1][y] = ROTTEN
				toVisit = append(toVisit, [2]int{x + 1, y})
				fresh--
			}
			if y > 0 && grid[x][y-1] == FRESH {
				grid[x][y-1] = ROTTEN
				toVisit = append(toVisit, [2]int{x, y - 1})
				fresh--
			}
			if y < n-1 && grid[x][y+1] == FRESH {
				grid[x][y+1] = ROTTEN
				toVisit = append(toVisit, [2]int{x, y + 1})
				fresh--
			}

			toVisit = toVisit[1:]
		}
		minutes++
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}) == 4)
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}) == -1)
	fmt.Println(orangesRotting([][]int{{20, 2}}) == 0)
}
