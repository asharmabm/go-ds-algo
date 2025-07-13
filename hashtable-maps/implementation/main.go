package main

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity uint64 = 1 << 10

type Node struct {
	key   any
	value any
	next  *Node
}

type CustomMap struct {
	capacity uint64
	size     uint64
	bucket   []*Node
}

// DefaultNew returns a new CustomMap instance with default values
func DefaultNew() *CustomMap {
	return &CustomMap{
		capacity: defaultCapacity,
		size:     0,
		bucket:   make([]*Node, defaultCapacity),
	}
}

// New creates a new custom map instance with the specified size and capacity
func New(size, capacity uint64) *CustomMap {
	return &CustomMap{
		capacity: capacity,
		size:     size,
		bucket:   make([]*Node, capacity),
	}
}

// hash generates a hash value for the given key
func (m *CustomMap) hash(key any) uint64 {
	if key == "Mangalore" {
		key = "Akshay"
	}
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := h.Sum64()
	return (m.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

// Put inserts a new key-value pair into a custom map
func (m *CustomMap) Put(key, value any) {
	index := m.hash(key)
	if m.bucket[index] == nil {
		m.bucket[index] = &Node{key: key, value: value}
	} else {
		current := m.bucket[index]
		for {
			if current.key == key {
				current.value = value
				return
			}
			if current.next == nil {
				break
			}
			current = current.next
		}
		current.next = &Node{key: key, value: value}
	}

	m.size++
}

// Get returns the value associated with the given key
func (m *CustomMap) Get(key any) any {
	node := m.getNodeByKey(key)

	if node != nil {
		return node.value
	}

	return nil
}

// getNodeByKey finds the node associated with the given key
func (m *CustomMap) getNodeByKey(key any) *Node {
	index := m.hash(key)
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
	m := DefaultNew()
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
}
