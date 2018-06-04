package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	fmt.Println(len(M))
	res := make([][]int, len(M))

	for i := range res {
		res[i] = make([]int, len(M[0]))
		for j := range res[i] {
			res[i][j] = getValue(M, i, j)
		}
	}

	return res
}

func getValue(M [][]int, r, c int) int {
	value, count := 0, 0
	for i := r - 1; i < r+2; i++ {
		for j := c - 1; j < c+2; j++ {
			if 0 <= i && i < len(M) && 0 <= j && j < len(M[0]) {
				value += M[i][j]
				count++
			}
		}
	}
	return value / count
}

func main() {
	a := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	imageSmoother(a)
}
