package main

import (
	"leetcode/utils"
	"math"
)

// The idea here is to perform a level order traversal (BFS) of the tree,
// calculating the sum of node values at each level.
// We keep track of the maximum sum encountered and the corresponding level.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *utils.TreeNode) int {
	queue := []*utils.TreeNode{root}
	minLevel, maxSum := 1, int(math.Inf(-1))

	level, levelSum := minLevel, 0
	var node *utils.TreeNode
	for len(queue) > 0 {
		levelSize := len(queue)

		for range levelSize {
			node = queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if levelSum > maxSum {
			maxSum = levelSum
			minLevel = level
		}

		levelSum = 0
		level++
	}

	return minLevel
}

func main() {
	root1 := utils.MakeBinaryTreeFromLevelOrder([]int{1, 7, 0, 7, -8, int(math.Inf(1)), int(math.Inf(1))}, int(math.Inf(1)))
	root2 := utils.MakeBinaryTreeFromLevelOrder([]int{989, int(math.Inf(1)), 10250, 98693, -89388, int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), -32127}, int(math.Inf(1)))
	root3 := utils.MakeBinaryTreeFromLevelOrder([]int{-100, -200, -300, -20, -5, -10, int(math.Inf(1))}, int(math.Inf(1)))

	utils.RunTests([]utils.TestCase[int]{
		{Input: []int{1, 7, 0, 7, -8}, Got: maxLevelSum(root1), Expected: 2},
		{Input: []int{989, 10250, 98693, -89388, -32127}, Got: maxLevelSum(root2), Expected: 2},
		{Input: []int{-100, -200, -300, -20, -5, -10}, Got: maxLevelSum(root3), Expected: 3},
	})
}
