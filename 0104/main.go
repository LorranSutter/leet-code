package main

import (
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
	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{3, 9, 20, -1, -1, 15, 7}, Got: maxDepth(utils.MakeBinaryTreeFromLevelOrder([]int{3, 9, 20, -1, -1, 15, 7}, -1)), Expected: 3},
		{Input: []int{1, -1, 2}, Got: maxDepth(utils.MakeBinaryTreeFromLevelOrder([]int{1, -1, 2}, -1)), Expected: 2},
	})
}
