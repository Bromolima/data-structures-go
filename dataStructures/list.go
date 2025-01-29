package dataStructures

import "fmt"

type List struct {
	head *Node
	tail *Node
	len  int
}

func (l *List) Insert(val int) {
	newNode := &Node{val: val, next: nil, prev: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.len = 1
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
		l.len++
	}
}

func (l *List) Add(val int) {
	newNode := &Node{val: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.len = 1
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		l.len++
	}
}

func (l *List) InsertInPosition(val int, pos int) error {
	newNode := &Node{val: val}

	if pos < 0 || pos >= l.len {
		return fmt.Errorf("invalid position")
	}

	if pos == 0 {
		l.Insert(val)
		return nil

	}

	if pos == l.len-1 {
		l.Add(val)
		return nil
	}

	aux := l.head
	for i := 0; i < pos-1; i++ {
		aux = aux.next
	}
	newNode.next = aux.next
	newNode.prev = aux
	aux.next = newNode
	aux.next.prev = newNode
	l.len++

	return nil

}

func (l *List) DeleteFromStart() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.len--

	return nil
}

func (l *List) DeleteFromEnd() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}

	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	l.len--

	return nil
}

func (l *List) Print() {
	fmt.Print("[")
	for aux := l.head; aux != nil; aux = aux.next {
		if aux.next == nil {
			fmt.Print(aux.val)
		} else {
			fmt.Print(aux.val, " ")
		}
	}
	fmt.Println("]")
}

func (l *List) PrintInverted() {
	fmt.Print("[")
	for aux := l.tail; aux != nil; aux = aux.prev {
		if aux.prev == nil {
			fmt.Print(aux.val)
		} else {
			fmt.Print(aux.val, " ")
		}
	}
	fmt.Println("]")
}

func (l *List) QuickSort() error {
	if l.head == nil || l.tail == nil || l.len == 0 {
		return fmt.Errorf("list is empty")
	}
	quickSort(l.head, l.tail)
	return nil
}

func partition(start, end *Node) *Node {
	if start == nil || end == nil {
		return nil
	}
	pivot := end
	i := start.prev

	for j := start; j != end && j != nil; j = j.next {
		if j.val <= pivot.val {
			if i == nil {
				i = start
			} else {
				i = i.next
			}
			i.val, j.val = j.val, i.val
		}
	}

	if i == nil {
		i = start
	} else {
		i = i.next
	}

	i.val, pivot.val = pivot.val, i.val
	return i
}

func quickSort(start, end *Node) {
	if start == nil || end == nil || start == end || end.next == start {
		return
	}
	pivot := partition(start, end)
	quickSort(start, pivot.prev)
	quickSort(pivot.next, end)
}
