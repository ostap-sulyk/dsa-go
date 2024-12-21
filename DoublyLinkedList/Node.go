package doublylinkedlist

type Node struct {
	Prev, Next *Node
	Value      int
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}
