package main

import "fmt"

// Given an array = [2,5,1,2,3,5,1,2,4]
// It should return 2

// Given an array = [2,1,1,2,3,5,1,2,4]
// It should return 1

// Given an array = [2,3,4,5]
// It should return no recurring character

func GetFirstRecurringChar(ar []int) result {
	m := make(map[int]bool)
	for _, k := range ar {
		if _, ok := m[k]; ok {
			return result{k, ""}
		}
		m[k] = true
	}

	return result{
		recurChar: 0,
		resultStr: "no recurring character",
	}
}

type result struct {
	recurChar int
	resultStr string
}

func main() {
	ar := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	v := GetFirstRecurringChar(ar)
	if v.resultStr == "" {
		fmt.Println("First Recurring Char is: ", v.recurChar)
	}

	ar1 := []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	v1 := GetFirstRecurringChar(ar1)
	if v1.resultStr == "" {
		fmt.Println("First Recurring Char is: ", v1.recurChar)
	}

	ar2 := []int{2, 3, 4, 5}
	v2 := GetFirstRecurringChar(ar2)
	if v2.resultStr == "" {
		fmt.Println("First Recurring Char is: ", v2.recurChar)
	} else {
		fmt.Println(v2.resultStr)
	}
}
