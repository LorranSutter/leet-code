package main

import "fmt"

type Data struct {
	UserId   int
	TaskId   int
	Priority int
}

type Node struct {
	Data Data
	Prev *Node
	Next *Node
}

type PriorityList struct {
	Head *Node
}

type TaskManager struct {
	// tasksToPriority map[int]int // Does not make sense bc when insert leads to lots of updates
	Tasks        map[int]*Node
	PriorityList PriorityList
}

func Constructor(tasks [][]int) TaskManager {

	tm := TaskManager{
		Tasks:        make(map[int]*Node),
		PriorityList: PriorityList{Head: nil},
	}

	for _, task := range tasks {
		tm.Add(task[0], task[1], task[2])
	}

	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	data := Data{userId, taskId, priority}
	newNode := &Node{Data: data, Prev: nil, Next: nil}

	fmt.Println(newNode)

	// Empty list
	if this.PriorityList.Head == nil {
		this.PriorityList.Head = newNode
		this.Tasks[taskId] = newNode
		return
	}

	// Check if new node should be the new head (highest priority)
	if this.PriorityList.Head.Data.Priority < priority ||
		(this.PriorityList.Head.Data.Priority == priority && this.PriorityList.Head.Data.TaskId < taskId) {
		this.PriorityList.Head.Prev = newNode
		newNode.Next = this.PriorityList.Head
		this.PriorityList.Head = newNode

		this.Tasks[taskId] = newNode
		return
	}

	// Find the correct position to insert
	current := this.PriorityList.Head

	for current != nil {
		// If we found the right spot (current has lower priority or same priority but higher taskId)
		if current.Data.Priority < priority ||
			(current.Data.Priority == priority && current.Data.TaskId < taskId) {
			break
		}
		current = current.Next
	}

	// Insert the new node
	if current != nil {
		// Inserting before current node
		newNode.Prev = current.Prev
		newNode.Next = current
		if current.Prev != nil {
			current.Prev.Next = newNode
		}
		current.Prev = newNode
	} else {
		// Inserting at the end (current is nil)
		// Find the last node
		last := this.PriorityList.Head
		for last.Next != nil {
			last = last.Next
		}
		last.Next = newNode
		newNode.Prev = last
	}

	this.Tasks[taskId] = newNode

}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	current := this.PriorityList.Head

	for current != nil {
		if current.Data.TaskId == taskId {
			this.Rmv(taskId)
			this.Add(current.Data.UserId, current.Data.TaskId, newPriority)
			return
		}
		current = current.Next
	}
}

func (this *TaskManager) Rmv(taskId int) {
	node, exists := this.Tasks[taskId]
	if !exists {
		return
	}

	// Update previous node's next pointer
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		// This was the head node
		this.PriorityList.Head = node.Next
	}

	// Update next node's previous pointer
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	// Remove from map
	delete(this.Tasks, taskId)
}

func (this *TaskManager) ExecTop() int {
	if this.PriorityList.Head == nil {
		return -1
	}

	userId := this.PriorityList.Head.Data.UserId
	taskId := this.PriorityList.Head.Data.TaskId

	// Update head
	this.PriorityList.Head = this.PriorityList.Head.Next
	if this.PriorityList.Head != nil {
		this.PriorityList.Head.Prev = nil
	}

	// Remove from map
	delete(this.Tasks, taskId)

	return userId
}

func (this *TaskManager) Print() {
	currentNode := this.PriorityList.Head
	for {
		if currentNode == nil {
			break
		}
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */

func main() {
	tasks := [][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}}
	taskManager := Constructor(tasks)

	taskManager.Print()
	fmt.Println()

	taskManager.Add(4, 104, 15)

	taskManager.Print()
	fmt.Println()

	taskManager.Add(5, 105, 15)

	taskManager.Print()
	fmt.Println()

	taskManager.Add(6, 106, 20)

	taskManager.Print()
	fmt.Println()

	taskManager.Add(7, 107, 23)

	taskManager.Print()
	fmt.Println()

	taskManager.Add(8, 108, 10)

	taskManager.Print()
	fmt.Println()

	fmt.Println(taskManager.ExecTop())

	taskManager.Rmv(101)

	taskManager.Print()

	fmt.Println()

	taskManager.Edit(105, 20)

	taskManager.Print()
}
