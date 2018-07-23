package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}, 1)
	done <- struct{}{}

	go func() {
		fmt.Println("Hello")
		time.Sleep(time.Second * 2)
		<-done
	}()

	done <- struct{}{}
	fmt.Println(" world!")
}
