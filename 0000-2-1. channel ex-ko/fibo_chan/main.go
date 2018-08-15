package main

import "fmt"

func fib(n int) chan int {
	c := make(chan int)

	go func() { // run goroutine
		for i, j := 0, 1; i < n; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()

	return c
}

func main() {
	for i := range fib(1000) {
		fmt.Println("피보나치 수 :", i)
	}
}
