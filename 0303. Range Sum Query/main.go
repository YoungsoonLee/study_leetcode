package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(i, j int) int {
	res := 0
	for i <= j {
		res += this.nums[i]
		i++
	}

	return res
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	param_1 := obj.SumRange(2, 5)
	fmt.Println(param_1)
}
