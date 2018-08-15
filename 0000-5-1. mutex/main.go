package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for k := 0; k < 1000000; k++ {
			mutex.Lock()
			data = append(data, k)
			mutex.Unlock()
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			mutex.Lock()
			data = append(data, k)
			mutex.Unlock()
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			mutex.Lock()
			data = append(data, k)
			mutex.Unlock()
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			mutex.Lock()
			data = append(data, k)
			mutex.Unlock()
		}
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println(len(data))
}
