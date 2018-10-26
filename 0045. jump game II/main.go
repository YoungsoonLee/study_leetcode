package main

import (
	"fmt"
)

func jump(nums []int) int {
	i, count, end := 0, 0, len(nums)-1

	var nextI, maxNextI, maxI int

	for i < end {
		if i+nums[i] >= end {
			return count + 1
		}

		nextI, maxNextI = i+1, i+nums[i]
		for nextI <= maxNextI {
			if nextI+nums[nextI] > maxI {
				maxI, i = nextI+nums[nextI], nextI
			}
			nextI++
		}

		count++
	}
	return count
}

/*
func jump(nums []int) int {
	res := 0
	count := 0

	for i := 0; i < len(nums); res++ {

		if res >= len(nums) {
			return count
		}

		f := nums[i]
		if f == 0 {
			return 0
		}

		if f >= len(nums) {
			return count
		}

		for j := i + 1; j < f; j++ {
			res = max(nums[j], res)
		}

		count++
	}

	return count

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

func main() {
	a := []int{9, 7, 7, 4, 7, 3}
	fmt.Println(jump(a))
}
