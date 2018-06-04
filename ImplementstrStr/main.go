package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {

	return strings.Index(haystack, needle)
}

func strStr2(haystack string, needle string) int {
	// brute-force
	ns := len(needle)
	for i := 0; i < len(haystack)-ns; i += ns {
		//fmt.Println(haystack[i:i+2], needle, i)
		if haystack[i:i+ns] == needle {
			return i
		}
	}
	return -1
}

func main() {
	h := "hello"
	n := "ll"
	//strStr(h, n)
	fmt.Println(strStr2(h, n))
}
