package main

import (
	"leetcode/utils"
	"math"
)

// The idea here is to use DFS to explore all paths in the tree.
// For each node, we check all paths starting from that node to see if they sum to targetSum.

// Definitely we can optimize this further using prefix sums and a hashmap to store cumulative sums

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *utils.TreeNode, targetSum int) int {
	count := 0
	dfs(root, targetSum, &count)
	return count
}

func dfs(node *utils.TreeNode, targetSum int, count *int) {
	if node == nil {
		return
	}
	currentSum := node.Val
	if currentSum == targetSum {
		(*count)++
	}

	pathSumHelper(node.Left, targetSum, currentSum, count)
	dfs(node.Left, targetSum, count)

	pathSumHelper(node.Right, targetSum, currentSum, count)
	dfs(node.Right, targetSum, count)
}

func pathSumHelper(node *utils.TreeNode, targetSum, currentSum int, count *int) {
	if node == nil {
		return
	}
	currentSum += node.Val
	if currentSum == targetSum {
		(*count)++
	}
	pathSumHelper(node.Left, targetSum, currentSum, count)
	pathSumHelper(node.Right, targetSum, currentSum, count)
}

func main() {
	root1 := utils.MakeBinaryTreeFromLevelOrder([]int{10, 5, -3, 3, 2, int(math.Inf(1)), 11, 3, -2, int(math.Inf(1)), 1}, int(math.Inf(1)))
	root2 := utils.MakeBinaryTreeFromLevelOrder([]int{5, 4, 8, 11, int(math.Inf(1)), 13, 4, 7, 2, int(math.Inf(1)), int(math.Inf(1)), 5, 1}, int(math.Inf(1)))
	root3 := utils.MakeBinaryTreeFromLevelOrder([]int{0, 1, 1}, int(math.Inf(1)))
	root4 := utils.MakeBinaryTreeFromLevelOrder([]int{10, 5, -3, 3, 2, int(math.Inf(1)), 11, 0, 1, int(math.Inf(1)), 1, int(math.Inf(1)), 0, 0, int(math.Inf(1)), 0, 0, 0, int(math.Inf(1)), 1, int(math.Inf(1)), 0, int(math.Inf(1)), 0}, int(math.Inf(1)))
	root5 := utils.MakeBinaryTreeFromLevelOrder([]int{1}, int(math.Inf(1)))
	root6 := utils.MakeBinaryTreeFromLevelOrder([]int{-2, int(math.Inf(1)), -3}, int(math.Inf(1)))

	utils.RunTests([]utils.TestCase[int]{
		{Input: root1, Got: pathSum(root1, 8), Expected: 3},
		{Input: root2, Got: pathSum(root2, 22), Expected: 3},
		{Input: root3, Got: pathSum(root3, 1), Expected: 4},
		{Input: root4, Got: pathSum(root4, 8), Expected: 8},
		{Input: root5, Got: pathSum(root5, 0), Expected: 0},
		{Input: root6, Got: pathSum(root6, -3), Expected: 1},
	})
}
