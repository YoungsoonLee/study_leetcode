package main

import "fmt"

func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	// dp init
	for i := 0; i < n+1; i++ {
		dp[i] = i - 1
	}
	fmt.Println("dp: ", dp)

	for i := 0; i < n+1; i++ {

	}
	return 0
}

func main() {
	s := "aab"
	minCut(s)
}
