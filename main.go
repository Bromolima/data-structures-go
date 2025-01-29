package main

import (
	"github.com/data-structs-go/dataStructures"
)

func main() {
	l := &dataStructures.List{}
	slice := []int{3, 5, 7, 8, 9, 4, 2, 10, 18, 9}
	for _, number := range slice {
		l.Add(number)
	}
	l.Print()
	l.DeleteFromStart()
	l.QuickSort()
	l.Print()
}
