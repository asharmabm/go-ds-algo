package main

import "fmt"

// Reverse - "Hi My name is Akshay"
func Reverse(str string) string {
	if str == "" || len(str) < 2 {
		return str
	}
	var res []rune
	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, rune(str[i]))
	}
	return string(res)
}

func Reverse2(str string) string {
	runes := []rune(str)
	for start, end := 0, len(runes)-1; start < end; start, end = start+1, end-1 {
		runes[start], runes[end] = runes[end], runes[start]
	}

	return string(runes)
}

func main() {
	res := Reverse("Hi My name is Akshay")
	fmt.Println("Reversed string", res)

	res = Reverse2("Hi My name is Akshay")
	fmt.Println("Reversed string", res)
}
