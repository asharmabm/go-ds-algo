package main

import "fmt"

func selectionSort(ar []int) []int { // O (n^2)
	var minIndex int
	//var temp int

	for i := 0; i < len(ar); i++ {
		minIndex = i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
		// can be done in this way also
		//temp = ar[i]
		//ar[i] = ar[minIndex]
		//ar[minIndex] = temp
	}

	return ar
}

func main() {
	ar := []int{99, 5, 2, 87, 67, 55, 4, 3}

	res := selectionSort(ar)

	fmt.Println("Sorted array", res)
}
