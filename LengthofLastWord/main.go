package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	var tail int
	tail = len(s) - 1

	len := 0

	fmt.Println("start tail: ", tail)

	for tail >= 0 && string(s[tail]) == " " {
		//fmt.Println("inner tail: ", tail)
		tail--
	}

	fmt.Println("tail: ", tail)

	for tail >= 0 && string(s[tail]) != " " {
		len++
		tail--
	}

	fmt.Println("len: ", len)
	fmt.Println("last tail: ", tail)

	return len
}

func main() {
	s := "       "
	fmt.Println(lengthOfLastWord(s))
}
