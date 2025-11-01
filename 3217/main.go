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
	head := utils.MakeList([]int{1, 2, 3, 4, 5})
	newHead := modifiedList([]int{1, 2, 3}, head)
	utils.PrintList(newHead)

	head = utils.MakeList([]int{1, 2, 1, 2, 1, 2})
	newHead = modifiedList([]int{1}, head)
	utils.PrintList(newHead)

	head = utils.MakeList([]int{1, 1, 1, 1, 2, 1, 2, 1, 2})
	newHead = modifiedList([]int{1}, head)
	utils.PrintList(newHead)

	head = utils.MakeList([]int{1, 2, 1, 2, 1, 2, 1, 1, 1})
	newHead = modifiedList([]int{1}, head)
	utils.PrintList(newHead)

	head = utils.MakeList([]int{1, 2, 3, 4})
	newHead = modifiedList([]int{5}, head)
	utils.PrintList(newHead)

	head = utils.MakeList([]int{1, 2, 3, 4})
	newHead = modifiedList([]int{1, 2, 3, 4}, head)
	utils.PrintList(newHead)
}
