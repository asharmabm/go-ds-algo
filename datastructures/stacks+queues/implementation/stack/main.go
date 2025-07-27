package main

import "fmt"

type node struct {
	value int
	next  *node
}

type stack struct {
	top    *node
	bottom *node
	length int
}

func newStack() *stack {
	return &stack{
		top:    nil,
		bottom: nil,
		length: 0,
	}
}

func (s *stack) peek() *node {
	return s.top
}

func (s *stack) push(data int) {
	newNode := &node{value: data, next: nil}

	if s.length == 0 {
		s.top = newNode
		s.bottom = newNode
	} else {
		next := s.top
		s.top = newNode
		s.top.next = next
	}
	s.length++
}

func (s *stack) pop() {
	if s.length == 0 {
		fmt.Println("Empty stack")
		return
	}
	if s.top == s.bottom {
		s.bottom = nil
	}
	s.top = s.top.next
	s.length--
}

func (s *stack) print() {
	current := s.top

	for current != nil {
		fmt.Print(current.value, ", ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	s := newStack()
	s.push(1)
	s.push(2)
	s.push(3)

	peek := s.peek()
	fmt.Println("Peek is", peek.value)
	s.print()

	fmt.Println("length is", s.length)
	s.pop()

	s.print()
	fmt.Println("length is", s.length)
	s.pop()
	s.print()

	s.pop()
	s.print()
}
