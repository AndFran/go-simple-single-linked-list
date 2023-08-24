package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value int) {
	n := &Node{value: value}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
}

func (l *List) Display() {
	fmt.Println("------------------")
	current := l.head
	for current != nil {
		fmt.Println(current)
		current = current.next
	}
	fmt.Println("------------------")
}

func main() {
	ll := &List{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	ll.Push(4)
	ll.Push(5)
	ll.Display()
}
