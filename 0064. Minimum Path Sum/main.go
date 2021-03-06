package main

import "fmt"

func minPathSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	fmt.Println(m, n)

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	fmt.Println(dp)

	return dp[m-1][n-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	g := [][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}
	//g := [][]int{[]int{1, 2, 5}, []int{3, 2, 1}}
	fmt.Println(minPathSum(g))
}
