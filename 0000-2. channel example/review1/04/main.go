package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(1 << 63)
	f := fibonacci()
	for x, ok := <-f; ok; x, ok = <-f {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}

func fibonacci() chan uint64 {
	c := make(chan uint64)
	go func() {
		var x, y uint64 = 0, 1
		for ; y < (1 << 63); c <- y {
			x, y = y, x+y
		}
		close(c)
	}()

	return c
}
