package main

import "fmt"

// brute-force
func singleNumber(nums []int) int {
	res := make(map[int]int, len(nums))
	for _, n := range nums {
		res[n]++
	}
	//fmt.Println(res)
	for i, v := range res {
		if v < 2 {
			return i
		}
	}

	return 0
}

func singleNumber2(nums []int) int {
	fmt.Println("test: ", 3^4)
	// n^n == 0
	// a^b^a^b^a == a
	res := 0
	for _, n := range nums {
		res ^= n
		fmt.Println(res)
	}
	return res
}

func main() {
	a := []int{2, 2, 3, 4, 4, 5, 5, 6, 6}
	//singleNumber(a)
	singleNumber2(a)
}
