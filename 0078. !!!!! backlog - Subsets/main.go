package main

import (
	"fmt"
	"math"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0, resSize(nums))
	rec(nums, []int{}, &res)

	fmt.Println(res)
	return res
}

func resSize(nums []int) int {
	return int(math.Pow(2, float64(len(nums))))
}

func rec(nums, temp []int, res *[][]int) {
	size := len(nums)
	if size == 0 {
		*res = append(*res, temp)
		return
	}

	// 对于 last 来说，要么存在于某个子集中，要么不存在
	// 마지막으로, 하위 집합에 있거나 존재하지 않습니다.
	nums, last := nums[:size-1], nums[size-1]

	fmt.Println("first rec: ", nums, temp, res)

	rec(nums, temp, res)

	withLast := make([]int, 1, 1+len(temp))
	withLast[0] = last
	withLast = append(withLast, temp...)

	fmt.Println("second rec: ", nums, temp, res)

	rec(nums, withLast, res)
}

/*
func subsets(nums []int) [][]int {
    res := [][]int{}
    backtrack([]int{}, nil, &res)
    for i, num := range nums {
        backtrack([]int{num}, nums[i+1:], &res)
    }
    return res
}

func backtrack(subset []int, nums []int, res *[][]int) {
    if len(nums) == 0 {
        *res = append(*res, subset)
        return
    }

    backtrack(subset, nil, res)
    for i, num := range nums {
        // Go-specific, just creating a new array and prepend a num
        newSubset := make([]int, len(subset)+1)
        copy(newSubset[0:len(subset)], subset)
        newSubset[len(subset)] = num

        backtrack(newSubset, nums[i+1:], res)
    }
}
*/

func main() {
	n := []int{1, 2, 3}
	subsets(n)
}
