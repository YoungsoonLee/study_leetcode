package main

import (
	"fmt"
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)

	m := make(map[int]bool, 0)

	for _, n := range nums {
		if n > 0 {
			m[n] = true
		}
	}

	fmt.Println("m: ", m)

	i := 1
	for {
		if exist := m[i]; !exist {
			return i
		}
		i++
	}

	// return i // not found
}

func main() {
	n := []int{1, 2, 3}
	fmt.Println(firstMissingPositive(n))
}
