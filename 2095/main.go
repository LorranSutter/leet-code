package main

import (
	"fmt"
	"leetcode/utils"
)

func deleteMiddle(head *utils.ListNode) *utils.ListNode {
	if head.Next == nil {
		return nil
	}

	preSlow := &utils.ListNode{Next: head}
	slow := head
	fast := head

	for {
		if fast.Next == nil {
			preSlow.Next = slow.Next
			break
		}
		if fast.Next.Next == nil {
			slow.Next = slow.Next.Next
			break
		}
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	return head
}

func main() {
	head := utils.MakeList([]int{1, 3, 4, 7, 1, 2, 6})
	newHead := deleteMiddle(head)
	expected := utils.MakeList([]int{1, 3, 4, 1, 2, 6})
	fmt.Println(utils.EqualListNodes(newHead, expected))

	head = utils.MakeList([]int{1, 2, 3, 4})
	newHead = deleteMiddle(head)
	expected = utils.MakeList([]int{1, 2, 4})
	fmt.Println(utils.EqualListNodes(newHead, expected))

	head = utils.MakeList([]int{2, 1})
	newHead = deleteMiddle(head)
	expected = utils.MakeList([]int{2})
	fmt.Println(utils.EqualListNodes(newHead, expected))

	head = utils.MakeList([]int{2})
	newHead = deleteMiddle(head)
	expected = utils.MakeList([]int{})
	fmt.Println(utils.EqualListNodes(newHead, expected))
}
