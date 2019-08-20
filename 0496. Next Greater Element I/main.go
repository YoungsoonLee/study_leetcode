package main

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	indexOf := make(map[int]int)
	for i, n := range nums {
		indexOf[n] = i
	}

	fmt.Println(indexOf)

	res := make([]int, len(findNums))

	for i, n := range findNums {
		res[i] = -1
		for j := indexOf[n] + 1; j < len(nums); j++ {
			fmt.Println(n, j, nums[j], i)
			if n < nums[j] {
				res[i] = nums[j]
				break
			}
		}
	}

	fmt.Println(res)
	return res
}

/*
func nextGreaterElement(findNums []int, nums []int) []int {
	// 1. use map
	m := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		// without dup. so i cannot check dup in map
		if nums[i] < nums[i+1] {
			m[nums[i]] = nums[i+1]
		} else {
			m[nums[i]] = -1
		}
	}

	// last one
	m[nums[len(nums)-1]] = -1

	fmt.Println(m)

	res := make([]int, len(findNums), len(findNums))
	for i, f := range findNums {
		if _, ok := m[f]; ok {
			res[i] = m[f]
		}
	}

	fmt.Println(res)
	return res
}
*/

func main() {
	n1 := []int{1, 3, 5, 2, 4}
	n2 := []int{6, 5, 4, 3, 2, 1, 7}

	nextGreaterElement(n1, n2)
}
