package main

import (
	"fmt"
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
	root := utils.MakeBinaryTreeFromLevelOrder([]int{10, 5, -3, 3, 2, int(math.Inf(1)), 11, 3, -2, int(math.Inf(1)), 1}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 8) == 3)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{5, 4, 8, 11, int(math.Inf(1)), 13, 4, 7, 2, int(math.Inf(1)), int(math.Inf(1)), 5, 1}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 22) == 3)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{0, 1, 1}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 1) == 4)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{10, 5, -3, 3, 2, int(math.Inf(1)), 11, 0, 1, int(math.Inf(1)), 1, int(math.Inf(1)), 0, 0, int(math.Inf(1)), 0, 0, 0, int(math.Inf(1)), 1, int(math.Inf(1)), 0, int(math.Inf(1)), 0}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 8) == 8)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1}, int(math.Inf(1)))
	fmt.Println(pathSum(root, 0) == 0)

	root = utils.MakeBinaryTreeFromLevelOrder([]int{-2, int(math.Inf(1)), -3}, int(math.Inf(1)))
	fmt.Println(pathSum(root, -3) == 1)
}
