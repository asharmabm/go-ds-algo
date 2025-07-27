package main

import "fmt"

type node struct {
	data int
	next *node
}

type queue struct {
	first  *node
	last   *node
	length int
}

func newQueue() *queue {
	return &queue{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

func (q *queue) peek() *node {
	return q.first
}

func (q *queue) Enqueue(data int) {
	newNode := &node{data: data}

	if q.length == 0 {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}

	q.length++
}

func (q *queue) Dequeue() {
	if q.first == nil {
		return
	}

	if q.first == q.last {
		q.last = nil
	}
	q.first = q.first.next
	q.length--
}

func (q *queue) print() {
	current := q.first

	for current != nil {
		fmt.Print(current.data, ", ")
		current = current.next
	}

	fmt.Println()
}

func main() {
	q := newQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	q.print()
	peek := q.peek()

	fmt.Println("Peek is", peek)

	q.Dequeue()

	q.print()

	peek = q.peek()
	fmt.Println("Peek is", peek)

	q.Dequeue()
	fmt.Println("a")
	q.print()

	q.Dequeue()
	fmt.Println("b")
	q.print()

	q.Dequeue()
	fmt.Println("c")
	q.print()
}
