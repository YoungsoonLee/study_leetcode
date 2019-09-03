package main

import (
	"sort"
)

/*
func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		sum += min(nums[i], nums[i+1])
	}
	fmt.Println(sum)
	return sum
}
*/

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums) / 2
	min := nums[0]
	for i := 0; i < n; i++ {
		if i != 0 {
			min += nums[i+2]
		}
	}
	return min
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
