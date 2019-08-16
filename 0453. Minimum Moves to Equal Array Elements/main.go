package main

import (
	"fmt"
)

func minMoves(nums []int) int {
	sum, min := 0, nums[0]
	for _, n := range nums {
		sum += n
		fmt.Println("sum: ", sum, min, n)
		if min > n {
			min = n
		}
	}

	fmt.Println("result: ", sum, min, len(nums))
	return sum - min*len(nums)
}

func main() {
	a := []int{1, 2, 3}
	minMoves(a)
}
