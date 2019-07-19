package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, v := range nums {
		if _, exists := m[v]; exists {
			return true
		}
		m[v]++
	}

	return false
}

/*
func containsDuplicate(nums []int) bool {
	hm := make(map[int]int)

	for _, n := range nums {
		if _, ok := hm[n]; ok {
			hm[n]++
		} else {
			hm[n] = 1
		}
	}

	for _, v := range hm {
		if v >= 2 {
			return true
		}
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	hm := make(map[int]bool, len(nums))

	for _, n := range nums {
		if hm[n] {
			return true
		}
		hm[n] = true
	}

	return false
}
*/

func main() {
	a := []int{1, 2, 2, 3, 4, 5}
	fmt.Println(containsDuplicate(a))
}
