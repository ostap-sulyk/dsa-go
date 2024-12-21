package stack

type Node struct {
	Next  *Node
	Value int
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}
