package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	// fmt.Println(bytes)

	for i := 0; i < len(s); i += 2 * k {
		j := min(i+k, len(s))
		fmt.Println(i+k, j)
		reverse(bytes[i:j])
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func reverse(bytes []byte) {
	i, j := 0, len(bytes)-1

	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	s := "abcdefg"
	k := 2

	reverseStr(s, k)
}
