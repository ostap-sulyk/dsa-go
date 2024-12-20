package linkedlist

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}