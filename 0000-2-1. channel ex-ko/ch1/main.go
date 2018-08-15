package main

import (
	"fmt"
	"time"
)

func main() {
	input := make(chan int)

	go func() {
		for data := range input {
			fmt.Println("고루틴에 들어온 녀셕을 처리 했다 : ", data)
		}

		fmt.Println("여기는 언제 찍힐까?")
	}()

	input <- 5
	input <- 10
	close(input) // get log and panic
	input <- 12

	time.Sleep(1 * time.Second)

}
