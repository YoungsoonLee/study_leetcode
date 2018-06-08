package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// Avoid adding duplicate results
		// i>0 is to prevent nums[i-1] overflow
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// The smaller l needs to be bigger
				l++
			case s > 0:
				// Larger r needs to be smaller
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// To avoid duplicate additions, both l and r need to be moved to different elements.
				l, r = next(nums, l, r)
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}

// threeSum_my ...
// not acccepted. because overtime.
func threeSum_my(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)

	var ret [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // !!!!!
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				fmt.Println(nums[i], nums[j], nums[k])
				if nums[i]+nums[j]+nums[k] == 0 {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	// fine duplicate.

	fmt.Println(ret)
	return ret
}

func main() {
	n := []int{-1, 0, 1, 2, -1, -4}
	threeSum_my(n)
}
