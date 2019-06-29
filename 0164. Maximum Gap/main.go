package main

import (
	"fmt"
	"sort"
)

func maximumGap(nums []int) int {
	n := len(nums)
	fmt.Println(n)
	if n <= 1 {
		return 0
	}

	sort.Ints(nums)
	m := 0

	for i := 1; i <= n-1; i++ {
		m = max(nums[i]-nums[i-1], m)
	}

	fmt.Println(m)
	return m
}

func max(a, b int) int {
	fmt.Println(a, b)
	if a > b {
		return a
	}
	return b
}

func main() {
	a := []int{1, 10000000}
	maximumGap(a)
}
