package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	maxEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = int(math.Max(float64(maxEndingHere+nums[i]), float64(nums[i])))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}
	return maxSoFar
}

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a))
}
