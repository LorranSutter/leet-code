package main

import (
	"fmt"
	"leetcode/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
	root := utils.MakeBinaryTreeFromLevelOrder([]int{3, 9, 20, -1, -1, 15, 7}, -1)
	fmt.Println(maxDepth(root) == 3)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1, -1, 2}, -1)
	fmt.Println(maxDepth(root) == 2)
}
