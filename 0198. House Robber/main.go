package main

import "fmt"

func rob(nums []int) int {
	// a is even max
	// b is odd max
	var a, b int
	for i, v := range nums {
		fmt.Println(a, b, v)
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, b+v)
		}
	}

	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	a := []int{2, 1, 4, 3}
	fmt.Println(rob(a))
}
