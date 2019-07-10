package main

import "fmt"

func trailingZeroes(n int) int {
	/*
		if n == 0 {
			return 1
		}
		return n * trailingZeroes(n-1)
	*/

	res := 0
	for n >= 5 {
		n /= 5
		fmt.Println(n)
		res += n
	}

	return res

}

func main() {
	n := 12

	fmt.Println(trailingZeroes(n))
}
