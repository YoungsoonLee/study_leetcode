package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		fmt.Println(n, n%3)
		n /= 3
	}

	return n == 1 // 3배수 맨끝이 3으로 나누어 져야 하니...
}

func main() {
	a := 27
	fmt.Println(isPowerOfThree(a))
}
