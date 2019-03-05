package main

import "fmt"

func removeDuplicates(nums []int) int {
	size := len(nums)

	i := 2

	for j := i; j < size; j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	fmt.Println(nums, i)
	return i
}

func main() {
	n := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(n)
}
