package main

import (
	"fmt"
	"sort"
)

/*
func findContentChildren(g []int, s []int) int {
	count := 0
	for si := 0; si < len(s); si++ {
		for gi := 0; gi < len(g); gi++ {
			if s[si] == g[gi] {
				count++
			}
		}

	}

	fmt.Println(count)
	return count
}
*/

// g: greed factor, s: size cookie?
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	fmt.Println(g, s)

	var i, j, res int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}
	fmt.Println(res)
	return res
}

func main() {
	a := []int{1, 2}
	s := []int{1, 2, 3}
	findContentChildren(a, s)
}
