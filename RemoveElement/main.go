package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	//fmt.Println(nums, i)
	nums = nums[i:]
	fmt.Println(nums)
	return len(nums)
}

func removeElement2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := i; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	fmt.Println(nums)
	return i
}

func main() {
	a := []int{3, 2, 2, 3}
	fmt.Println(removeElement2(a, 3))
}
