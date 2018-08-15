package main

import (
	"fmt"
	"time"
)

func main() {
	s := make(chan int)

	// Sender (값을 보내는 녀석)
	s <- 3

	go func() {
		// Receiver (값을 받는 녀석)
		f := <-s
		fmt.Println(f)
	}()

	fmt.Println("보고싶다 이 로그")
	time.Sleep(3 * time.Second)
}
