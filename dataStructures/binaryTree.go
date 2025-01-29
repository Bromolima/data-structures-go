package dataStructures

import "fmt"

type Tree struct {
	root *Node
}

func (t *Tree) Insert(val int) {
	newNode := &Node{val: val, next: nil, prev: nil}
	if t.root == nil {
		t.root = newNode
		return
	}

	aux := t.root
	for aux.next != nil && aux.prev != nil {
		if newNode.val > aux.val {
			aux = aux.next
		} else {
			aux = aux.prev
		}
	}

	if newNode.val > aux.val {
		aux.next = newNode
		return
	}
	aux.prev = newNode
}

func (t *Tree) Search(val int) (*Node, error) {
	if t.root == nil {
		return nil, fmt.Errorf("tree is empty")
	}

	aux := t.root
	for aux != nil {
		if val == aux.val {
			return aux, nil
		}

		if val > aux.val {
			aux = aux.next
		}

		aux = aux.prev
	}

	return nil, fmt.Errorf("could not find node")
}
