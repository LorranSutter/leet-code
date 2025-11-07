package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
			fmt.Println("Different values", s1[i], s2[i])
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
