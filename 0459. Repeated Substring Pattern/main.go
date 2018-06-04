package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	size := len(s)

	fmt.Println(s + s)

	ss := (s + s)[1 : size*2-1]
	fmt.Println(ss)

	return strings.Contains(ss, s)
}

func main() {
	a := "abab"
	fmt.Println(repeatedSubstringPattern(a))
}
