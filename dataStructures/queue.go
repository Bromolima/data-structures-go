package dataStructures

import "fmt"

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) Enqueue(val int) {
	newNode := &Node{val: val, next: nil, prev: nil}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.head.prev = newNode
		newNode.next = q.head
		q.head = newNode
	}
}

func (q *Queue) Dequeue() error {
	if q.head == nil {
		return fmt.Errorf("queue is empty")
	}

	q.tail = q.tail.prev
	if q.tail != nil {
		q.tail.next = nil
	} else {
		q.head = nil
	}

	return nil
}

func (q *Queue) Print() {
	for aux := q.head; aux != nil; aux = aux.next {
		fmt.Print(aux.val, "")
	}
	fmt.Println()
}
