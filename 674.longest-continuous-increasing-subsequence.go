import "fmt"

/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/
 *
 * algorithms
 * Easy (44.77%)
 * Likes:    496
 * Dislikes: 101
 * Total Accepted:    77.4K
 * Total Submissions: 172.7K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 *
 * Given an unsorted array of integers, find the length of longest continuous
 * increasing subsequence (subarray).
 *
 *
 * Example 1:
 *
 * Input: [1,3,5,4,7]
 * Output: 3
 * Explanation: The longest continuous increasing subsequence is [1,3,5], its
 * length is 3.
 * Even though [1,3,5,7] is also an increasing subsequence, it's not a
 * continuous one where 5 and 7 are separated by 4.
 *
 *
 *
 * Example 2:
 *
 * Input: [2,2,2,2,2]
 * Output: 1
 * Explanation: The longest continuous increasing subsequence is [2], its
 * length is 1.
 *
 *
 *
 * Note:
 * Length of the array will not exceed 10,000.
 *
 */
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 1

	i, j := 0, 1

	for j < len(nums) {
		for j < len(nums) && nums[j-1] < nums[j] {
			//fmt.Println(j, i, res)
			j++
		}
		fmt.Println(j, i, res)

		if res < j-i {
			res = j - i
		}

		i = j
		j++
	}

	return res
}


