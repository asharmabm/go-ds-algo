package main

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	capacity int
	size     int
	data     []int
}

func New(capacity int) *DynamicArray {
	return &DynamicArray{
		capacity: capacity,
		size:     0,
		data:     make([]int, capacity),
	}
}

func (ar *DynamicArray) Update(index int, element int) error {
	if index >= ar.size || index < 0 {
		return errors.New("index out of range")
	}

	newSize := ar.size
	if ar.size == ar.capacity {
		newSize--
	}

	// index = 2, element = 100
	// 1,2, 3,4
	for i := newSize; i > index; i-- {
		ar.data[i] = ar.data[i-1]
	}

	ar.data[index] = element

	return nil
}

func (ar *DynamicArray) Add(element int) error {
	if ar.size == ar.capacity {
		ar.capacity = ar.capacity * 2
		newAr := make([]int, ar.capacity)
		copy(newAr, ar.data)
		ar.data = newAr
	}

	ar.data[ar.size] = element
	ar.size++
	return nil
}

func (ar *DynamicArray) Remove(index int) {
	if index > ar.capacity {
		fmt.Errorf("index out of bounds")
	}
	for i := index; i < len(ar.data)-1; i++ {
		ar.data[i] = ar.data[i+1]
	}
	ar.size--
}

func (ar *DynamicArray) Print() {
	for i := 0; i < ar.size; i++ {
		if i == 0 {
			fmt.Print("[")
		}
		if i == ar.size-1 {
			fmt.Println(fmt.Sprintf("%d] ", ar.data[i]))
		} else {
			fmt.Print(fmt.Sprintf("%d, ", ar.data[i]))
		}
	}
}

func main() {
	ar := New(5)
	ar.Add(10)
	ar.Add(20)
	ar.Add(30)
	ar.Add(40)
	ar.Add(50)
	ar.Add(60)

	ar.Print()

	ar.Update(1, 100)
	ar.Print()

	err := ar.Update(3, 200)
	fmt.Println(err)
	ar.Print()

	ar.Update(0, 300)
	ar.Print()
	//
	//ar.Add(500)
	ar.Remove(1)
	ar.Print()

	//fmt.Println("Sfddfsfdf", ar.size)
}
