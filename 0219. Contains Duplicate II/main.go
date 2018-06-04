package main

import "fmt"

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

func main() {
	a := []int{1, 2, 1, 3, 4, 5}
	k := 2
	fmt.Println(containsNearbyDuplicate(a, k))
}
