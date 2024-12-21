package doublylinkedlist

import (
	"strconv"
)

type DoublyLinkedList struct {
	head, tail *Node
	length     int
}

func NewDoubleLinkedList() *DoublyLinkedList { return &DoublyLinkedList{} }

func NewDoubleLinkedListWithValue(value int) *DoublyLinkedList {
	node := NewNode(value)
	return &DoublyLinkedList{head: node, tail: node, length: 1}
}

func (dll *DoublyLinkedList) Append(value int) {
	newNode := NewNode(value)
	if dll.length == 0 {
		dll.head, dll.tail = newNode, newNode
	} else {
		dll.tail.Next = newNode
		newNode.Prev = dll.tail
		dll.tail = newNode
	}
	dll.length++
}

func (dll *DoublyLinkedList) RemoveLast() *Node {
	if dll.length == 0 {
		return nil
	}

	var temp *Node = nil
	if dll.length == 1 {
		temp = dll.tail
		dll.head, dll.tail = nil, nil
	} else {
		temp = dll.tail
		dll.tail = dll.tail.Prev
		dll.tail.Next = nil
	}

	temp.Prev = nil
	dll.length--

	return temp
}

func (dll *DoublyLinkedList) PrePend(value int) {
	newNode := NewNode(value)
	if dll.length == 0 {
		dll.head, dll.tail = newNode, newNode
	} else {
		newNode.Next = dll.head
		dll.head.Prev = newNode
		dll.head = newNode
	}
	dll.length++
}

func (dll *DoublyLinkedList) RemoveFirst() *Node {
	var temp *Node
	if dll.length == 0 {
		return nil
	} else if dll.length == 1 {
		temp = dll.head
		dll.head, dll.tail = nil, nil
	} else {
		temp = dll.head
		dll.head = dll.head.Next
		dll.head.Prev, temp.Next = nil, nil
	}
	dll.length--
	return temp
}

func (dll *DoublyLinkedList) Get(index int) *Node {
	if index < 0 || index >= dll.length {
		return nil
	}
	var temp *Node
	if index <= dll.length/2 {
		temp = dll.head
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
	} else {
		temp = dll.tail
		for i := 0; i > index; i++ {
			temp = temp.Prev
		}
	}

	return temp
}

func (dll DoublyLinkedList) Set(index, value int) bool {
	node := dll.Get(index)
	if node != nil {
		node.Value = value
		return true
	}
	return false
}

func (dll *DoublyLinkedList) Insert(index, value int) bool {
	if index < 0 || index > dll.length {
		return false
	} else if index == 0 {
		dll.PrePend(value)
	} else if index == dll.length {
		dll.Append(value)
	} else {
		newNode := NewNode(value)
		after := dll.Get(index)
		before := after.Prev

		newNode.Prev, newNode.Next = before, after
		before.Next, after.Prev = newNode, newNode
		dll.length++
	}
	return true
}

func (dll *DoublyLinkedList) Remove(index int) *Node {
	if index < 0 || index >= dll.length {
		return nil
	} else if index == 0 {
		return dll.RemoveFirst()
	} else if index == dll.length-1 {
		return dll.RemoveLast()
	} else {
		node := dll.Get(index)
		before, after := node.Prev, node.Next
		before.Next, after.Prev = after, before
		node.Prev, node.Next = nil, nil
		return node
	}
}

// printing
func (dll *DoublyLinkedList) PrintList() {
	s := ""
	temp := dll.head
	for temp != nil {
		s += strconv.Itoa(temp.Value) + " -> "
		temp = temp.Next
	}
	s += "nil"
	println(s)
}
func (dll *DoublyLinkedList) PrintHead() {
	if dll.head != nil {
		println("Head :", dll.head.Value)
	} else {
		println("Head : nil")
	}
}
func (dll *DoublyLinkedList) PrintTail() {
	if dll.tail != nil {
		println("Tail :", dll.tail.Value)
	} else {
		println("Tail : nil")
	}
}
func (dll *DoublyLinkedList) PrintLength() { println("Length :", dll.length) }
