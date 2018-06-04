package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(values)

	done := make(chan struct{})
	go func() { // the sorting goroutine
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // notify soritng is done
	}()

	// do some other things ...
	<-done // waiting here for notification
	fmt.Println(values[0], values[len(values)-1])
}
