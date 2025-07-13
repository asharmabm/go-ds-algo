package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

type Node struct {
	key   any
	value any
	next  *Node
}

type CustomMap struct {
	capacity int
	size     int
	bucket   []*Node
}

var defaultCapacity = 50

func New() *CustomMap {
	return &CustomMap{
		capacity: defaultCapacity,
		size:     0,
		bucket:   make([]*Node, defaultCapacity),
	}
}

func (m *CustomMap) Hash(key any) int {

	if key == "Mangalore" || key == "Mangalore1" {
		key = "Akshay"
	}
	str := fmt.Sprintf("%v", key)
	bytes := md5.Sum([]byte(str))
	decimalBigEndian := binary.BigEndian.Uint16(bytes[:])

	return int(decimalBigEndian) % defaultCapacity
}

func (m *CustomMap) Put(key any, value any) {
	index := m.Hash(key)

	newNode := &Node{key: key, value: value}
	if m.bucket[index] == nil {
		m.bucket[index] = newNode
	} else {
		current := m.bucket[index]
		if key == "Mangalore1" {
			fmt.Printf("Address of current %p, current.next %p \n", current, current.next)
		}
		for current != nil {
			if current.key == key {
				current.value = value
				return
			}

			if current.next == nil {
				break
			}

			current = current.next

			if key == "Mangalore1" {
				fmt.Printf("Address of current %p \n", current)
			}
		}

		current.next = newNode

	}
	m.size++
}

func (m *CustomMap) Get(key any) any {
	node := m.getNodeByKey(key)

	if node != nil {
		return node.value
	}
	return nil
}

func (m *CustomMap) getNodeByKey(key any) *Node {
	index := m.Hash(key)
	current := m.bucket[index]

	for current != nil {
		if current.key == key {
			return current
		}

		current = current.next
	}
	return nil
}

func main() {
	m := New()
	m.Put("Akshay", 1)
	m.Put("Rohit", 2)
	m.Put("Mahesh", 3)

	v := m.Get("Akshay")
	fmt.Println("Value is: ", v)

	// Update value for same key
	m.Put("Akshay", 100)
	v = m.Get("Akshay")
	fmt.Println("Value after update: ", v)

	m.Put("Mangalore", 900)
	v = m.Get("Mangalore")
	fmt.Println("Value after update: ", v)

	m.Put("Mangalore1", 9000)
	v = m.Get("Mangalore1")
	fmt.Println("Value after update: ", v)
}
