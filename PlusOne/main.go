package main

import "fmt"

func plusOne(digits []int) []int {
	fmt.Println(digits)
	size := len(digits)
	for i := size - 1; i >= 0; i-- {
		fmt.Println(digits[i], i)
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}

	//
	if digits[0] == 0 {
		nn := make([]int, size+1, size+1)
		nn[0] = 1
		return nn
	}

	return digits
}

func main() {
	a := []int{3, 5, 9}
	fmt.Println(plusOne(a))
}
