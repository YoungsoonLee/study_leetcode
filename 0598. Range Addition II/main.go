package main

import (
	"fmt"
)

func maxCount(m int, n int, ops [][]int) int {
	for _, o := range ops {
		fmt.Println(o[0], o[1])
		
		m = min(m, o[0])
		n = min(n, o[1])
	}
	return m * n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func main() {
	m := 3
	n := 3
	ops := [][]int{{2,2}, {3,3}}
	maxCount(m,n,ops)
}
