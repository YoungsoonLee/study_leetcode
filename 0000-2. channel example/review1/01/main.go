package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		x := <-c
		c <- x * x
	}()

	c <- 3
	y := <-c
	fmt.Println(y)
}
