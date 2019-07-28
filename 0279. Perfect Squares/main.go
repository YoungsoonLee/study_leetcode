package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}
	fmt.Println(perfects) // n 까지의 퍼펙트 수를 담았네 ...
	// dp[i] 表示 the least number of perfect square numbers which sum to i
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	fmt.Println("dp1: ", dp)

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				// because i = ( i - p ) + p,p is the square number
				// So dp[i] = dp[i-p] + 1
				dp[i] = dp[i-p] + 1
			}
		}
	}

	fmt.Println("dp2: ", dp)
	return dp[n]
}

func main() {
	n := 13
	numSquares(n)
}
