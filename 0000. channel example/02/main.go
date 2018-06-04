package main

import "fmt"

func main() {
	c := make(chan int, 2) // a buffered channel
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c))

	x, ok := <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))

	x, ok = <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))

	x, ok = <-c
	fmt.Println(x, ok) // 0 false

	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2

	close(c) // panic!
	c <- 7   // also panic if the last line is removed.
}
