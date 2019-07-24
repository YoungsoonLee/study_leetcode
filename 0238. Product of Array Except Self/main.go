package main

import "fmt"

func productExceptSelf(nums []int) []int {
	size := len(nums)

	res := make([]int, size)

	left := 1
	for i := 0; i < size; i++ {
		res[i] = left
		left *= nums[i]
	}

	fmt.Println("left: ", res)

	right := 1
	for i := size - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	fmt.Println("right: ", res)
	return res
}

func main() {
	a := []int{1, 2, 3, 4}
	productExceptSelf(a)
}
