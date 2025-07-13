package main

import "fmt"

func addNumbers(nums []int) []int {
	fmt.Printf("Address of nums in addNumbers1: %p %d \n ", &nums, cap(nums))
	nums = append(nums, 4)
	fmt.Printf("Address of nums in addNumbers2: %p, %d \n", &nums, cap(nums))

	nums = append(nums, 5)

	fmt.Println("Nums", nums)

	return nums
}

func main() {
	var numbers []int
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)

	fmt.Printf("Address of nums in addNumbers0: %p %d \n ", &numbers, cap(numbers))

	n := addNumbers(numbers)

	fmt.Printf("Address of nums in addNumbers3: %p %d \n ", &n, cap(n))

	fmt.Println("Updated slice:", numbers)
}
