package queue

import (
	"fmt"
	"strconv"
)

type Queue struct {
	first, last *Node
	length      int
}

func NewEmptyQue() *Queue { return &Queue{} }

func NewQueue(value int) *Queue {
	n := NewNode(value)
	return &Queue{n, n, 1} // WHAT??? I could do that?????
}

func (q *Queue) Enque(value int) {
	n := NewNode(value)
	if q.length == 0 {
		q.first, q.last = n, n
	} else {
		q.last.Next, q.last = n, n
	}
	q.length++
}

func (q *Queue) DeQueue() *Node {
	if q.length == 0 {
		return nil
	}
	temp := q.first
	if q.length == 1 {
		q.first, q.last = nil, nil
	} else {
		q.first = temp.Next
		temp.Next = nil
	}
	q.length--
	return temp
}

// printing
func (q *Queue) Print() {
	s := ""
	for n := q.first; n != nil; n = n.Next { // WHAT??? I could do that too?????
		s += strconv.Itoa(n.Value) + " "
	}
	println(s)
}
func (q *Queue) PrintLenght() { println("Length:", q.length) }
func (q *Queue) PrintFirst()  { fmt.Printf("First: %+v\n", q.first) }
func (q *Queue) PrintLast()   { fmt.Printf("Last: %+v\n", q.last) }
