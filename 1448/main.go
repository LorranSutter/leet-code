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
// 67 ms / 13.02 MB
func goodNodes1(root *utils.TreeNode) int {
	result := goodNodesHelper1(root, root.Val)
	return result
}

func goodNodesHelper1(node *utils.TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	if node.Val >= maxVal {
		maxVal = node.Val
		return 1 + goodNodesHelper1(node.Left, maxVal) + goodNodesHelper1(node.Right, maxVal)
	}
	return goodNodesHelper1(node.Left, maxVal) + goodNodesHelper1(node.Right, maxVal)
}

// 88 ms / 12.73 MB
func goodNodes2(root *utils.TreeNode) int {
	count := 0
	goodNodesHelper2(root, root.Val, &count)
	return count
}

func goodNodesHelper2(node *utils.TreeNode, maxVal int, count *int) {
	if node == nil {
		return
	}
	if node.Val >= maxVal {
		maxVal = node.Val
		(*count)++
	}
	goodNodesHelper2(node.Left, maxVal, count)
	goodNodesHelper2(node.Right, maxVal, count)
}

func main() {
	fmt.Println("Solution 01")
	root := utils.MakeBinaryTreeFromLevelOrder([]int{3, 1, 4, 3, -1, 1, 5}, -1)
	fmt.Println(goodNodes1(root) == 4)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{3, 3, -1, 4, 2}, -1)
	fmt.Println(goodNodes1(root) == 3)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1}, -1)
	fmt.Println(goodNodes1(root) == 1)

	fmt.Println()

	fmt.Println("Solution 02")
	root = utils.MakeBinaryTreeFromLevelOrder([]int{3, 1, 4, 3, -1, 1, 5}, -1)
	fmt.Println(goodNodes2(root) == 4)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{3, 3, -1, 4, 2}, -1)
	fmt.Println(goodNodes2(root) == 3)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1}, -1)
	fmt.Println(goodNodes2(root) == 1)
}
