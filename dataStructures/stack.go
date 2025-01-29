package dataStructures

import "fmt"

type Stack struct {
	head *SimpleNode
}

func (s *Stack) Push(val int) {
	newNode := &SimpleNode{val: val, next: nil}
	if s.head == nil {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
}

func (s *Stack) Pop() error {
	if s.head == nil {
		return fmt.Errorf("stack is empty")
	} else {
		s.head = s.head.next
	}
	return nil
}

func (s *Stack) Print() {
	for aux := s.head; aux != nil; aux = aux.next {
		fmt.Print(aux.val, " ")
	}
	fmt.Println()
}
