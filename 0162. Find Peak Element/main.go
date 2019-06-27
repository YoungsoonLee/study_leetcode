package main

import "fmt"

// return value is index of peak nums
func findPeakElement(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n - 1
	}

	if n == 2 {
		if nums[0] < nums[1] {
			return 0
		}
		return 1
	}

	if n == 3 {
		if nums[0] < nums[1] && nums[1] > nums[1] {
			return 1
		}
	}

	peakIndex := make([]int, 0)

	for i := 1; i < n-1; i++ {
		if nums[i] == nums[n-1] {
			break
		} else if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			peakIndex = append(peakIndex, i)
		}
	}

	fmt.Println(peakIndex)
	return peakIndex[0]
}

func main() {
	n := []int{1, 2, 1, 3, 5, 6, 4}
	findPeakElement(n)
}
