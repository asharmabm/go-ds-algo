package main

import "fmt"

func findFactorialRecursive(num int) int {
	if num == 1 {
		return 1
	}
	return num * findFactorialRecursive(num-1)
}

func findFactorialIterative(num int) int {
	var res = 1
	if num == 2 {
		res = 2
	}

	for i := 2; i <= num; i++ {
		res = res * i
	}
	return res
}

func main() {
	res := findFactorialRecursive(5)

	fmt.Printf("Factorial of num is, %d \n", res)
}
