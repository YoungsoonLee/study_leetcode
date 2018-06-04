package main

import (
	"fmt"
	"math"
)

/*
func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)

	if len(nums) < 3 {
		return nums[0]
	}
	return nums[2]
}
*/
func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	fmt.Println(max1, max2, max3)
	for _, n := range nums {
		if n == max1 || n == max2 {
			continue
		}

		switch {
		case max1 < n:
			max3, max2, max1 = max2, max1, n
		case max2 < n:
			max3, max2 = max2, n
		case max3 < n:
			max3 = n
		}
	}

	// 세 번째로 큰 숫자가 없다.
	if max3 == math.MinInt64 {
		fmt.Println("2: ", max2)
		return max1
	}

	return max3
}

func main() {
	a := []int{2, 1}
	fmt.Println(thirdMax(a))
}
