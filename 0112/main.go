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
func hasPathSum(root *utils.TreeNode, targetSum int) bool {
	return hasPathSumHelper(root, targetSum, 0)
}

func hasPathSumHelper(node *utils.TreeNode, targetSum, currentSum int) bool {
	if node == nil {
		return false
	}
	currentSum += node.Val
	if node.Left == nil && node.Right == nil {
		return currentSum == targetSum
	}
	return hasPathSumHelper(node.Left, targetSum, currentSum) || hasPathSumHelper(node.Right, targetSum, currentSum)
}

func main() {
	root := utils.MakeBinaryTreeFromLevelOrder([]int{5, 4, 8, 11, int(math.Inf(1)), 13, 4, 7, 2, int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), 1}, int(math.Inf(1)))
	fmt.Println(hasPathSum(root, 22) == true)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1, 2, 3}, int(math.Inf(1)))
	fmt.Println(hasPathSum(root, 5) == false)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{}, int(math.Inf(1)))
	fmt.Println(hasPathSum(root, 0) == false)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1, 2}, int(math.Inf(1)))
	fmt.Println(hasPathSum(root, 1) == false)
}
