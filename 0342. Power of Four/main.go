package main

import "fmt"

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}

	for num%4 == 0 {
		num /= 4
	}

	return num == 1
}

func main() {
	a := 5
	fmt.Println(isPowerOfFour(a))
}
