package main

import "fmt"

func main() {
	// create an array of size 4
	var array [4]string

	// This also can be followed
	// array := [4]string{"a", "b", "c", "d"}

	// prints default value of type when not initialised or inserted any value.
	// cannot use make to create arrays. only used for slices, maps and channels.

	array[0] = "a"
	array[1] = "b"
	array[2] = "c"
	array[3] = "d"

	fmt.Println(array)

	fmt.Println(array[1])

	// Slices (Dynamic Arrays)
	// create a slice
	var slice []int

	// below also creates slice with capacity of 100 and length of 0
	// slice := make([]int, 0, 100)

	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	// when we append to a slice and if it exceeds the capacity then internally
	// it creates new array with usually double the capacity of the existing array
	// and copies the data with appended element. So time complexity of O(n) and usually
	// space complexity will be more since new underlying array is created.

	// Slices will use pointer to reference for the underlying array

	slice = append(slice, 1, 2)

	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	slice = append(slice, 3)

	// capacity doubled from 2 to 4
	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	slice = []int{}

	// When not initialised with length of slice we cannot directly assign value like s[0] = 1 like
	// Otherwise it panics with index out of range
	slice = make([]int, 2)

	slice[0] = 1
	slice[1] = 2

	fmt.Printf("slice: %v\n", slice)

	// Once we exceed length of slice use append to add elements
	slice = append(slice, 3)

	fmt.Printf("slice: %v\n", slice)

	// Append more values at a time
	slice = append(slice, 4, 5, 6)

	fmt.Printf("slice: %v\n", slice)

	// Lookup for a value in the 2nd index
	slice = []int{}
	if len(slice) >= 2 {
		fmt.Printf("Element %d found at index %d\n", slice[2], 2)
	} else {
		fmt.Printf("No element found at index %d\n", 2)
	}

	// Add value at the beginning
	slice = []int{1, 2, 3}
	slice1 := []int{0}
	slice = append(slice1, slice...)

	fmt.Printf("slice after adding at first index: %v\n", slice)

	// Remove second index element
	slice = []int{1, 2, 3, 4, 5}
	slice = append(slice[:2], slice[3:]...)

	fmt.Printf("slice after removal: %v\n", slice)

	// Add an element at the fourth index
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice1 = []int{10, 11, 12}
	slice = append(slice[:4], append([]int{100}, slice[4:]...)...)

	fmt.Printf("slice after adding an element in the middle: %v\n", slice)

	// Add an element at the third index and new slice
	index := 3
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice1 = []int{10, 11, 12}
	slice = append(slice[:index], append([]int{100}, append(slice[index:], slice1...)...)...)

	fmt.Printf("slice after adding an element in the middle and a slice at the end: %v\n", slice)
}
