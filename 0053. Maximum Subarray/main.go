package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	maxEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(maxEndingHere+nums[i], nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a))
}
