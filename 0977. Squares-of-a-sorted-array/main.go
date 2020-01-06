package main

import (
	"fmt"
	"sort"
)

func sortedSquares(A []int) []int {
	for i := range A {
		A[i] = A[i] * A[i]
	}

	fmt.Println(A)

	sort.Ints(A)
	return A
}

func main() {
	a := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(a))
}
