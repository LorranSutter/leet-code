package main

import (
	"fmt"
	"leetcode/utils"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *utils.TreeNode, targetSum int) [][]int {
	paths := [][]int{}
	currentPath := make([]int, 0)
	pathSumHelper(root, targetSum, 0, currentPath, &paths)

	return paths
}

func pathSumHelper(node *utils.TreeNode, targetSum, currentSum int, currentPath []int, paths *[][]int) {
	if node == nil {
		return
	}
	currentSum += node.Val
	currentPath = append(currentPath, node.Val)
	if node.Left == nil && node.Right == nil && currentSum == targetSum {
		// Create a copy of currentPath before appending
		pathCopy := make([]int, len(currentPath))
		copy(pathCopy, currentPath)
		(*paths) = append(*paths, pathCopy)
	}
	pathSumHelper(node.Left, targetSum, currentSum, currentPath, paths)
	pathSumHelper(node.Right, targetSum, currentSum, currentPath, paths)
}

func main() {
	root := utils.MakeBinaryTreeFromLevelOrder([]int{5, 4, 8, 11, int(math.Inf(1)), 13, 4, 7, 2, int(math.Inf(1)), int(math.Inf(1)), 5, 1}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 22))

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1, 2, 3}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 5))

	root = utils.MakeBinaryTreeFromLevelOrder([]int{}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 0))

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1, 2}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 1))

	root = utils.MakeBinaryTreeFromLevelOrder([]int{-2, int(math.Inf(1)), -3}, int(math.Inf(1)))
	fmt.Println(pathSum(root, -5))
}
