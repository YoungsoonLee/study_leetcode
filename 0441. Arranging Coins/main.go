package main

import "fmt"

func arrangeCoins(n int) int {
	n *= 2
	//fmt.Println(n)
	x := n
	for !(x*(x+1) <= n && n < (x+1)*(x+2)) {
		fmt.Println((x * (x + 1)), (x+1)*(x+2))
		x = (x + n/(x+1)) / 2
	}
	fmt.Println(x)
	return x
}

func main() {
	a := 8
	arrangeCoins(a)
}
