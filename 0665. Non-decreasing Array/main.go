package main

import (
	"fmt"
	"sort"
)

func checkPossibility(nums []int) bool {
	/*
		pre := deepCopy(nums)
		//pre := nums
		pre[0] = 100
		fmt.Println(nums)
		fmt.Println(pre)
	*/

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			pre := deepCopy(nums)
			// nums [i-1]을 nums [i]로 변경합니다.
			pre[i-1] = pre[i]

			next := deepCopy(nums)
			// nums [i]를 nums [i-1]로 변경합니다.
			next[i] = next[i-1]

			fmt.Println(sort.IntSlice(pre), sort.IntSlice(next))
			return sort.IsSorted(sort.IntSlice(pre)) || sort.IsSorted(sort.IntSlice(next))
		}
	}

	return true
}

func deepCopy(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	return res
}

func main() {
	a := []int{4, 2, 3}
	fmt.Println(checkPossibility(a))

	a = []int{4, 2, 1}
	fmt.Println(checkPossibility(a))
}
