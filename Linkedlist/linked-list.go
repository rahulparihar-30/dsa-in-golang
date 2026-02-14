package Linkedlist

type Node struct {
	Data int
	Next *Node
}

func Insert(head *Node, data int) *Node {
	if head == nil {
		head = &Node{Data: data}
		return head
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Data: data}
	return head
}
