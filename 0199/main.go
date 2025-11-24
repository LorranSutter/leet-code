package main

import (
	"fmt"
	"leetcode/utils"
	"math"
)

// The idea here is to perform a level order traversal (BFS) of the tree
// but always enqueue the right child before the left child.
// This way, the first node we encounter at each level is the rightmost node.

func rightSideView(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	rightmostValues := []int{}
	queue := []*utils.TreeNode{root}
	queue = append(queue, root)

	levelSize := 0
	var node *utils.TreeNode
	for len(queue) > 0 {
		levelSize = len(queue)
		rightmostValues = append(rightmostValues, queue[0].Val)

		for i := 0; i < levelSize; i++ {
			node = queue[0]
			queue = queue[1:]

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}

	return rightmostValues
}

func main() {
	root := utils.MakeBinaryTreeFromLevelOrder([]int{1, 2, 3, int(math.Inf(1)), 5, int(math.Inf(1)), 4}, int(math.Inf(1)))
	result := rightSideView(root)
	expected := []int{1, 3, 4}
	fmt.Println(utils.DeepEqualSlices(result, expected))

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1, 2, 3, 4, int(math.Inf(1)), int(math.Inf(1)), int(math.Inf(1)), 5}, int(math.Inf(1)))
	result = rightSideView(root)
	expected = []int{1, 3, 4, 5}
	fmt.Println(utils.DeepEqualSlices(result, expected))

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1, int(math.Inf(1)), 3}, int(math.Inf(1)))
	result = rightSideView(root)
	expected = []int{1, 3}
	fmt.Println(utils.DeepEqualSlices(result, expected))

	root = utils.MakeBinaryTreeFromLevelOrder([]int{}, int(math.Inf(1)))
	result = rightSideView(root)
	expected = []int{}
	fmt.Println(utils.DeepEqualSlices(result, expected))

	root = utils.MakeBinaryTreeFromLevelOrder([]int{1, 2, 3, int(math.Inf(1)), 5, 6, int(math.Inf(1)), 4}, int(math.Inf(1)))
	result = rightSideView(root)
	expected = []int{1, 3, 6, 4}
	fmt.Println(utils.DeepEqualSlices(result, expected))
}
