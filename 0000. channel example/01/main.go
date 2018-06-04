package main

import "fmt"

func main() {
	c := make(chan int) // an unbuffered channel
	go func() {
		x := <-c   // blocking here until a value is received.
		c <- x * x // blocking here until the result is received
	}()

	c <- 3   // blocking here until the value is received.
	y := <-c // blocking here until the result is received.
	fmt.Println(y)

}
