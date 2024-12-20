package linkedlist

import "strconv"

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList(value int) *LinkedList {
	node := NewNode(value)
	return &LinkedList{head: node, tail: node, length: 1}
}

func (l *LinkedList) Append(value int) {
	newNode := NewNode(value)
	if l.length == 0 {
		l.head, l.tail = newNode, newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.length++
}

func (l *LinkedList) Prepend(value int) {
	newNode := NewNode(value)
	if l.length == 0 {
		l.head, l.tail = newNode, newNode
	} else {
		newNode.Next = l.head
		l.head = newNode
	}
	l.length++
}

func (l *LinkedList) RemoveLast() *Node {
	if l.length == 0 {
		return nil
	}
	pre, temp := l.head, l.head

	for temp.Next != nil {
		pre = temp
		temp = temp.Next
	}
	l.tail = pre
	l.tail.Next = nil
	l.length--
	if l.length == 0 {
		l.head, l.tail = nil, nil
	}
	return temp
}

func (l *LinkedList) RemoveFirst() *Node {
	if l.length == 0 {
		return nil
	}
	temp := l.head
	l.head = l.head.Next
	temp.Next = nil
	l.length--

	if l.length == 0 {
		l.tail = nil
	}
	return temp
}

func (l *LinkedList) Get(index int) *Node {
	if index < 0 || index >= l.length {
		return nil
	}
	temp := l.head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp
}

func (l *LinkedList) Set(index, value int) bool {
	node := l.Get(index)
	if node == nil {
		return false
	}
	node.Value = value
	return true
}

func (l *LinkedList) Insert(index, value int) bool {
	if index < 0 || index > l.length {
		return false
	} else if index == 0 {
		l.Prepend(value)
	} else if index == l.length {
		l.Append(value)
	} else {
		node := NewNode(value)
		temp := l.head
		for i := 0; i < index-1; i++ {
			temp = temp.Next
		}
		node.Next = temp.Next
		temp.Next = node
		l.length++
	}

	return true
}

func (l *LinkedList) Remove(index int) bool {
	if index < 0 || index >= l.length {
		return false
	} else if index == 0 {
		l.RemoveFirst()
	} else if index == l.length-1 {
		l.RemoveLast()
	} else {
		prev := l.Get(index - 1)
		temp := prev.Next
		prev.Next = temp.Next
		temp.Next = nil
		l.length--
	}
	return true
}

func (l *LinkedList) Reverse() {
	var before, after *Node
	l.head, l.tail = l.tail, l.head
	temp := l.tail

	for i := 0; i < l.length; i++ {
		after = temp.Next
		temp.Next = before
		before, temp = temp, after
	}
}

// Printing
func (l *LinkedList) PrintList() {
	current := l.head
	output := ""
	for current != nil {
		output += strconv.Itoa(current.Value) + " -> "
		current = current.Next
	}
	output += "nil"
	println(output)
}

func (l *LinkedList) PrintHead()   { println("Head:", l.head.Value) }
func (l *LinkedList) PrintTail()   { println("Tail:", l.tail.Value) }
func (l *LinkedList) PrintLength() { println("Size:", l.length) }
