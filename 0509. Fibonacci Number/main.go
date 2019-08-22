package main

import "fmt"

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	return fib(N-2) + fib(N-1)
}

func main() {
	n := 4
	fmt.Println(fib(n))
}
