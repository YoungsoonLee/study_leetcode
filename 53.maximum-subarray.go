/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (45.38%)
 * Likes:    6961
 * Dislikes: 318
 * Total Accepted:    913.4K
 * Total Submissions: 2M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 * 
 * Example:
 * 
 * 
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 * 
 * 
 * Follow up:
 * 
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 * 
 */

// @lc code=start
func maxSubArray(nums []int) int {
	sum, maxSum := -1<<31, -1<<31
	for _, n:= range nums{
		sum = max(sum+n, n)
		fmt.Printf("sum: %d\n", sum)
		maxSum = max(maxSum, sum)
		fmt.Printf("maxSum: %d\n", maxSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

