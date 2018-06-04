package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := 0, -1 // 왼쪽과 오른쪽의 값은 right-left + 1 == 0을 만족하는 한.
	min, max := nums[n-1], nums[0]

	for i := 1; i < n; i++ {
		if max <= nums[i] { // nums가 증가하면 if 문은 항상 참이어야합니다.
			max = nums[i]
		} else {
			// 설명 nums [i]가 앞에있는 숫자보다 작 으면 잘못 배치되어 재정렬해야합니다.
			right = i
		}

		j := n - 1 - i
		if min >= nums[i] {
			min = nums[i]
		} else {
			left = j
		}
	}

	return right - left + 1
}

func findUnsortedSubarray_my(nums []int) int {
	snums := make([]int, len(nums))
	copy(snums, nums)
	sort.Ints(snums)

	fmt.Println(nums)
	fmt.Println(snums)

	start, end := len(snums), 0

	for i := 0; i < len(snums); i++ {
		if snums[i] != nums[i] {
			start = min(start, i)
			end = max(end, i)
		}

		fmt.Println(start, end)
	}

	if end-start >= 0 {
		return end - start + 1
	} else {
		return 0
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray_my(a))
}
