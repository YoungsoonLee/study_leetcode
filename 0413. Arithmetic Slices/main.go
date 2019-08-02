package main

import "sort"

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}

	sort.Ints(A)

	m := 0
	j := 0
	for i := 0; i < len(A); i++ {
		j = i + 1

	}
}

func main() {

}
