package main

import "fmt"

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type ElementsList struct {
	Head *Node
	Tail *Node
}

type LRUCache struct {
	Capacity        int
	CurrentCapacity int
	Elements        ElementsList
	KeysMap         map[int]*Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity,
		0,
		ElementsList{},
		make(map[int]*Node, capacity),
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	node := this.KeysMap[key]

	if node == nil {
		return -1
	}
	this.MoveElementToHead(node)

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node := this.KeysMap[key]

	if node != nil {
		node.Val = value
		this.MoveElementToHead(node)
		return
	}

	newNode := Node{Key: key, Val: value}

	if this.Elements.Head == nil {
		this.Elements.Head = &newNode
		this.Elements.Tail = &newNode
	} else {
		newNode.Next = this.Elements.Head
		this.Elements.Head.Prev = &newNode
		this.Elements.Head = &newNode
	}

	if this.CurrentCapacity < this.Capacity {
		this.CurrentCapacity++
	} else {
		delete(this.KeysMap, this.Elements.Tail.Key)
		this.Elements.Tail = this.Elements.Tail.Prev
		this.Elements.Tail.Next = nil
	}

	this.KeysMap[key] = &newNode
}

func (this *LRUCache) MoveElementToHead(node *Node) {
	if node != this.Elements.Head {
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
		if node == this.Elements.Tail {
			this.Elements.Tail = node.Prev
		}

		node.Next = this.Elements.Head
		this.Elements.Head.Prev = node

		if this.Elements.Head != nil {
			this.Elements.Head.Prev = node
		}
		this.Elements.Head = node

		if this.Elements.Tail == nil {
			this.Elements.Tail = node
		}
	}
}

func (this *LRUCache) PrintList() {
	currentNode := this.Elements.Head
	for currentNode != nil {
		fmt.Print(currentNode.Key, ",", currentNode.Val, " ")
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1) == 1)
	obj.Put(3, 3)
	fmt.Println(obj.Get(2) == -1)
	obj.Put(4, 4)
	fmt.Println(obj.Get(1) == -1)
	fmt.Println(obj.Get(3) == 3)
	fmt.Println(obj.Get(4) == 4)

	fmt.Println()

	obj = Constructor(2)
	fmt.Println(obj.Get(2) == -1)
	obj.Put(2, 6)
	fmt.Println(obj.Get(1) == -1)
	obj.Put(1, 5)
	obj.Put(1, 2)
	fmt.Println(obj.Get(1) == 2)
	fmt.Println(obj.Get(2) == 6)

	fmt.Println()

	obj = Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	fmt.Println(obj.Get(1) == -1)
	fmt.Println(obj.Get(2) == 3)
}
