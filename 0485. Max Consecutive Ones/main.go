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

func main() {
	a := []int{1, 0, 1, 1, 0, 1}
	findMaxConsecutiveOnes(a)
}
