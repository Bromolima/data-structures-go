package main

import (
	"github.com/data-structs-go/dataStructures"
)

func main() {
	l := &dataStructures.List{}
	l.Add(7)
	l.Add(2)
	l.Add(1)
	l.Add(6)
	l.Add(4)
	l.Print()
}
