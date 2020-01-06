package main

import "fmt"

// @lc code=start
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, 0, len(A))

	// brute force
	for i := range queries {
		idx := queries[i][1]
		val := queries[i][0]
		A[idx] = A[idx] + val

		//check even and sum
		evenSum(A, &res)

		//fmt.Println("A: ", A)
		//fmt.Println("res: ", res)
	}
	return res
}

func evenSum(a []int, res *[]int) {
	temp := 0
	for _, v := range a {
		if v%2 == 0 {
			temp += v
		}
	}

	*res = append(*res, temp)
}

func main() {
	A := []int{1, 2, 3, 4}
	queries := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	fmt.Println(sumEvenAfterQueries(A, queries))
}
