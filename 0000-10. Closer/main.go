package main

import "fmt"

func addTwo() func() int {
	sum := 0
	return func() int {
		sum += 2
		return sum
	}
}

func addTwoDynamic() func() int {
	return func() int {
		sum += 2
		return sum
	}
}

func main() {
	twoMore := addTwo()
	fmt.Println(twoMore())
	fmt.Println(twoMore())
}
