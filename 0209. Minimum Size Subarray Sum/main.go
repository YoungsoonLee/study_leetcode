package main

import (
	"fmt"
	"sort"
)

func minSubArrayLen(s int, nums []int) int {
	sort.Ints(nums)
	count := 0

	//var n int

	for i := len(nums) - 1; i > 0; i-- {
		s = s - nums[i]
		if s != 0 {
			count++
		} else {
			break
		}

	}

	return count + 1
}

func main() {
	s := 7
	n := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, n))
}
