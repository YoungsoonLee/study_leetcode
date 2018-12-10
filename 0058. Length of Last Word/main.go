package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	word := strings.Fields(s)
	fmt.Println("word: ", word)

	n := len(word)
	if n == 0 {
		return n
	}

	if n == 1 {
		return len(word[0])
	} else {
		return len(word[n-1])
	}

	/*
		var tail int
		tail = len(s) - 1

		len := 0

		//fmt.Println("start tail: ", tail)

		for tail >= 0 && string(s[tail]) == " " {
			//fmt.Println("inner tail: ", tail)
			tail--
		}

		//fmt.Println("tail: ", tail)

		for tail >= 0 && string(s[tail]) != " " {
			len++
			tail--
		}

		//fmt.Println("len: ", len)
		//fmt.Println("last tail: ", tail)

		return len
	*/
}

func main() {
	a := "Hello World"
	fmt.Println(lengthOfLastWord(a))
}
