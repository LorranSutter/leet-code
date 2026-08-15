package main

import (
	"leetcode/utils"
)

func modifiedList(nums []int, head *utils.ListNode) *utils.ListNode {
	numberSet := make(map[int]struct{})

	for _, num := range nums {
		numberSet[num] = struct{}{}
	}

	node, prev := head, &utils.ListNode{}
	for node != nil {
		if _, ok := numberSet[node.Val]; ok {
			if node == head {
				head = head.Next
			} else {
				prev.Next = node.Next
			}
		} else {
			prev = node
		}
		node = node.Next
	}

	return head
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{
			Input:    []any{[]int{1, 2, 3}, []int{1, 2, 3, 4, 5}},
			Got:      utils.EqualListNodes(modifiedList([]int{1, 2, 3}, utils.MakeList([]int{1, 2, 3, 4, 5})), utils.MakeList([]int{4, 5})),
			Expected: true,
		},
		{
			Input:    []any{[]int{1}, []int{1, 2, 1, 2, 1, 2}},
			Got:      utils.EqualListNodes(modifiedList([]int{1}, utils.MakeList([]int{1, 2, 1, 2, 1, 2})), utils.MakeList([]int{2, 2, 2})),
			Expected: true,
		},
		{
			Input:    []any{[]int{1}, []int{1, 1, 1, 1, 2, 1, 2, 1, 2}},
			Got:      utils.EqualListNodes(modifiedList([]int{1}, utils.MakeList([]int{1, 1, 1, 1, 2, 1, 2, 1, 2})), utils.MakeList([]int{2, 2, 2})),
			Expected: true,
		},
		{
			Input:    []any{[]int{1}, []int{1, 2, 1, 2, 1, 2, 1, 1, 1}},
			Got:      utils.EqualListNodes(modifiedList([]int{1}, utils.MakeList([]int{1, 2, 1, 2, 1, 2, 1, 1, 1})), utils.MakeList([]int{2, 2, 2})),
			Expected: true,
		},
		{
			Input:    []any{[]int{5}, []int{1, 2, 3, 4}},
			Got:      utils.EqualListNodes(modifiedList([]int{5}, utils.MakeList([]int{1, 2, 3, 4})), utils.MakeList([]int{1, 2, 3, 4})),
			Expected: true,
		},
		{
			Input:    []any{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			Got:      utils.EqualListNodes(modifiedList([]int{1, 2, 3, 4}, utils.MakeList([]int{1, 2, 3, 4})), utils.MakeList(nil)),
			Expected: true,
		},
	})
}
