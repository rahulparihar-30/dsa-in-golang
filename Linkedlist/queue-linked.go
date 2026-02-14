package Linkedlist

import "errors"

type QueueNode struct {
	value int
	next  *QueueNode
}

var FRONT *QueueNode = nil
var REAR *QueueNode = nil

func Enqueue(value int) {
	newNode := &QueueNode{value: value}
	newNode.next = nil
	if FRONT == nil {
		FRONT = newNode
		REAR = newNode
		return
	}
	REAR.next = newNode
	REAR = REAR.next
}

func Dequeue() (int, error) {
	if FRONT == nil {
		return 0, errors.New("queue is empty")
	}
	value := FRONT.value
	FRONT = FRONT.next
	return value, nil
}

func Front() (int, error) {
	if FRONT == nil {
		return 0, errors.New("queue is empty")
	}
	return FRONT.value, nil
}
