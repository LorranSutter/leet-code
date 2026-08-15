package main

import (
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
	utils.RunTests([]utils.TestCase[bool]{
		{
			Input:    []int{1, 3, 4, 7, 1, 2, 6},
			Got:      utils.EqualListNodes(deleteMiddle(utils.MakeList([]int{1, 3, 4, 7, 1, 2, 6})), utils.MakeList([]int{1, 3, 4, 1, 2, 6})),
			Expected: true,
		},
		{
			Input:    []int{1, 2, 3, 4},
			Got:      utils.EqualListNodes(deleteMiddle(utils.MakeList([]int{1, 2, 3, 4})), utils.MakeList([]int{1, 2, 4})),
			Expected: true,
		},
		{
			Input:    []int{2, 1},
			Got:      utils.EqualListNodes(deleteMiddle(utils.MakeList([]int{2, 1})), utils.MakeList([]int{2})),
			Expected: true,
		},
		{
			Input:    []int{2},
			Got:      utils.EqualListNodes(deleteMiddle(utils.MakeList([]int{2})), utils.MakeList([]int{})),
			Expected: true,
		},
	})
}
