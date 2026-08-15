package main

import (
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
	head1 := utils.MakeList([]int{1, 2, 3, 4, 5})
	reversed1 := utils.MakeList([]int{5, 4, 3, 2, 1})

	head2 := utils.MakeList([]int{1, 2})
	reversed2 := utils.MakeList([]int{2, 1})

	head3 := utils.MakeList([]int{})
	reversed3 := utils.MakeList([]int{})

	utils.RunTests([]utils.TestCase[bool]{
		{Input: head1, Got: utils.EqualListNodes(reverseList(head1), reversed1), Expected: true},
		{Input: head2, Got: utils.EqualListNodes(reverseList(head2), reversed2), Expected: true},
		{Input: head3, Got: utils.EqualListNodes(reverseList(head3), reversed3), Expected: true},
	})
}
