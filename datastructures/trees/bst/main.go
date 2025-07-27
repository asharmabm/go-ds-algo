package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) {
	newNode := &Node{data: value}
	if t.root == nil {
		t.root = newNode
	} else {
		current := t.root
		for {
			if value < current.data {
				// Left
				if current.left == nil {
					current.left = newNode
					return
				}
				current = current.left
			} else {
				// Right
				if current.right == nil {
					current.right = newNode
					return
				}
				current = current.right
			}
		}
	}
}

func (t *Tree) Lookup(value int) *Node {
	if t.root == nil {
		return nil
	}

	current := t.root

	for current != nil {
		if value < current.data {
			current = current.left
		} else if value > current.data {
			current = current.right
		} else if current.data == value {
			return current
		}
	}

	return nil
}

func (t *Tree) Remove(value int) bool {
	if t.root == nil {
		return false
	}

	current := t.root
	var parent *Node

	for current != nil {
		if value < current.data {
			parent = current
			current = current.left
		} else if value > current.data {
			parent = current
			current = current.right
		} else if current.data == value {
			if current.right == nil {
				if parent == nil {
					t.root = current.left
				} else {
					if current.data < parent.data {
						parent.left = current.left
					} else if current.data > parent.data {
						parent.right = current.left
					}
				}
			}
		}
	}

	return false
}

func (t *Tree) PrintInOrder(node *Node) {
	if node == nil {
		return
	}
	t.PrintInOrder(node.left)  // Step 1: Visit left subtree
	fmt.Println(node.data)     // Step 2: Process current node (print data)
	t.PrintInOrder(node.right) // Step 3: Visit right subtree
}

func main() {
	tree := &Tree{}

	tree.Insert(9)

	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)

	node := tree.Lookup(9)
	fmt.Println("Found", node)
}
