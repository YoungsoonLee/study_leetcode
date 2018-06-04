package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	ret := make([]string, n)

	for i := range ret {
		x := i + 1
		switch {
		case x%15 == 0:
			ret[i] = "FizzBuzz"
		case x%3 == 0:
			ret[i] = "Fizz"
		case x%5 == 0:
			ret[i] = "Buzz"
		default:
			ret[i] = strconv.Itoa(x)
		}
	}
	fmt.Println(ret)
	return ret
}

func main() {
	n := 15
	fizzBuzz(n)
}
