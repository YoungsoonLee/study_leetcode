package main

import (
	"fmt"
	"strings"
)

func rotateString(A string, B string) bool {
	//fmt.Println(A + A)
	//return false
	return len(A) == len(B) && strings.Contains(A+A, B)
}

func rotateString_my(A string, B string) bool {
	if len(A) == 0 && len(B) == 0 {
		return true
	}

	for _, c := range A {
		A = A[1:] + string(c)
		//fmt.Println(A)
		if A == B {
			return true
		}
	}

	return false
}

func main() {
	a := "abcde"
	b := "cdeab"

	fmt.Println(rotateString(a, b))
}
