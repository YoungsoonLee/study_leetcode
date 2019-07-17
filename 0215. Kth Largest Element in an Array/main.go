package main

import (
	"fmt"
	"sort"

	"github.com/lithammer/shortuuid"
)

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[k-1]
}

func main() {
	m := make(map[string]int)
	for i := 0; i < 1000000; i++ {
		u := shortuuid.New()
		if _, ok := m[u]; ok {
			m[u]++
		} else {
			m[u] = 1
		}
	}

	for k, v := range m {
		if v >= 2 {
			fmt.Println("dup: ", k, v)
		}
	}

	//fmt.Println(m)
}
