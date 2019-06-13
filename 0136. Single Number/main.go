package main

import (
	"fmt"
	"sort"
)

// ^ operator
// the binary XOR operator. It copies the bit if it is set in only one operand and not both operators.
func singleNumber2(nums []int) int {
	res := 0
	for _, n := range nums {
		// n^n == 0
		// a^b^a^b^a == a
		fmt.Println(res, n)
		res ^= n
	}
	return res
}

func singleNumber(nums []int) int {
	sort.Ints(nums)

	if len(nums) == 1 {
		return nums[0]
	}

	i := 0
	for i < len(nums)-2 {

		if nums[i] != nums[i+1] {
			return nums[i]
		}

		i = i + 2
	}

	if i == len(nums)-1 {
		return nums[i]
	}

	return 0
}

func main() {
	n := []int{4, 1, 2, 1, 2}
	fmt.Println(4 ^ 1)
	fmt.Println(singleNumber2(n))
}
