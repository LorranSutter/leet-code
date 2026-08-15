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
func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	node := &utils.ListNode{}
	head := &utils.ListNode{}
	head = node

	partialSum := 0
	carry := 0
	for l1 != nil && l2 != nil {
		partialSum = l1.Val + l2.Val + carry
		if partialSum > 9 {
			partialSum %= 10
			carry = 1
		} else {
			carry = 0
		}
		node.Next = &utils.ListNode{Val: partialSum}
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	var missing *utils.ListNode
	if l1 != nil {
		missing = l1
	} else if l2 != nil {
		missing = l2
	}

	for missing != nil {
		partialSum = missing.Val + carry
		if partialSum > 9 {
			partialSum %= 10
			carry = 1
		} else {
			carry = 0
		}
		node.Next = &utils.ListNode{Val: partialSum}
		node = node.Next
		missing = missing.Next
	}

	if carry != 0 {
		node.Next = &utils.ListNode{Val: carry}
	}

	return head.Next
}

func main() {
	l1 := utils.MakeList([]int{2, 4, 3})
	l2 := utils.MakeList([]int{5, 6, 4})
	l3 := utils.MakeList([]int{7, 0, 8})
	c1 := utils.EqualListNodes(addTwoNumbers(l1, l2), l3)

	l1 = utils.MakeList([]int{0})
	l2 = utils.MakeList([]int{0})
	l3 = utils.MakeList([]int{0})
	c2 := utils.EqualListNodes(addTwoNumbers(l1, l2), l3)

	l1 = utils.MakeList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = utils.MakeList([]int{9, 9, 9, 9})
	l3 = utils.MakeList([]int{8, 9, 9, 9, 0, 0, 0, 1})
	c3 := utils.EqualListNodes(addTwoNumbers(l1, l2), l3)

	l1 = utils.MakeList([]int{2, 4, 9})
	l2 = utils.MakeList([]int{5, 6, 4, 9})
	l3 = utils.MakeList([]int{7, 0, 4, 0, 1})
	c4 := utils.EqualListNodes(addTwoNumbers(l1, l2), l3)

	utils.RunTests([]utils.TestCase[bool]{
		{Input: [][]int{{2, 4, 3}, {5, 6, 4}}, Got: c1, Expected: true},
		{Input: [][]int{{0}, {0}}, Got: c2, Expected: true},
		{Input: [][]int{{9, 9, 9, 9, 9, 9, 9}, {9, 9, 9, 9}}, Got: c3, Expected: true},
		{Input: [][]int{{2, 4, 9}, {5, 6, 4, 9}}, Got: c4, Expected: true},
	})
}
