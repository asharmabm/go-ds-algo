package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// Add appends a new node at the end of the list
func (ll *LinkedList) Add(value int) {
	newNode := &Node{data: value}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Print displays the elements of the linked list
func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// Delete removes a node with the specified value
func (ll *LinkedList) Delete(value int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == value {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

func main() {
	ll := &LinkedList{}
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Print()

	ll.Delete(20)
	ll.Print()

	var m []int

	if m == nil {
		fmt.Println("Camer her er")
	}
	fmt.Printf("%T, value is %v", m, m)
}
