package queue

type Node struct {
	Next  *Node
	Value int
}

func NewNode(value int) *Node { return &Node{Value: value} }
