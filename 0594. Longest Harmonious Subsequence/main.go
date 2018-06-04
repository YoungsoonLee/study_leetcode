package main

import "fmt"

func findLHS(nums []int) int {
	r := make(map[int]int, len(nums))
	for _, n := range nums {
		r[n]++
	}
	fmt.Println(r)

	max := 0
	for n, c1 := range r {
		c2, ok := r[n+1]
		fmt.Println(n, c1, c2, ok)

		if ok {
			t := c1 + c2
			if max < t {
				max = t
			}
		}
	}

	return max
}

func findLHS_my(nums []int) int {
	rr := make(map[int]int)
	rec := make(map[int][]int)

	for _, v := range nums {
		rr[v]++
	}

	fmt.Println(rr)

	i := 0
	for k, v := range rr {
		rec[i] = []int{k, v}
		i++
	}

	fmt.Println(rec)
	count := 0
	for i := 0; i < len(rec)-1; i++ {
		if absSub(rec[i][0], rec[i+1][0]) == 1 {
			count += (rec[i][1] + rec[i+1][1])
		}
	}

	fmt.Println(count)
	return count
}

func absSub(a, b int) int {
	if a-b < 0 {
		return -1 * (a - b)
	} else {
		return a - b
	}
}

func main() {
	a := []int{1, 3, 2, 2, 5, 2, 3, 7}
	findLHS(a)
}
