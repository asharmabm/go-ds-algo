package main

import "fmt"

// ar := []int{5, 4, 1}
func bubbleSort(ar []int) []int { // O (n^2)
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			// j:=0;j<3-0-1;j++
			// if 5 > 4
			// ar[4,5,1]
			// j:= 1;j<2;j++
			// 5>1
			// ar[4,1,5]
			// i:=1
			// j:= 0;j<3-1-1;j++
			// 4>1
			// ar[1,4,5]
			// j:=1;j<1;j++
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}

	return ar
}

func main() {
	ar := []int{4, 5, 77, 8, 3, 1, 99, 105}
	res := bubbleSort(ar)

	fmt.Println("Result", res)
}
