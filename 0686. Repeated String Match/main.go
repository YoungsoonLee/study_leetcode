package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	//times := max(len(b)-len(a), 1)
	times := (len(b)-1)/len(a) + 1 // !!!!!

	fmt.Println(times)
	fmt.Println(strings.Repeat(a, times))

	if strings.Contains(strings.Repeat(a, times), b) {
		return times
	}

	if strings.Contains(strings.Repeat(a, times+1), b) { // 왜 딱 1번만 더할가...
		return times + 1
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := "baa"
	b := "abaab"
	fmt.Println(repeatedStringMatch(a, b))
}
