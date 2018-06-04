package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	dp := make([]int, n+1)
	//fmt.Println(dp)

	for i := 2; i <= n; i++ {
		//fmt.Println(dp[i-1], cost[i-1], dp[i-2], cost[i-2])
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
		//fmt.Println(dp)
	}
	//fmt.Println(dp, n)
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	c := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	minCostClimbingStairs(c)
}
