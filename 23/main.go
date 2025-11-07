package main

import (
	"leetcode/utils"
)

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
	if len(lists) == 0 {
		return &utils.ListNode{}
	}
	if len(lists) == 1 {
		return lists[0]
	}

	finishedLists := make([]bool, len(lists))
	numFinishedLists := 0
	node := &utils.ListNode{}
	head := &utils.ListNode{}

	minValue := 100000
	for i, list := range lists {
		if list == (&utils.ListNode{}) {
			finishedLists[i] = true
			numFinishedLists++
		}
		if list.Val < head.Val {
			head = list
			minValue = list.Val
		}
	}
	node = head

	for numFinishedLists < len(lists) {
		for i := range lists {
			if finishedLists[i] || lists[i] == nil {
				continue
			}
			for lists[i] != nil && lists[i].Val == minValue {
				// node.Next = &utils.ListNode{Val: lists[i].Val, Next: nil}
				node.Next = lists[i]
				node = node.Next
				lists[i] = lists[i].Next
			}
			if lists[i] == nil {
				finishedLists[i] = true
				numFinishedLists++
			}
		}
		minValue = node.Val + 1
	}

	return head.Next
}

func main() {
	list1 := utils.MakeList([]int{1, 1, 4, 5})
	list2 := utils.MakeList([]int{1, 3, 4})
	list3 := utils.MakeList([]int{2, 6})
	print(mergeKLists([]*utils.ListNode{list1}))
	print(mergeKLists([]*utils.ListNode{list1, list2, list3}))
	// print(mergeKLists([]*utils.ListNode{list1, list2, list3, {}})) // Fix this case
	print(mergeKLists([]*utils.ListNode{}))   // Fix this case
	print(mergeKLists([]*utils.ListNode{{}})) // Fix this case ?
}
