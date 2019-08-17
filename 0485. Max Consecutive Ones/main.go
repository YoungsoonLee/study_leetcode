package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	max := 0

	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j = i
		} else {
			if max < i-j {
				max = i - j
			}
		}
	}
	fmt.Println(max)
	return max
}

/*
func findMaxConsecutiveOnes(nums []int) int {
	mc := 0
	i, j := 0, 1
	for j < len(nums) {
		if i == 0 && nums[i] == 1 {
			mc++
		} else if nums[i] == nums[j] {
			mc++
			j++
		} else {
			i = j + 1
			j = i + 1
		}
	}
	fmt.Println(mc)
	return mc
}
*/

func main() {
	a := []int{1, 1, 0, 1, 1, 1}
	findMaxConsecutiveOnes(a)
}
