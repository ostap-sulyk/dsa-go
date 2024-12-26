package trees

type BinarySearchTree struct{ Root *Node }

func NewEmptyBinaryTree() *BinarySearchTree     { return &BinarySearchTree{} }
func NewBinaryTree(value int) *BinarySearchTree { return &BinarySearchTree{Root: NewNode(value)} }

func (bst *BinarySearchTree) Insert(val int) bool {
	newNode := NewNode(val)

	if bst.Root == nil {
		bst.Root = newNode
		return true
	}

	temp := bst.Root
	for {
		if newNode.Value == temp.Value {
			return false
		}
		if newNode.Value < temp.Value {
			if temp.Left == nil {
				temp.Left = newNode
				return true
			}
			temp = temp.Left
		} else {
			if temp.Right == nil {
				temp.Right = newNode
				return true
			}
			temp = temp.Right
		}
	}
}

func (bst *BinarySearchTree) Contains(value int) bool {
	if bst.Root == nil {
		return false
	}

	temp := bst.Root

	for temp != nil {
		if value < temp.Value {
			temp = temp.Left
		} else if value > temp.Value {
			temp = temp.Right
		} else {
			return true
		}
	}
	return false
}
