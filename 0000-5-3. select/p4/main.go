package main

import (
	"fmt"
	"time"
)

func consuming(scheduler chan string) {
	select {
	case <-scheduler:
		fmt.Println("이름을 입력 받았습니다.")
	case <-time.After(5 * time.Second):
		fmt.Println("시간이 지났습니다.")
	}
}

func producing(scheduler chan string) {
	var name string
	fmt.Print("이름: ")
	fmt.Scanln(&name)
	scheduler <- name
}

func main() {
	scheduler := make(chan string)

	go consuming(scheduler)
	go producing(scheduler)

	time.Sleep(10 * time.Second)
}
