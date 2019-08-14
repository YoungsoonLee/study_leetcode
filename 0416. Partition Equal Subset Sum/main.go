package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	total := 0
	for _, val := range nums {
		total += val
	}

	if total%2 != 0 {
		return false
	}
	target := total / 2
	return backTrack(nums, target, 0)
}

func backTrack(nums []int, target int, idx int) bool {
	if idx == len(nums) || target < 0 {
		return false
	}
	if target == 0 {
		return true
	}
	return backTrack(nums, target-nums[idx], idx+1) || backTrack(nums, target, idx+1)
}

func canPartition(nums []int) bool {
	// use DP
	sum := 0

	for _, n := range nums {
		sum += n
	}

	fmt.Println("1. sum: ", sum)
	fmt.Println("2. sum & 1: ", sum&1) // sum과 & 1 연산은 1인지 0인지 확인 한다.

	if sum&1 == 1 {
		return false
	}

	fmt.Println("3. sum >> 1: ", sum>>1) // 나누기 2
	sum = sum >> 1
	n := len(nums)

	// dp[i][j] represents the elements in nums[:i] ,
	// you can find some, their sum is j

	// 2d array dp init
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	fmt.Println("4. dp: ", dp)

	for i := 0; i < n+1; i++ {
		// Select 0 elements out of any number of elements, the sum is 0
		dp[i][0] = true
	}

	for j := 1; j < sum+1; j++ {
		// From the nums containing 0 elements, the elements are not picked up,
		// making their sum j
		dp[0][j] = false
	}

	fmt.Println("5. dp: ", dp)

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			dp[i][j] = dp[i-1][j]

			if j >= nums[i-1] {
				// nums[:i] has nums[i-1] more than nums[:i-1], so
				// Either an element in nums[:i-1] can be synthesized j-nums[i-1]
				// Either an element in nums[:i-1] can be synthesized j
				// nums[:i] may have element synthesis j
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	fmt.Println("6. dp: ", dp)

	return dp[n][sum]
}

func main() {
	n := []int{1, 2, 3, 5}
	fmt.Println(canPartition(n))
}
