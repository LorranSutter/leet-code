package main

import (
	"leetcode/utils"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightIsValidBST(num int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val <= num {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
	}

	return leftIsValidBST(num, root.Left) && rightIsValidBST(num, root.Right)
}

func leftIsValidBST(num int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val >= num {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
	}

	return leftIsValidBST(num, root.Left) && rightIsValidBST(num, root.Right)
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	return leftIsValidBST(root.Val, root.Left) && rightIsValidBST(root.Val, root.Right)
}

func main() {
	// TODO Implement solution
	//  2
	// / \
	//1   3
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	c1 := isValidBST(root)

	//  5
	// / \
	//1   4
	//   / \
	//  3   6
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	c2 := isValidBST(root)

	//  5
	// / \
	//1   6
	//   / \
	//  5   7
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	c3 := isValidBST(root)

	//  5
	// /
	//4
	// \
	//  5
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	c4 := isValidBST(root)

	//  5
	// /
	//4
	// \
	//  1
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 1}
	c5 := isValidBST(root)

	//  2
	// / \
	//2   2
	root = &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	c6 := isValidBST(root)

	//  5
	// / \
	//4   6
	//   / \
	//  3   7
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 7}
	c7 := isValidBST(root)

	utils.RunTests([]utils.TestCase[bool]{
		{Input: []int{2, 1, 3}, Got: c1, Expected: true},
		{Input: []int{5, 1, 4, 3, 6}, Got: c2, Expected: false},
		{Input: []int{5, 1, 6, 5, 7}, Got: c3, Expected: true},
		{Input: []int{5, 4, 5}, Got: c4, Expected: true},
		{Input: []int{5, 4, 1}, Got: c5, Expected: false},
		{Input: []int{2, 2, 2}, Got: c6, Expected: false},
		{Input: []int{5, 4, 6, 3, 7}, Got: c7, Expected: false},
	})
}
