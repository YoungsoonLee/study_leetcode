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

func main() {
	n1 := []int{4, 1, 2}
	n2 := []int{1, 3, 4, 2}

	nextGreaterElement(n1, n2)
}
