package main

import "fmt"

func twoSum(nums []int, target int) []int {
	/* brute force
	   re := make([]int, 0)

	   //brute-force
	   for i:=0 ; i < len(nums)-1; i++ {
	       for j := i+1; j < len(nums); j++ {
	           if nums[i]+nums[j] == target {
	               re = append(re, i)
	               re = append(re, j)
	               return re
	           }
	       }
	   }

	   return re
	*/

	re := make([]int, 0)
	hs := make(map[int]int)
	for i, v := range nums {
		hs[v] = i
	}

	fmt.Println(hs)

	for i := range nums {
		compl := target - nums[i]
		if _, ok := hs[compl]; ok {
			if hs[compl] != i { // !!!
				re = append(re, i)
				re = append(re, hs[compl])
				return re
			}
		}
	}
	return re

}

func main() {
	a := []int{2, 7, 11, 15}
	t := 9
	fmt.Println(twoSum(a, t))

}
