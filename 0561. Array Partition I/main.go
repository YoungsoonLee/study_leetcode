package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		sum += min(nums[i], nums[i+1])
	}
	fmt.Println(sum)
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	a := []int{1, 4, 3, 2}
	arrayPairSum(a)
}
