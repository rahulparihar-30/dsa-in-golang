package Linkedlist

import "errors"

type deQueNode struct {
	prev *deQueNode
	data int
	next *deQueNode
}

var start *deQueNode = nil
var tail *deQueNode = nil

func NilInsertion(node *deQueNode) {
	node.prev = nil
	node.next = nil
	start = node
	tail = node
}

func PushFront(value int) {
	newNode := &deQueNode{data: value}
	if start == nil {
		NilInsertion(newNode)
		return
	}
	newNode.prev = nil
	newNode.next = start
	start = newNode
}
func PushBack(value int) {
	newNode := &deQueNode{data: value}
	if start == nil {
		NilInsertion(newNode)
		return
	}
	newNode.prev = tail
	tail.next = newNode
	newNode.next = nil
	tail = newNode
}
func PushAt(value, location int) (bool, error) {

	if start == nil {
		return false, errors.New("head is nil")
	}

	newNode := &deQueNode{data: value}

	if location == 0 {
		newNode.next = start
		newNode.prev = nil
		start = newNode
		return true, nil
	}

	curr := start
	index := 0

	for curr != nil && index < location-1 {
		curr = curr.next
		index++
	}

	if curr == nil {
		return false, errors.New("index out of bound")
	}
	newNode.next = curr.next
	newNode.prev = curr
	curr.next = newNode
	return true, nil
}
