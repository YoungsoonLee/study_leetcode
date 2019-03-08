package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix) // row
	if m == 0 {
		return 0
	}

	n := len(matrix[0]) // col
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0') //change first column
		for i := 1; i < m; i++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}

	fmt.Println("first dp: ", dp)

	return 0
}

func main() {
	m := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	maximalRectangle(m)
}
