package main

import "fmt"

// brute-force,
// time : O(n^2), space: O(n)
func climbStairs(n int) int {
	return climbing(0, n)
}

func climbing(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return climbing(i+1, n) + climbing(i+2, n)
}

// using memo array.
// time: O(n), space: O(n)
func climbStairs2(n int) int {
	memo := make([]int, n+1)
	return climbing2(0, n, memo)
}

func climbing2(i, n int, memo []int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	if memo[i] > 0 {
		return memo[i]
	}

	memo[i] = climbing2(i+1, n, memo) + climbing2(i+2, n, memo)
	return memo[i]
}

// Dynamic programing
// tiem: O(n), space: O(n) <- dp
func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	fmt.Println(dp)
	for i := 3; i <= n; i++ {

		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp)
	return dp[n]
}

// fibonacci algo
// tiem: O(n), space: O(1)
func climbStairs4(n int) int {
	if n == 1 {
		return 1
	}

	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

// Binets Method (using matrix)
// Time: O(log(n)), space: O(1)

// Fibonacci Formula
// time: O(log(n)), space: O(1)

func main() {
	n := 5
	//fmt.Println(climbStairs4(n))
	n = 6
	fmt.Println(climbStairs3(n))

	n = 3
	//fmt.Println(climbStairs4(n))
	n = 2
	//fmt.Println(climbStairs4(n))
	n = 44
	//fmt.Println(climbStairs4(n))
}
