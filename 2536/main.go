package main

import (
	"fmt"
	"leetcode/utils"
)

// The idea is to use a 2D Difference Array Technique
// 1. For each query, we just have to know what happens with the beging and ending indices
// 2. We will use a prefix sum for each row and column after we calculate the boundaries
// 3. Iterate over queries
//       Query -> [0, 0, 1, 1]
// 	3.1. We always increment the top left corner
//       1 0 0 0
//       0 0 0 0
//       0 0 0 0
//       0 0 0 0
// 	3.2. We decrement the right of top right corner and bottom of the bottom left corner, because when we prefix sum after we don't carry the sum after the query limit
//       1 0 0 0     1 0 -1 0
//       0 0 0 0 ->  0 0  0 0
//       0 0 0 0    -1 0  0 0
//       0 0 0 0     0 0  0 0
//  3.3. Intuitively we should also decrement the right and bottom of the right bottom corner
//       However, it causes a problem when we prefix sum afer, because we double decrement the corner after the query
//        1 0 -1 0      1  0 -1 0     1  1  0  0
//        0 0  0 0  ->  0  0 -1 0 ->  1  1 -1 -1
//       -1 0  0 0     -1 -1  0 0     0 -1 -3 -3
//        0 0  0 0      0  0  0 0     0  0 -2 -2
//  3.4. Instead, we increment the corner below the bottom left corner to avoid double decrement
//        1 0 -1 0      1  0 -1 0     1 1 0 0
//        0 0  0 0  ->  0  0  0 0 ->  1 1 0 0
//       -1 0  0 0     -1  0  1 0     0 0 0 0
//        0 0  0 0      0  0  0 0     0 0 0 0
// 4. After all queries, we just need to prefix sum the row and column

func rangeAddQueries(n int, queries [][]int) [][]int {
	m := make([][]int, n)
	for i := range n {
		m[i] = make([]int, n)
	}

	for _, q := range queries {
		fmt.Println(q)
		// Top left corner
		m[q[0]][q[1]]++

		// Bottom right corner
		if q[2] < n-1 && q[3] < n-1 {
			m[q[2]+1][q[3]+1]++
		}

		// Top right corner
		if q[3] < n-1 {
			m[q[0]][q[3]+1]--
		}

		// Bottom left corner
		if q[2] < n-1 {
			m[q[2]+1][q[1]]--
		}
	}

	// Prefix sum rows
	for i := range n {
		for j := 1; j < n; j++ {
			m[i][j] += m[i][j-1]
		}
	}

	// Prefix sum columns
	for i := 1; i < n; i++ {
		for j := range n {
			m[i][j] += m[i-1][j]
		}
	}

	return m
}

func main() {
	result := rangeAddQueries(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}})
	fmt.Println(utils.DeepEqualMatrix(result, [][]int{{1, 1, 0}, {1, 2, 1}, {0, 1, 1}}))

	result = rangeAddQueries(2, [][]int{{0, 0, 1, 1}})
	fmt.Println(utils.DeepEqualMatrix(result, [][]int{{1, 1}, {1, 1}}))
}
