package main

import (
	"fmt"
)

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(value any) {
	newNode := &Node{data: value, next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *LinkedList) Print() {
	current := l.head

	for current != nil {
		fmt.Print(current.data, ", ")
		current = current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) Delete(index int) {

	if l.head == nil || index < 0 {
		_ = fmt.Errorf("cannot delete on a nil list or index less than 0")
	}
	current := l.head

	fmt.Println("Index is", index)
	for i := 0; i <= index; i++ {
		if i == index {
			current.data = current.next.data
			current.next = current.next.next
			break
		}
		current = current.next
	}

}

func main() {
	l := &LinkedList{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Print()

	//l.Delete(3)
	l.Print()
}
