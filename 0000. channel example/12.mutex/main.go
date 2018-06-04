package main

import "fmt"

func main() {
	mutex := make(chan struct{}, 1) // the capacity must be one

	counter := 0
	increse := func() {
		mutex <- struct{}{} //lock
		counter++
		<-mutex //unlock
	}

	increse1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increse()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increse1000(done)
	go increse1000(done)
	<-done
	<-done
	fmt.Println(counter) // 2000
}
