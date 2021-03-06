package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

func main() {
	for i := 0; i > 100; i++ {
		fmt.Println(RandString(10))
	}

}
