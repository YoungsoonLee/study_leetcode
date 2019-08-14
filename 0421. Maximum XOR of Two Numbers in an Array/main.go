package main

import (
	"fmt"
)

func findMaximumXOR(nums []int) int {
	var max, mask int

	for i := 31; i >= 0; i-- {
		mask |= 1 << uint(i)

		nMap := make(map[int]bool)
		for _, num := range nums {

		}
	}

	fmt.Println(mask)

	return max
}

func main() {
	n := []int{3, 10, 5, 25, 2, 8}
	findMaximumXOR(n)
}
