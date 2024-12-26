package trees

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node { return &Node{value, nil, nil} }
