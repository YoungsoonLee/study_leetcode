package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / pow(x, -n)
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {
	fmt.Println("n: ", n, "n>>1: ", n>>1)

	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	res := pow(x, n>>1)

	fmt.Println("res:", res, " n:", n, " n&1:", n&1)

	if n&1 == 0 {
		return res * res
	}

	return res * res * x
}

func main() {
	x := 2.0
	n := 5
	fmt.Println(myPow(x, n))
}
