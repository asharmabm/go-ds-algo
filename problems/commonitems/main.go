package main

import "fmt"

func hasCommonItems(ar1 []string, ar2 []string) bool {
	m := make(map[string]int)

	less := len(ar1)
	lessAr := ar1
	high := len(ar1)
	highAr := ar1

	if len(ar2) < len(ar1) {
		less = len(ar2)
		lessAr = ar2
	} else {
		high = len(ar2)
		highAr = ar2
	}

	for i := 0; i < less; i++ {
		m[lessAr[i]] = 0
		if _, ok := m[highAr[i]]; ok {
			return true
		}
	}

	for i := less; i < high; i++ {
		if _, ok := m[highAr[i]]; ok {
			return true
		}
	}
	return false
}

func main() {

	res := hasCommonItems([]string{"a", "b", "c", "x", "i"}, []string{"z", "y", "i"})
	fmt.Println(res)
	res = hasCommonItems([]string{"a", "b", "c", "a"}, []string{"z", "y", "a"})
	fmt.Println(res)

}
