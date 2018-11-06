package main

import (
	"fmt"
)

func rotate(m [][]int) {
	n := len(m)
	//fmt.Println(n, n/2)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			fmt.Println(n, i, j, n-i-1)

			temp := m[i][j]

			// 왼쪽 행이 오른쪽 열에 동일 함
			m[i][j] = m[n-j-1][i]
			m[n-j-1][i] = m[n-i-1][n-j-1]
			m[n-i-1][n-j-1] = m[j][n-i-1]
			m[j][n-i-1] = temp
		}
	}

	fmt.Println(m)

}

/*
func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
*/

func main() {
	//fmt.Println(abs(-1))
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(a)
}
