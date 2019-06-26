package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	res := ""
	temp := ""
	scount := 0

	for _, c := range s {
		fmt.Println("c: ", string(c))
		if string(c) != " " {
			temp += string(c)
			scount = 0
		} else {
			if scount >= 1 {
				continue
			} else {
				scount++
				res = string(c) + temp + res
				temp = ""
			}
		}
	}

	res = temp + res

	fmt.Println("[" + res + "]")
	return strings.TrimRight(strings.TrimLeft(res, " "), " ")
}

func main() {
	s := "a good   example"
	reverseWords(s)
}
