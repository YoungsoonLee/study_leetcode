package main

import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	x ^= y
	fmt.Println(x, y)

	res := 0
	for x > 0 {
		res += x & 1
		fmt.Println("res: ", res)
		x >>= 1
		fmt.Println("x: ", x)
	}
	return res
}

func main() {
	a := 1
	b := 4
	fmt.Println("main: ", 5&1)
	hammingDistance(a, b)
}
