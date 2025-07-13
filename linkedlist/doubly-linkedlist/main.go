package main

import "fmt"

type linkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	data     any
	next     *node
	previous *node
}

func (ll *linkedList) Append(data int) {
	newNode := &node{data: data, next: nil, previous: nil}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		ll.length = 1
		return
	} else {
		newNode.previous = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
		ll.length++
	}
}

func (ll *linkedList) Prepend(data int) {
	newNode := &node{data: data, next: nil, previous: nil}

	newNode.next = ll.head
	ll.head.previous = newNode
	ll.head = newNode
	ll.length++
}

// 1, 2, 3, nil

func (ll *linkedList) Insert(index, data int) {
	if index >= ll.length {
		fmt.Errorf("index out of range, so adding at the end")
		ll.Append(data)
		return
	}

	if index == 0 {
		ll.Prepend(data)
	}

	newNode := &node{data: data, previous: nil, next: nil}
	current := ll.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	next := current.next
	current.next = newNode
	newNode.previous = current
	next.previous = newNode
	newNode.next = next
	ll.length++
}

func (ll *linkedList) Remove(index int) {
	if index >= ll.length {
		fmt.Errorf("index out of range, cannot delete")
		return
	}

	if index == 0 {
		ll.head = ll.head.next
		ll.head.previous = nil
		ll.length--
		return
	}

	current := ll.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	prevNode := current
	current.next = current.next.next
	current.next.previous = prevNode

	ll.length--
}

func (ll *linkedList) print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}

	fmt.Println("nil")
}

func main() {
	l := linkedList{}
	l.Append(1)
	//l.print()
	l.Append(2)
	l.Append(3)
	l.Append(4)

	l.Prepend(10)
	l.Append(88)
	l.print()

	l.Insert(2, 99)
	l.print()

	l.Remove(2)
	l.print()

	l.Remove(0)
	l.print()

	l.Remove(1)
	l.print()

}
