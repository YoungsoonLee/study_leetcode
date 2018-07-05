package main

import "fmt"

func searchRange(nums []int, target int) []int {
	index := search(nums, target)
	fmt.Println(index)
	if index == -1 {
		return []int{-1, -1}
	}

	// 이분법을 사용하여 첫 번째 타겟을 찾습니다.
	first := index
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}
		first = f
	}

	// 利用二分法，查找最后一个target
	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}
		// 注意这里与查找first的不同
		last += l + 1
	}

	return []int{first, last}
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var median int

	for low <= high {
		median = (low + high) / 2

		switch {
		case nums[median] < target:
			low = median + 1
		case nums[median] > target:
			high = median - 1
		default:
			return median
		}
	}
	return -1
}

func main() {
	n := []int{2, 2}
	t := 2
	fmt.Println(searchRange(n, t))
}
