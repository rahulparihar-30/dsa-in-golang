package Linkedlist

type node struct {
	data int
	next *node
}

var top *node = nil

func Push(data int) {
	if top == nil {
		top = &node{data: data}
		top.next = top
	}
	newNode := &node{data: data}
	newNode.next = top
	top = newNode
}

func Pop() int {
	if top == nil {
		return -1
	}
	data := top.data
	top = top.next
	return data
}

func Peek() int {
	if top == nil {
		return -1
	}
	return top.data
}
