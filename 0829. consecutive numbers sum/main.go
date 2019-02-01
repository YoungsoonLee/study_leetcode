package main

import (
	"fmt"
	"math"
)

func consecutiveNumbersSum(N int) int {
	res := 0
	root := int(math.Sqrt(float64(2 * N))) // 제곱근 ...

	for k := 2; k <= root; k++ {
		kx := N - k*(k-1)/2
		fmt.Println("kx: ", kx)
		fmt.Println("kx%k : ", kx%k)
		if kx%k == 0 {
			res++
		}
	}

	fmt.Println("root: ", root)
	fmt.Println("res: ", res)

	return res
}

func main() {
	n := 15
	consecutiveNumbersSum(n)
}
