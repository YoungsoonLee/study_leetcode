package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = countIdx(s)
	}
	return s
}

func countIdx(st string) string {
	sb := make([]string, 0)
	c := string(st[0])
	count := 1
	for i := 1; i < len(st); i++ {
		if string(st[i]) == c {
			count++
		} else {
			sb = append(sb, strconv.Itoa(count))
			sb = append(sb, c)
			c = string(st[i])
			count = 1
		}
	}

	sb = append(sb, strconv.Itoa(count))
	sb = append(sb, c)

	//fmt.Println(sb)

	return strings.Join(sb, "")
}

func main() {
	n := 3
	fmt.Println(countAndSay(n))
}
