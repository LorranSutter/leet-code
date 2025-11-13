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
	root := utils.MakeTree([]int{4, 2, 7, 1, 3})
	result := searchBST(root, 2)
	fmt.Println(result.Val == 2)

	root = utils.MakeTree([]int{4, 2, 7, 1, 3})
	result = searchBST(root, 5)
	fmt.Println(result == nil)
}
