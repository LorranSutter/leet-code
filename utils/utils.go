package utils

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[len(nums)-1], Next: nil}

	for i := len(nums) - 2; i >= 0; i-- {
		newNode := &ListNode{Val: nums[i], Next: head}
		head = newNode
	}

	return head
}

func MakeBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}

	for i := 1; i < len(nums); i++ {
		addBinaryTreeNode(root, nums[i])
	}

	return root
}

// Add TreeNode to the binary tree
func addBinaryTreeNode(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			addBinaryTreeNode(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			addBinaryTreeNode(root.Right, val)
		}
	}
}

func PrintList(list *ListNode) {
	fmt.Println(list)
	currentNode := list
	for currentNode != nil {
		fmt.Print(currentNode.Val, " ")
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func PrintMatrix[T any](m [][]T) {
	for i := range m {
		fmt.Println(m[i])
	}
}

func EqualListNodes(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil || l2 != nil {
		return false
	}

	return true
}

func DeepEqualSlices[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		fmt.Println("Different lengths", len(s1), len(s2))
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			fmt.Printf("Different values at (%d): [s1 -> %v] [s2 -> %v]\n", i, s1[i], s2[i])
			return false
		}
	}

	return true
}

func EqualSlices[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		fmt.Println("Different lengths", len(s1), len(s2))
		return false
	}

	counts := make(map[T]int)
	for _, v := range s1 {
		counts[v]++
	}

	for _, v := range s2 {
		counts[v]--
		if counts[v] < 0 {
			fmt.Println("Element not foud or too many occurences", v)
			return false
		}
	}

	// All counts should be zero if slices are unordered equal
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func DeepEqualMatrix[T comparable](m1, m2 [][]T) bool {
	if len(m1) != len(m2) {
		fmt.Println("Different lengths", len(m1), len(m2))
		return false
	}

	for i := range m1 {
		for j := range m1[i] {
			if m1[i][j] != m2[i][j] {
				fmt.Printf("Different values at (%d,%d): [m1 -> %v] [m2 -> %v]\n", i, j, m1[i][j], m2[i][j])
				return false
			}
		}
	}

	return true
}
