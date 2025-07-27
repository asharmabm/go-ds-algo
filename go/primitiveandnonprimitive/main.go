package main

import "fmt"

// Primitive Types
// These are the basic, built-in types provided by the Go language.
// They represent single, indivisible values.
// Predefined in Go: These types are defined by the Go compiler.

// Primitive Data Types in Go
// Category	    Data Types					    Examples
// Numeric	    int, float64, complex128	    int, int8, float32, byte
// Boolean	    bool							true, false
// String	    string							"hello", "Go"
// Character	rune						    'a', '✓' (represents Unicode code points)

// Non-Primitive Types
// These are more complex types and are derived from primitive types.
// They can hold multiple values and are often user-defined.

// Non-Primitive Data Types in Go
// Category	    	Data Types					Examples
// Aggregates		array, slice, struct		[]int, struct {name string}
// Reference		map, channel, pointer		map[string]int, chan int, *int
// User-Defined		interface, function, type	interface{}, func(int) int, type

// Summary Table
// Type			Comparison Basis			Direct Comparison with ==?
// Arrays		Based on values				Yes
// Structs		Based on field values		Yes
// Slices		Based on reference			No (except nil)
// Maps			Based on reference			No (except nil)
// Pointers		Based on memory address		Yes
// Channels		Based on reference			Yes

// Value-based types: Arrays, structs.
// Reference-based types: Slices, maps, channels, pointers.
// For reference-based types, equality checks only compare whether the references
// point to the same underlying data, not the content.

func main() {

	// We can compare two different arrays with == and we get either tru or false
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{4, 5, 6}

	fmt.Println(a == b) // Output: true (values are the same)
	fmt.Println(a == c) // Output: false (values are different)

	// Get the address of an array
	fmt.Println("Array a:", a)
	fmt.Println("Address of array a:", &a)
	fmt.Printf("Type of array a: %T\n", &a)

	// Access and modify elements via the pointer
	arr := [3]int{10, 20, 30}
	ptr := &arr // Pointer to the array
	(*ptr)[1] = 50
	// This also can be used
	// (&arr)[1] = 50

	fmt.Println("Array after modification:", arr) // Output: Array after modification: [10 50 30]

	// Comparing Slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	// Cannot compare slice1 == slice2, can be compared with slice1 == nil
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("Is slice1 nil?:", slice1 == nil)

	address := &slice1 // Get the address of the slice

	fmt.Println("slice1:", slice1)              // Output: Slice: [10 20 30]
	fmt.Println("Address:", address)            // Output: Address: (some memory address)
	fmt.Printf("slice1 address: %p\n", &slice1) // Output: Actual address of slice
	fmt.Printf("Type: %T\n", address)           // Output: Type: *[]int

	// Pointer: Points to the first element of the underlying array.
	slice := []int{10, 20, 30} // Define a slice

	fmt.Printf("Slice address: %p\n", &slice)               // Address of the slice
	fmt.Printf("Underlying array address: %p\n", &slice[0]) // Address of the underlying array

	// Empty Slice - For empty slice we cannot get the underlying array address via &emptySlice[0]
	var emptySlice []int

	fmt.Println("Empty slice:", emptySlice)
	fmt.Printf("Slice address: %p\n", &emptySlice)

	// Accessing the underlying array of an empty slice is invalid
	// fmt.Printf("Underlying array address: %p\n", &emptySlice[0]) // Runtime panic

	// For Structs

	p1 := person{"Ace", 20}
	p2 := person{"Ace", 20}
	p3 := person{"Bob", 21}

	fmt.Println(p1 == p2) // We can compare directly, it works based on each field values
	fmt.Println(p1 == p3)

	map1 := map[string]int{"Alice": 1}
	map2 := map[string]int{"Alice": 1}

	//fmt.Println(map1 == map2) // Cannot compare like this

	fmt.Println(&map1 == &map2)

	fmt.Printf("map1 address: %p\n", &map1) // Address of the map1
	fmt.Printf("map2 address: %p\n", &map2) // Address of the map2
	fmt.Println("Is map1 nil?:", map1 == nil)

	// make map1 nil
	map1 = nil
	fmt.Println("Is map1 nil?:", map1 == nil)

	// For Pointers
	ptrA := 3
	ptrB := 3
	ptr1 := &ptrA
	ptr2 := &ptrB
	fmt.Println("Is ptr1 == ptr2:", ptr1 == ptr2)

	// Make reference of ptr2 to &ptrA
	ptr2 = &ptrA
	fmt.Println("Is ptr1 == ptr2:", ptr1 == ptr2)

	// For Channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	fmt.Println("Is chan1 == chan2:", ch1 == ch2)
	fmt.Printf("Address of chan1 %p:\n", &ch1)
	fmt.Printf("Address of chan2 %p:\n", &ch2)

	// Make chan2 = chan1
	ch2 = ch1
	fmt.Println("Is chan1 == chan2:", ch1 == ch2)
	fmt.Printf("Address of chan1 %p:\n", &ch1)
	fmt.Printf("Address of chan2 %p:\n", &ch2)
}

type person struct {
	name string
	age  int
}
