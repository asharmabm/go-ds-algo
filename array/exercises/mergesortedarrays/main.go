package main

import (
	"fmt"
	"sort"
)

func MergeSortedArrays(ar1 []int, ar2 []int) []int {
	if len(ar1) == 0 {
		return ar2
	}

	if len(ar2) == 0 {
		return ar1
	}

	ar1 = append(ar1, ar2...)
	sort.Ints(ar1)
	return ar1
}

func MergeSortedArrays2(ar1 []int, ar2 []int) []int {
	if len(ar1) == 0 {
		return ar2
	}

	if len(ar2) == 0 {
		return ar1
	}
	ar3 := make([]int, 0)
	i, j := 0, 0

	for i < len(ar1) && j < len(ar2) {
		if ar1[i] < ar2[j] {
			//ar3[k] = ar1[i]
			ar3 = append(ar3, ar1[i])
			//k++
			i++
		} else {
			//ar3[k] = ar2[j]
			ar3 = append(ar3, ar2[j])
			//k++
			j++
		}
	}

	for i < len(ar1) {
		//ar3[k] = ar1[i]
		ar3 = append(ar3, ar1[i])
		//k++
		i++
	}

	for j < len(ar2) {
		//ar3[k] = ar2[j]
		ar3 = append(ar3, ar2[j])
		//k++
		j++
	}

	return ar3
}

func main() {
	ar1 := []int{0, 3, 4, 5, 31, 34}
	ar2 := []int{4, 6, 7, 30, 45}

	res := MergeSortedArrays(ar1, ar2)
	fmt.Println("After Merge: ", res)

	res = MergeSortedArrays2(ar1, ar2)
	fmt.Println("After Merge: ", res)
}

// For understanding
//ints := []int{4, 2, 7, 1, 9}
//sort.Sort(sort.Reverse(sort.IntSlice(ar1))) // Sort in descending order
//fmt.Println("Sorted in descending order:", ints) // Output: [9 7 4 2 1]
