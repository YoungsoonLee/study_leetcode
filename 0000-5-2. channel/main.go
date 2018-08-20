package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data = []int{}

	var c = make(chan int)

	go func() {
		for k := 0; k < 1000000; k++ {
			<-c
			data = append(data, k)
			c <- 1
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			<-c
			data = append(data, k)
			c <- 1
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			<-c
			data = append(data, k)
			c <- 1
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			<-c
			data = append(data, k)
			c <- 1
		}
	}()

	c <- 1 // start

	time.Sleep(500 * time.Millisecond)
	fmt.Println(len(data))

}
