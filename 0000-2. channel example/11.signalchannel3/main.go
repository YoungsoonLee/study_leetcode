package main

import (
	"log"
	"time"
)

func worker(id int, ready <-chan struct{}, done chan<- struct{}) {
	<-ready // wait until ready is closed
	log.Print("Worker #", id, " started to process.")
	time.Sleep(time.Second) // simulate a workload
	log.Print("Worker #", id, " finished its job.")
	done <- struct{}{} // notify the main goroutine. (N-to-1)
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan struct{}), make(chan struct{})
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	time.Sleep(time.Second * 2) // simulate an initialazation phase
	// 1-to-N notification
	ready <- struct{}{}
	ready <- struct{}{}
	ready <- struct{}{}
	// being N-to-1 notified
	<-done
	<-done
	<-done

}
