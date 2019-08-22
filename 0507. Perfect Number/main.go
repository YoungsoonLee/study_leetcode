package main

import (
	"fmt"
	"math"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	root := int(math.Sqrt(float64(num)))

	fmt.Println("root: ", root)

	for i := 2; i <= root; i++ {
		fmt.Println(i, num%i, num/i, sum, i+(num/i))
		if num%i == 0 {
			sum += i + (num / i)
		}
	}

	return sum == num
}

func main() {
	a := 28
	fmt.Println(checkPerfectNumber(a))
}
