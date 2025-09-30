package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(nums []int) *ListNode {
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
