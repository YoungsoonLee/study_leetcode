package main

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	fmt.Println(sum)

	left := 0
	for i := range nums {
		if left*2+nums[i] == sum {
			return i
		}
		left += nums[i]
	}

	return -1
}

func main() {
	a := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(a))
}
