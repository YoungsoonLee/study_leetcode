package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 1

	i, j := 0, 1
	for j < len(nums) {
		//fmt.Println(j, i, res)
		for j < len(nums) && nums[j-1] < nums[j] {
			fmt.Println(j, i, res)
			j++
		}
		fmt.Println(j, i, res)
		// !!!!!
		if res < j-i {
			res = j - i
		}

		i = j
		j++
	}

	return res
}

func main() {
	a := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(a))

	//a := []int{2, 2, 2, 2, 2}
	//fmt.Println(findLengthOfLCIS(a))
}
