package main

import (
	"fmt"
	"leetcode/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}

	var prev *utils.ListNode
	temp := &utils.ListNode{}
	node := head
	for node != nil {
		temp = node.Next
		node.Next = prev
		prev = node
		node = temp
	}
	return prev
}

func main() {
	head := utils.MakeList([]int{1, 2, 3, 4, 5})
	reversed := utils.MakeList([]int{5, 4, 3, 2, 1})
	result := reverseList(head)
	fmt.Println(utils.EqualListNodes(result, reversed))

	head = utils.MakeList([]int{1, 2})
	reversed = utils.MakeList([]int{2, 1})
	result = reverseList(head)
	fmt.Println(utils.EqualListNodes(result, reversed))

	head = utils.MakeList([]int{})
	reversed = utils.MakeList([]int{})
	result = reverseList(head)
	fmt.Println(utils.EqualListNodes(result, reversed))
}
