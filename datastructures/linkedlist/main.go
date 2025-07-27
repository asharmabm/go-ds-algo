package main

import (
	"fmt"
	"strings"
	"time"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

//func (ll *LinkedList) Add(value int) {
//	newNode := &Node{data: value}
//
//	if ll.head == nil {
//		ll.head = newNode
//	} else {
//		current := ll.head
//		for current.next != nil {
//			current = current.next
//		}
//		current.next = newNode
//	}
//}

func (ll *LinkedList) Append(value int) {
	newNode := &Node{data: value}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		ll.length++
	} else {
		ll.tail.next = newNode
		ll.length++
		ll.tail = newNode
	}
}

func (ll *LinkedList) Prepend(value int) {
	newNode := &Node{data: value}
	newNode.next = ll.head

	//ll.last = newNode.next
	ll.head = newNode
	ll.length++
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}

	fmt.Println("nil")
}

func (ll *LinkedList) Insert(index int, value int) {
	if index >= ll.length {
		fmt.Println(fmt.Errorf("index out of range so adding at the end"))
		ll.Append(value)
		return
	}
	if index == 0 {
		ll.Prepend(value)
		return
	}

	// index 2, value 87
	// 1, 2, 3, 4, 5
	// 1, 2, 87, 3, 4, 5
	newNode := &Node{data: value}
	current := ll.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	next := current.next
	current.next = newNode
	//current.next.next = next
	// can use this also
	newNode.next = next

	ll.length++
}

func (ll *LinkedList) Delete(index int) {
	if index >= ll.length {
		fmt.Println(fmt.Errorf("error: index out of range"))
		return
	}

	count := 0
	current := ll.head
	for count != index-1 {
		current = current.next
		count++
	}

	current.next = current.next.next
	// can be done like this also
	//unwantedNode := current.next
	//current.next = unwantedNode.next

	ll.length--
}

func (ll *LinkedList) Reverse() {
	if ll.length == 1 {
		return
	}

	first := ll.head
	ll.tail = ll.head
	second := first.next

	for second != nil { // 1, 2, 3
		temp := second.next // 3
		second.next = first // third = 1
		fmt.Printf("First first %+v\n", first)
		first = second // first = 2
		fmt.Printf("Second first %+v\n", first)
		second = temp // second = 3  // if arranged in order 2, 3, 1

		fmt.Println("second", temp)
		fmt.Println(strings.Repeat("*", 20))
		fmt.Println()
	}

	ll.head.next = nil
	ll.head = first
}

// Reverse1 - Best one to understand
func (ll *LinkedList) Reverse1() {
	cur := ll.head
	var prev *Node = nil

	for cur != nil {
		temp := cur.next

		cur.next = prev
		prev = cur

		cur = temp
	}

	ll.head = prev
}

func main() {
	l := &LinkedList{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	l.Print()
	//l = &LinkedList{}

	//for i := 0; i < 100; i++ {
	//	l.Prepend(i + 1)
	//}
	//l.Append(1)
	//l.Append(2)
	//l.Append(3)
	l.Insert(3, 4)
	//l.Insert(12, 87)

	l.Print()

	//l.Delete(1)
	l.Delete(3)

	l.Print()

	l.Delete(6)

	l.Print()

	//l.Reverse()
	l.Print()
	l.Reverse1()

	l.Print()

	msoLP := time.Date(2025, time.April, 22, 15, 06, 00, 00, time.Local)
	//mpLPFetchTime := time.Date(2025, time.April, 22, 15, 10, 22, 00, time.UTC)

	fmt.Println("Debug, msoLP", msoLP)

	v := time.Until(msoLP)

	// 19:51:20.923594
	// 15:06:00

	// 15:06:00 - 19:51:20
	// msoLP - current.Now
	fmt.Println("Curreent time", time.Now())

	//fmt.Println("Debug minus value", t)
	fmt.Println("Final", v)
}
