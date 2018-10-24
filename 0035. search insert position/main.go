package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := 0

	for low <= high {
		mid = (low + high) / 2
		fmt.Println(mid)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
			// expect
			if nums[low] < target && target < low+1 {
				return low + 1
			}
		} else {
			high = mid - 1
			if nums[high] < target && target < high+1 {
				return high + 1
			}
		}
	}

	return low
}

func searchInsert2(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] <= target {
		if nums[i] == target {
			return i
		}

		i++
	}

	return i
}

func main() {
	n := []int{1, 3, 5, 6}
	t := 7
	fmt.Println(searchInsert(n, t))
	//fmt.Println(searchInsert2(n, t))
}
