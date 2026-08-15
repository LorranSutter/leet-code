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
func searchBST(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}

func main() {
	root := utils.MakeBinaryTree([]int{4, 2, 7, 1, 3})
	result := searchBST(root, 2)

	root2 := utils.MakeBinaryTree([]int{4, 2, 7, 1, 3})
	result2 := searchBST(root2, 5)

	utils.RunTests([]utils.TestCase[bool]{
		{Input: 2, Got: result.Val == 2, Expected: true},
		{Input: 5, Got: result2 == nil, Expected: true},
	})
}
