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
	root1a := utils.MakeBinaryTreeFromLevelOrder([]int{3, 1, 4, 3, -1, 1, 5}, -1)
	root2a := utils.MakeBinaryTreeFromLevelOrder([]int{3, 3, -1, 4, 2}, -1)
	root3a := utils.MakeBinaryTreeFromLevelOrder([]int{1}, -1)

	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{3, 1, 4, 3, -1, 1, 5}, Got: goodNodes1(root1a), Expected: 4},
		{Input: []int{3, 3, -1, 4, 2}, Got: goodNodes1(root2a), Expected: 3},
		{Input: []int{1}, Got: goodNodes1(root3a), Expected: 1},
	})

	root1b := utils.MakeBinaryTreeFromLevelOrder([]int{3, 1, 4, 3, -1, 1, 5}, -1)
	root2b := utils.MakeBinaryTreeFromLevelOrder([]int{3, 3, -1, 4, 2}, -1)
	root3b := utils.MakeBinaryTreeFromLevelOrder([]int{1}, -1)

	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{3, 1, 4, 3, -1, 1, 5}, Got: goodNodes2(root1b), Expected: 4},
		{Input: []int{3, 3, -1, 4, 2}, Got: goodNodes2(root2b), Expected: 3},
		{Input: []int{1}, Got: goodNodes2(root3b), Expected: 1},
	})
}
