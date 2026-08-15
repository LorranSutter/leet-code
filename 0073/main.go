package main

import (
	"leetcode/utils"
)

// 0 ms / 7.94 MB
func setZeroes1(matrix [][]int) [][]int {
	var zeroesPos [][]int

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				zeroesPos = append(zeroesPos, [][]int{{i, j}}...)
			}
		}
	}

	for _, pos := range zeroesPos {
		for i := range matrix {
			matrix[i][pos[1]] = 0
		}
		for j := range matrix[0] {
			matrix[pos[0]][j] = 0
		}
	}

	return matrix
}

// 0 ms / 7.88 MB
func setZeroes2(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	var zeroesPos []int

	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				zeroesPos = append(zeroesPos, i*n+j)
			}
		}
	}

	for _, index := range zeroesPos {
		index_i := index / n
		index_j := index % n

		for i := range matrix {
			matrix[i][index_j] = 0
		}
		for j := range matrix[0] {
			matrix[index_i][j] = 0
		}
	}

	return matrix
}

// 0 ms / 7.81 MB
func setZeroes3(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	var firstRowHasZero, firstColumnHasZero bool

	// Check if any element in the first row is 0
	for i := range m {
		if matrix[i][0] == 0 {
			firstRowHasZero = true
		}
	}
	// Check if any element in the first column is 0
	for j := range n {
		if matrix[0][j] == 0 {
			firstColumnHasZero = true
		}
	}

	// Assign 0 to corresponding row 0 and column 0
	// of a found zero out of the first row and column
	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// Sets every 0 column
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// Sets every 0 row
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// Sets first column to zero
	if firstRowHasZero {
		for i := range m {
			matrix[i][0] = 0
		}
	}
	// Sets first row to zero
	if firstColumnHasZero {
		for j := range n {
			matrix[0][j] = 0
		}
	}

	return matrix
}

func main() {
	utils.RunTests([]utils.TestCase[[][]int]{
		{Input: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, Got: setZeroes1([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}), Expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{Input: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, Got: setZeroes1([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}), Expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{Input: [][]int{{4, 1, 2, 5}, {3, 0, 5, 2}, {1, 3, 1, 0}}, Got: setZeroes1([][]int{{4, 1, 2, 5}, {3, 0, 5, 2}, {1, 3, 1, 0}}), Expected: [][]int{{4, 0, 2, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{4, 1, 2, 5}, {3, 5, 5, 2}, {0, 0, 0, 0}}, Got: setZeroes1([][]int{{4, 1, 2, 5}, {3, 5, 5, 2}, {0, 0, 0, 0}}), Expected: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{0, 1, 2, 5}, {0, 5, 5, 2}, {0, 2, 4, 5}}, Got: setZeroes1([][]int{{0, 1, 2, 5}, {0, 5, 5, 2}, {0, 2, 4, 5}}), Expected: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{0, 1, 2}, {9, 0, 5}, {1, 2, 0}}, Got: setZeroes1([][]int{{0, 1, 2}, {9, 0, 5}, {1, 2, 0}}), Expected: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
	})

	utils.RunTests([]utils.TestCase[[][]int]{
		{Input: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, Got: setZeroes2([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}), Expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{Input: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, Got: setZeroes2([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}), Expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{Input: [][]int{{4, 1, 2, 5}, {3, 0, 5, 2}, {1, 3, 1, 0}}, Got: setZeroes2([][]int{{4, 1, 2, 5}, {3, 0, 5, 2}, {1, 3, 1, 0}}), Expected: [][]int{{4, 0, 2, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{4, 1, 2, 5}, {3, 5, 5, 2}, {0, 0, 0, 0}}, Got: setZeroes2([][]int{{4, 1, 2, 5}, {3, 5, 5, 2}, {0, 0, 0, 0}}), Expected: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{0, 1, 2, 5}, {0, 5, 5, 2}, {0, 2, 4, 5}}, Got: setZeroes2([][]int{{0, 1, 2, 5}, {0, 5, 5, 2}, {0, 2, 4, 5}}), Expected: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{0, 1, 2}, {9, 0, 5}, {1, 2, 0}}, Got: setZeroes2([][]int{{0, 1, 2}, {9, 0, 5}, {1, 2, 0}}), Expected: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
	})

	utils.RunTests([]utils.TestCase[[][]int]{
		{Input: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, Got: setZeroes3([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}), Expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{Input: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, Got: setZeroes3([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}), Expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{Input: [][]int{{4, 1, 2, 5}, {3, 0, 5, 2}, {1, 3, 1, 0}}, Got: setZeroes3([][]int{{4, 1, 2, 5}, {3, 0, 5, 2}, {1, 3, 1, 0}}), Expected: [][]int{{4, 0, 2, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{4, 1, 2, 5}, {3, 5, 5, 2}, {0, 0, 0, 0}}, Got: setZeroes3([][]int{{4, 1, 2, 5}, {3, 5, 5, 2}, {0, 0, 0, 0}}), Expected: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{0, 1, 2, 5}, {0, 5, 5, 2}, {0, 2, 4, 5}}, Got: setZeroes3([][]int{{0, 1, 2, 5}, {0, 5, 5, 2}, {0, 2, 4, 5}}), Expected: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{Input: [][]int{{0, 1, 2}, {9, 0, 5}, {1, 2, 0}}, Got: setZeroes3([][]int{{0, 1, 2}, {9, 0, 5}, {1, 2, 0}}), Expected: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{Input: [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}, Got: setZeroes3([][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}), Expected: [][]int{{0, 0, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
	})
}
