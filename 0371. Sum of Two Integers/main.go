package main

import "fmt"

func getSum(a int, b int) int {

	for a != 0 {
		fmt.Println(a & b)
		fmt.Println((a & b) << 1)
		fmt.Println(a ^ b)
		a, b = (a&b)<<1, a^b
	}
	return b
}

func main() {
	a := 10
	b := 2
	fmt.Println(getSum(a, b))
}
