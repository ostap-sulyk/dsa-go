package main

import doublylinkedlist "dsa-go/DoublyLinkedList"

func main() {
	dll := doublylinkedlist.NewDoubleLinkedList(7)

	dll.PrintHead()
	dll.PrintTail()
	dll.PrintSize()
	dll.PrintList()
}
