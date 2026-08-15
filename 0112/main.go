package main

import (
	"math"

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
	utils.RunTests([]utils.TestCase[bool]{
		{
			Input:    []int{5, 4, 8, 11, 13, 4, 7, 2, 1},
			Got:      hasPathSum(utils.MakeBinaryTreeFromLevelOrder([]int{5, 4, 8, 11, int(math.Inf(1)), 13, 4, 7, 2, int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), 1}, int(math.Inf(1))), 22),
			Expected: true,
		},
		{
			Input:    []int{1, 2, 3},
			Got:      hasPathSum(utils.MakeBinaryTreeFromLevelOrder([]int{1, 2, 3}, int(math.Inf(1))), 5),
			Expected: false,
		},
		{
			Input:    []int{},
			Got:      hasPathSum(utils.MakeBinaryTreeFromLevelOrder([]int{}, int(math.Inf(1))), 0),
			Expected: false,
		},
		{
			Input:    []int{1, 2},
			Got:      hasPathSum(utils.MakeBinaryTreeFromLevelOrder([]int{1, 2}, int(math.Inf(1))), 1),
			Expected: false,
		},
	})
}
