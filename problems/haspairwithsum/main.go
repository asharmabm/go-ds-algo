package main

import "fmt"

func hasPairWithSum(ar []int, sum int) bool {
	m := make(map[int]int)
	for i := 0; i < len(ar); i++ {
		if _, ok := m[sum-ar[i]]; ok {
			return true
		} else {
			m[ar[i]] = 0
		}
	}
	return false
}

// O(n)  - Time complexity
// O(n) - Space complexity

func main() {
	res := hasPairWithSum([]int{1, 2, 3, 4, 5}, 8)
	fmt.Println(res)
}
