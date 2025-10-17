package main

import "fmt"

type Node struct {
	Val        int
	CurrentMin int
	Next       *Node
}

type MinStack struct {
	Head *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	newNode := &Node{Val: val}
	if ms.Head == nil {
		newNode.CurrentMin = val
		ms.Head = newNode
	} else {
		if newNode.Val < ms.Head.CurrentMin {
			newNode.CurrentMin = newNode.Val
		} else {
			newNode.CurrentMin = ms.Head.CurrentMin
		}

		newNode.Next = ms.Head
		ms.Head = newNode
	}
}

func (ms *MinStack) Pop() {
	if ms.Head == nil {
		return
	}

	ms.Head = ms.Head.Next
}

func (ms *MinStack) Top() int {
	return ms.Head.Val
}

func (ms *MinStack) GetMin() int {
	return ms.Head.CurrentMin
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	fmt.Println(obj.Top() == -2)
	obj.Push(0)
	fmt.Println(obj.Top() == 0)
	obj.Push(-3)
	fmt.Println(obj.Top() == -3)
	fmt.Println(obj.GetMin() == -3)
	obj.Pop()
	fmt.Println(obj.Top() == 0)
	fmt.Println(obj.GetMin() == -2)
}
