package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}, 1) // the signal channel
	done <- struct{}{}             // fill the channel

	// now the channel done is full. A new send will block.
	go func() {
		fmt.Println("Hello")
		time.Sleep(time.Second * 2) // simulate a workload
		<-done
	}()

	// do something...
	done <- struct{}{} // block here until a receive is mode
	fmt.Println(" world!")

}
