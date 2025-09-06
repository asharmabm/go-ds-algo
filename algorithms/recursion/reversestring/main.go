package main

import "fmt"

// a, b, c, d, e
// e, d, c, b, a

func reverseString(s string) string { // O(n2)
	if len(s) == 1 {
		return s
	}
	newStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		newStr = newStr + string(s[i])
	}

	return newStr
}

func reverseStringBytes(s string) string { // O(n)
	n := len(s)
	b := make([]byte, n) // allocate once
	for i := 0; i < n; i++ {
		b[i] = s[n-1-i]
	}
	return string(b)
}

func reverseStringUnicode(s string) string {
	r := []rune(s) // convert to runes
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
}

func reverseStringRecursive(s string) string {
	r := []rune(s)
	return string(reverseRunes(r))
}

func reverseRunes(r []rune) []rune {
	n := len(r)
	if n <= 1 {
		return r
	}
	// put last rune at front, recurse in the middle, put first rune at end
	return append([]rune{r[n-1]}, append(reverseRunes(r[1:n-1]), r[0])...)
}

func main() {
	res := reverseString("abcde")
	fmt.Println("Result reversed", res)
}
