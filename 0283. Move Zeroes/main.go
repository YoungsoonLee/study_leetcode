package main

import "fmt"

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums, nums[i])
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	// last 0 check
	if nums[0] == 0 {
		nums = append(nums, nums[0])
		nums = nums[1:]
	}
	fmt.Println(nums)
}

func moveZeroes2(nums []int) {
	l := len(nums)
	i, j := 0, 0

	for j < l {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	fmt.Println(i, nums)
	for i < l {
		nums[i] = 0
		i++
	}
	fmt.Println(nums)
}

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes2(a)
}
