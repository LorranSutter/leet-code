package main

import "fmt"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func insertTree(root *TreeNode, elem int) {
// 	if root.Left == nil && root.Right == nil {
// 		root.Left = &TreeNode{Val: elem}
// 		return
// 	} else if root.Left != nil && root.Right == nil {
// 		root.Right = &TreeNode{Val: elem}
// 		return
// 	} else {
// 		insertTree(root.Left, elem)
// 		// insertTree(root.Right, elem)
// 	}
// }

// func makeTree(input [][]int) *TreeNode {
// 	if len(input) == 0 {
// 		return nil
// 	}

// 	tree := &TreeNode{}

// 	for i := range input {
// 		for _, item := range input[i] {
// 			insertTree(root, item)
// 			printTree(root)
// 			fmt.Println()
// 		}
// 	}

// 	return root
// }

// func printTree(root *TreeNode) {
// 	if root != nil {
// 		fmt.Println(root.Val)
// 		printTree(root.Left)
// 		printTree(root.Right)
// 	}
// }

func search(triangle [][]int, row int, rowIndex int, sum int) int {
	if row >= len(triangle) || rowIndex >= len(triangle[row]) {
		return sum
	}
	return sum + triangle[row][rowIndex] + search(triangle, row+1, rowIndex, sum)
}

func minimumTotal(triangle [][]int) int {
	for row := range triangle {
		for item := range row {

		}
	}
	// printTree(root)
	return 0
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}) == 11)
	// fmt.Println(minimumTotal([][]int{{-10}}) == -10)
}
