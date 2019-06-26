package main

import "fmt"

func maxProduct(nums []int) int {
	cur, neg, max := 1, 1, nums[0]
	for i := 0; i < len(nums); i++ {
		fmt.Println(cur, neg, max)
		switch {
		case nums[i] > 0:
			cur, neg = nums[i]*cur, nums[i]*neg
		case nums[i] < 0:
			cur, neg = nums[i]*neg, nums[i]*cur
		default:
			cur, neg = 0, 1
		}

		if max < cur {
			max = cur
		}

		if cur <= 0 {
			cur = 1
		}
	}

	return max
}

/*
func maxProduct(nums []int) int {
	n := len(nums) - 1
	m := nums[0]

	for i := 0; i < n; i++ {
		c, n := nums[i], nums[i+1]
		if checkSubArray(c, n) {
			//sfmt.Println("true")
			m = max((c * n), m)
		}
	}

	fmt.Println(m)
	return m
}

func checkSubArray(a, b int) bool {
	if abs(abs(a)-abs(b)) == 1 {
		// subarray
		return true
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
*/

func main() {
	a := []int{2, 3, -2, 4}
	maxProduct(a)
}
