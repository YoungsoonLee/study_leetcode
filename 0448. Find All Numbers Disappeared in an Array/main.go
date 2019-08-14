package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	// 1. I solve it with brute force
	// 2. and then I will modify this to get more perfermance.

	// s1. sort.
	// s2. use for loop, i = 0 is number 1 position.
	//     i count number and I add one into idx postion.
	// s3. end for loop,  number 0 is disapper posintion number
	n := len(nums)
	tmp := make([]int, n)

	for i := 0; i < n; i++ {
		tmp[nums[i]-1]++
	}
	fmt.Println(tmp)
	ret := []int{}

	for i := 0; i < n; i++ {
		if tmp[i] == 0 {
			ret = append(ret, i+1)
		}
	}
	fmt.Println(ret)

	// S.C is O(n)
	// T.C is O(n)
	return ret

}

func main() {

}
