package main

func rotate(matrix [][]int) {
	// Transponse and reverse columns
	// 1 2 3    1 4 7   0 0 1   7 4 1
	// 4 5 6 -> 2 5 8 X 0 1 0 = 8 5 2
	// 7 8 9    3 6 9   1 0 0.  9 6 3

	// Transpose matrix
	for i := range matrix {
		for j := 0; j < i+1; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	mLen := len(matrix)
	// Invert columns
	for i := range matrix {
		for j := 0; j < mLen/2; j++ {
			col := mLen - j - 1
			matrix[i][j], matrix[i][col] = matrix[i][col], matrix[i][j]
		}
	}
}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}
