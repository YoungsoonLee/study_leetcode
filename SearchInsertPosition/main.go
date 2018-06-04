package main

import "fmt"

// sorted nums
// use binary serarch
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := 0

	for low <= high {
		//fmt.Println(low, high, mid)
		mid = (low + high) / 2
		//fmt.Println("new mid: ", mid)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println(low, high, mid)
	return low
}

func main() {
	a := []int{1, 3, 5, 6}
	t := 7
	fmt.Println("result: ", searchInsert(a, t))
}
