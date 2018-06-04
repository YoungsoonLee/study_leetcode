package main

import "fmt"

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num <= 6 {
		// 1,2,3,4,5,6 ugly
		return true
	}

	if num%2 == 0 {
		return isUgly(num / 2)
	}

	if num%3 == 0 {
		return isUgly(num / 3)
	}

	if num%5 == 0 {
		return isUgly(num / 5)
	}

	return false
}

func main() {
	a := 14
	fmt.Println(isUgly(a))
}
