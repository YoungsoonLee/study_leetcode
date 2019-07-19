package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if m[n] != 0 && i-(m[n]-1) <= k {
			fmt.Println(m)
			return true
		}

		m[n] = i + 1
	}
	fmt.Println(m)
	return false
}

/*
func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	m := make(map[int]int, len(nums))

	for i, n := range nums {
		fmt.Println(i, m[n], m)
		if m[n] != 0 && i-(m[n]-1) <= k { // !!!
			return true
		}

		m[n] = i + 1
	}

	return false
}
*/

func main() {
	a := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(a, k))
}
