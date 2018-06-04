package main

import (
	"fmt"
	"sort"
)

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
