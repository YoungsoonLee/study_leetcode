package main

import "fmt"

func isPerfectSquare(num int) bool {
	r := intSqrt(num)
	fmt.Println("r: ", r)
	return r*r == num
}

func intSqrt(n int) int {
	r := n
	for r*r > n { // !!!
		fmt.Println(r)
		r = (r + n/r) / 2 // !!!
	}
	return r
}

func main() {
	a := 14
	fmt.Println(isPerfectSquare(a))
}
