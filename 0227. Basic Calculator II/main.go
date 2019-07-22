package main

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	c := 0

	var dfs func(string)
	dfs = func(s string) {

		if len(s) < 3 {
			return
		}

		for i := 1; i < len(s); i++ {
			if s[i] == ' ' {
				s = s[:i-1]
				break
			}

			a, _ := strconv.Atoi(string(s[i-1]))
			b, _ := strconv.Atoi(string(s[i+1]))

			if s[i] == '*' {
				c += a * b
			}
			if s[i] == '/' {
				c += a / b
			}
			if s[i] == '+' {
				c += a + b
			}
			if s[i] == '-' {
				c += a - b
			}

			s = s[:i-1]
			fmt.Println("s: ", s)
			break
		}

		dfs(s)
	}

	dfs(s)

	fmt.Println(s)

	for i := 1; i < len(s); i++ {

		a, _ := strconv.Atoi(string(s[i-1]))
		//b, _ := strconv.Atoi(string(s[i+1]))

		if s[i] == '*' {
			c += a * c
		}
		if s[i] == '/' {
			c += a / c
		}
		if s[i] == '+' {
			c += a + c
		}
		if s[i] == '-' {
			c += a - c
		}
		s = s[:i-1]
		break
	}
	fmt.Println(c)
	return c
}

func main() {
	a := "3+2*2"
	calculate(a)
}
