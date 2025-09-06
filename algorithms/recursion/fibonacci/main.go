package main

import "fmt"

// Given a number N return the index value of the Fibonacci sequence, where the sequence is:

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 ...
// the pattern of the sequence is that each value is the sum of the 2 previous values,
// that means that for N=5 → 2 + 3

//For example: fibonacciRecursive(6) should return 8

func fibonacciIterative(num int) int {
	if num < 2 {
		return num
	}
	last := 1
	beforeLast := 0
	next := 0
	for i := 0; i <= num-2; i++ {
		next = last + beforeLast
		beforeLast = last
		last = next
	}

	return next
}

func fibonacciIterativeSlice(num int) int {
	slice := []int{0, 1}
	for i := 2; i <= num; i++ {
		slice = append(slice, slice[i-2]+slice[i-1])
	}
	return slice[num]
}

func fibonacciRecursive(num int) int { // O(2^n)
	if num < 2 {
		return num
	}

	return fibonacciRecursive(num-1) + fibonacciRecursive(num-2)
}

func main() {
	res := fibonacciIterative(43)

	fmt.Println("Answer", res)

	res = fibonacciRecursive(43)

	fmt.Println("Answer", res)

	//res = fibonacciIterativeSlice(3)
	//
	//fmt.Println("Answer", res)
}
