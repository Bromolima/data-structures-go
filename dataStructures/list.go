package dataStructures

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func swap(node_1, node_2 *Node) {
	aux := node_1
	node_1 = node_2
	node_2 = aux
}

type List struct {
	Head *Node
	Tail *Node
	len  int
}

func (l *List) Insert(val int) {
	newNode := &Node{Val: val}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.len = 1
	} else {
		aux := l.Head
		l.Head = newNode
		aux.Prev = newNode
		newNode.Next = aux
		l.len++
	}
}

func (l *List) Add(val int) {
	newNode := &Node{Val: val}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.len = 1
	} else {
		aux := l.Tail
		l.Tail = newNode
		aux.Next = newNode
		newNode.Prev = aux
		l.len++
	}
}

func (l *List) InsertInPosition(val int, pos int) error {
	newNode := &Node{Val: val}

	if pos < 0 || pos >= l.len {
		return fmt.Errorf("Invalid position")
	}

	if pos == 0 {
		l.Insert(val)
		return nil

	}

	if pos == l.len-1 {
		l.Add(val)
		return nil
	}

	aux := l.Head
	for i := 0; i < pos-1; i++ {
		aux = aux.Next
	}
	newNode.Next = aux.Next
	newNode.Prev = aux
	aux.Next = newNode
	aux.Next.Prev = newNode
	l.len++

	return nil

}

func (l *List) DeleteFromStart() error {
	if l.Head == nil {
		return fmt.Errorf("list is empty")
	}
	l.Head = l.Head.Next
	l.Head.Prev = nil
	l.len--

	return nil
}

func (l *List) DeleteFromEnd() error {
	if l.Head == nil {
		return fmt.Errorf("list is empty")
	}

	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.len--

	return nil
}

func (l *List) Print() {
	aux := l.Head
	print("[")
	for aux != nil {
		if aux.Next == nil {
			fmt.Print(aux.Val)
		} else {
			fmt.Print(aux.Val, " ")
		}
		aux = aux.Next
	}
	print("]")
}

func (l *List) PrintInverted() {
	aux := l.Tail
	for aux != nil {
		fmt.Print(aux.Val, " ")
		aux = aux.Prev
	}
}
