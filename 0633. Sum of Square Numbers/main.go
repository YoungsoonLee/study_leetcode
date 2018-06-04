package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	a := intSqrt(c)
	fmt.Println(a)

	for a >= 0 {
		//fmt.Println(a, c-a*a)
		if isSquare(c - a*a) {
			return true
		}

		a--
	}

	return false
}

func intSqrt(c int) int {
	return int(math.Sqrt(float64(c)))
}

func isSquare(b2 int) bool {
	b := intSqrt(b2)
	return b*b == b2
}

func main() {
	a := 181
	fmt.Println(judgeSquareSum(a))
}
