package main

import "fmt"

type stack struct {
	array []int
}

func newStack() *stack {
	return &stack{
		array: make([]int, 0),
	}
}

func (s *stack) peek() int {
	return s.array[len(s.array)-1]
}

func (s *stack) push(data int) {
	s.array = append(s.array, data)
}

func (s *stack) pop() {
	s.array = s.array[:len(s.array)-1]
}

func (s *stack) print() {
	for _, v := range s.array {
		fmt.Print(v, ", ")
	}
	fmt.Println()
}

func main() {
	s := newStack()
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)

	peek := s.peek()
	fmt.Println("Peek is", peek)

	s.print()

	s.pop()
	s.print()

	s.pop()
	s.print()
}
