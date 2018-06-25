package main

import "fmt"

func removeDuplicates(nums []int) int {

	i := 0

	for {
		if nums[i] == nums[i+1] {
			nums = append(nums, nums[i])
		} else {
			nums = append(nums, nums[i+1])
		}

		if nums[i] > nums[i+1] {
			nums = nums[i+1:]
			break
		}
		i++
	}

	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= nums[i+1] {
			nums = nums[:i+1]
			break
		}
	}
	fmt.Println(nums)
	return len(nums)

}

// !!!
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := i; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums)
	return i + 1
}

func main() {
	a := []int{1, 1, 2, 2}
	fmt.Println(removeDuplicates2(a))
}
