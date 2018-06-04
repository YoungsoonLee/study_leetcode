package main

import "fmt"

// my solution - failed
/*
func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	for n >= 0 {
		if n == 0 {
			return true
		} else if n == 1 {
			return false
		}
		n = n % 2
	}

	return false
}
*/

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n /= 2
	}

	return true
}

func main() {
	a := 16
	fmt.Println(isPowerOfTwo(a))
}
