package main

import (
	"fmt"
	"os"
	"strings"
)

// CustomArray - Define a custom array structure
type CustomArray struct {
	data     []int // Underlying storage for elements
	capacity int   // Maximum capacity of the array
	size     int   // Current number of elements in the array
}

// NewCustomArray - Initialise a new CustomArray
func NewCustomArray(capacity int) *CustomArray {
	return &CustomArray{
		data:     make([]int, capacity), // Preallocate storage
		capacity: capacity,
		size:     0,
	}
}

// Add an element at a specific index
func (ar *CustomArray) Add(index int, value int) error {
	// Ensure the index is within the bounds
	if index > ar.size || index < 0 || ar.size >= ar.capacity {
		return fmt.Errorf("index out of bounds or array size is full")
	}

	// Shift the elements to the right to make space
	for i := ar.size; i > index; i-- {
		ar.data[i] = ar.data[i-1] // Update current index value to previous index value
	}

	ar.data[index] = value
	ar.size++ // Increment the size
	return nil
}

// Get - Retrieve an element at a specific index
func (ar *CustomArray) Get(index int) (int, error) {
	if index < 0 || index > ar.size {
		return 0, fmt.Errorf("index out of bounds")
	}
	return ar.data[index], nil
}

// Remove an element at a specific index
func (ar *CustomArray) Remove(index int) error {
	if index < 0 || index > ar.size {
		return fmt.Errorf("index out of bounds")
	}

	// Shift elements to the left to fill the gap
	for i := index; i < ar.size-1; i++ {
		ar.data[i] = ar.data[i+1] // Update next index value to current index
	}
	ar.data[ar.size-1] = 0

	ar.size-- // Decrement the size

	fmt.Println("hfgg", ar)

	return nil
}

func (ar *CustomArray) Display() {
	fmt.Print("CustomArray: [")
	for i := 0; i < ar.size; i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(ar.data[i])
	}
	fmt.Println("]")
}

func main() {

	ar := NewCustomArray(5)

	ar.Add(0, 10)
	ar.Add(1, 20)
	ar.Add(2, 30)
	ar.Display()

	ar.Add(1, 100)
	ar.Display()

	ar.Add(3, 200)
	ar.Display()

	ar.Add(0, 300)
	ar.Display()
	//
	//ar.Add(500)
	//ar.Remove(1)
	ar.Display()

	fmt.Println(strings.Repeat("*", 100))

	ar.Add(0, 10)

	ar.Add(1, 20)
	ar.Add(2, 30)
	//ar.Add(3, 40)
	//err := ar.Add(4, 50)
	ar.Display()

	//err = ar.Add(5, 60)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//ar.Display()

	// Insert an element at a specific index
	err := ar.Add(1, 15) // Insert 15 at index 1
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ar.Display() // Output: CustomArray: [10, 15, 30]

	// Get an element at a specific index
	value, err := ar.Get(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element at index 2:", value) // Output: Element at index 2: 20
	}

	// Remove an element
	ar.Remove(1) // Remove the element at index 1
	ar.Display() // Output: CustomArray: [10, 30, 40, 50]
}
