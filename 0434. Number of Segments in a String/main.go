package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	ret := strings.Fields(s)
	fmt.Println(ret, len(ret))

	return len(ret)
}

func countSegments2(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(s); i++ {
		if (i == 0 || string(s[i-1]) == " ") && string(s[i]) != " " {
			count++
		}
	}

	fmt.Println(count)
	return count
}

func main() {
	s := "Hello, my name is John"
	countSegments2(s)
}
