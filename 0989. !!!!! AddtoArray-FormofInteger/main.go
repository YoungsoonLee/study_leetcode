package main

import (
	"fmt"
)

func addToArrayForm(A []int, K int) []int {
	j := 1
	sum := 0
	for i := len(A) - 1; i >= 0; i-- {
		sum += A[i] * j
		j *= 10

	}

	sum += K
	fmt.Println(sum)
	fmt.Println(numToArray(sum))
	return []int{}
}

func numToArray(a int) []int {
	res := make([]int, 0, 8)
	for a > 0 {
		res = append(res, a%10) // 나머지
		a /= 10                 // 몫업데이트
	}

	reverse(res)
	return res
}

func reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func main() {
	A := []int{1, 2, 0, 0}
	K := 34
	addToArrayForm(A, K)
}
