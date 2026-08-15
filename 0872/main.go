package main

import (
	"leetcode/utils"
	"math"
)

func leafSimilar(root1 *utils.TreeNode, root2 *utils.TreeNode) bool {
	leafs1, leafs2 := []int{}, []int{}
	getLeafNodes(root1, &leafs1)
	getLeafNodes(root2, &leafs2)

	if len(leafs1) != len(leafs2) {
		return false
	}

	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}

	return true
}

func getLeafNodes(node *utils.TreeNode, leafs *[]int) {
	if node.Left == nil && node.Right == nil {
		*leafs = append(*leafs, node.Val)
		return
	}
	if node.Left != nil {
		getLeafNodes(node.Left, leafs)
	}
	if node.Right != nil {
		getLeafNodes(node.Right, leafs)
	}
}

func main() {
	root1 := utils.MakeBinaryTreeFromLevelOrder([]int{3, 5, 1, 6, 2, 9, 8, int(math.Inf(1)), int(math.Inf(1)), 7, 4}, int(math.Inf(1)))
	root2 := utils.MakeBinaryTreeFromLevelOrder([]int{3, 5, 1, 6, 7, 4, 2, int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), 9, 8}, int(math.Inf(1)))

	root3 := utils.MakeBinaryTreeFromLevelOrder([]int{1, 2, 3}, int(math.Inf(1)))
	root4 := utils.MakeBinaryTreeFromLevelOrder([]int{1, 3, 2}, int(math.Inf(1)))

	utils.RunTests([]utils.TestCase[bool]{
		{Input: []int{3, 5, 1, 6, 2, 9, 8, 7, 4}, Got: leafSimilar(root1, root2), Expected: true},
		{Input: []int{1, 2, 3}, Got: leafSimilar(root3, root4), Expected: false},
	})
}
