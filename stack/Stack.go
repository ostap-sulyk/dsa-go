package stack

type Stack struct {
	top    *Node
	height int
}

func NewStack(value int) *Stack { return &Stack{top: NewNode(value), height: 1} }
func NewEmptyStack() *Stack     { return &Stack{} }

func (s *Stack) Push(value int) {
	node := NewNode(value)
	if s.height == 0 {
		s.top = node
	} else {
		node.Next = s.top
		s.top = node
	}
	s.height++
}

func (s *Stack) Pop() *Node {
	if s.height == 0 {	return nil	}

	temp := s.top
	s.top = temp.Next
	temp.Next = nil
	s.height--

	return temp
}

func (s *Stack) PrintStack() {
	temp := s.top
	for temp != nil {
		println(temp.Value)
		println("↓")
		temp = temp.Next
	}
}
func (s *Stack) PrintTop()    { println("Top:", s.top.Value) }
func (s *Stack) PrintHeight() { println("Height:", s.height) }
