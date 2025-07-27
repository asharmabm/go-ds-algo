package main

import (
	"fmt"
)

// QueueUsingStacks struct holds two stacks for queue implementation
type QueueUsingStacks struct {
	inputStack  []int // Stack for enqueue operations
	outputStack []int // Stack for dequeue operations
}

// Enqueue adds a new value to the queue
func (qs *QueueUsingStacks) Enqueue(value int) {
	qs.inputStack = append(qs.inputStack, value) // Push value onto the input stack
}

// dequeueFromInput transfers all elements from the input stack to the output stack
func (qs *QueueUsingStacks) dequeueFromInput() {
	// Transfer elements to reverse their order for FIFO behavior
	for len(qs.inputStack) > 0 {
		top := qs.inputStack[len(qs.inputStack)-1]           // Get the top element
		qs.inputStack = qs.inputStack[:len(qs.inputStack)-1] // Remove top element from input stack
		qs.outputStack = append(qs.outputStack, top)         // Push it onto the output stack
	}
}

// Dequeue removes and returns the front element of the queue
func (qs *QueueUsingStacks) Dequeue() int {
	// If output stack is empty, transfer elements from input stack
	if len(qs.outputStack) == 0 {
		qs.dequeueFromInput()
	}
	// Return the top element from the output stack if available
	if len(qs.outputStack) > 0 {
		top := qs.outputStack[len(qs.outputStack)-1]            // Get the front element
		qs.outputStack = qs.outputStack[:len(qs.outputStack)-1] // Remove it from output stack
		return top                                              // Return the dequeued value
	}
	return -1 // Return -1 if the queue is empty
}

func main() {
	queue := &QueueUsingStacks{} // Create a new queue instance

	// Enqueue elements
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Three elements are enqueued in the following order: 10, 20, and 30.")

	fmt.Println("\nNow, let's dequeue three times!")
	// Dequeue elements and print the results
	fmt.Println("\nDequeue:", queue.Dequeue()) // Should print 10
	fmt.Println("Dequeue:", queue.Dequeue())   // Should print 20
	fmt.Println("Dequeue:", queue.Dequeue())   // Should print 30
}

//type MyQueue struct {
//	stk1 []any
//	stk2 []any
//}
//
//
//func Constructor() MyQueue {
//	return MyQueue{
//		make([]any, 0),
//		make([]any, 0),
//	}
//}
//
//
//func (this *MyQueue) Push(x int)  {
//	this.stk1 = append(this.stk1, x)
//}
//
//
//func (this *MyQueue) Pop() int {
//	if len(this.stk2) == 0{
//		for len(this.stk1) >0{
//			top:= this.stk1[len(this.stk1)-1]
//			this.stk1 = this.stk1[:len(stk1)-1]
//			this.stk2 = append(this.stk2, top)
//		}
//	}
//
//	if len(this.stk2) >0 {
//		top := this.stk2[len(this.stk2)-1]
//		this.stk2 = this.stk2[:len(this.stk2)-1]
//		return top
//	}
//	return -1
//}
//
//
//func (this *MyQueue) Peek() int {
//	this.stk2[len(this.stk2)-1]
//}
//
//
//func (this *MyQueue) Empty() bool {
//	if len(this.stk1) == 0 && len(stk2) == 0{
//		return true
//	}
//
//	return false
//}
//
//
///**
// * Your MyQueue object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Push(x);
// * param_2 := obj.Pop();
// * param_3 := obj.Peek();
// * param_4 := obj.Empty();
// */
