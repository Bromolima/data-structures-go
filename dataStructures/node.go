package dataStructures

type Node struct {
	val  int
	next *Node
	prev *Node
}

type SimpleNode struct {
	val  int
	next *SimpleNode
}
