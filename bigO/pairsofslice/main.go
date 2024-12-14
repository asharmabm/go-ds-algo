package main

import "fmt"

// Log all pairs of slice
func pairsOfSlice(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			fmt.Printf("%d, %d\n", s[i], s[j])
		}
	}
}

// O(n * n) = O(n^2)  - Quadratic time
// ^ means square

func main() {
	pairsOfSlice([]int{1, 2, 3, 4, 5})
}
