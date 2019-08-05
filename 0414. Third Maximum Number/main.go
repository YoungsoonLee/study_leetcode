package main

import (
	"fmt"
	"math"
)

/*
func thirdMax(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	m := make([]int, 0)

	max := nums[0]
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if len(m) < 3 {
			m = append(m, nums[i])
		}

		if nums[i] > max {
			max = nums[i]
			if len(m) >= 3 {
				sort.Ints(m)
				m[0] = nums[i]
			}
		} else if nums[i] < min {
			min = nums[i]
		} else {
			continue
		}

	}

	sort.Ints(m)
	fmt.Println(m)

	if len(m) < 3 {
		return m[1]
	}
	return m[0]
}
*/

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
	a := []int{2, 2, 3, 1}
	fmt.Println(thirdMax(a))
}
