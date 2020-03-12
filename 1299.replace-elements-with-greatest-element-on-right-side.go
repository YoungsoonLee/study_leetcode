/*
 * @lc app=leetcode id=1299 lang=golang
 *
 * [1299] Replace Elements with Greatest Element on Right Side
 *
 * https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/description/
 *
 * algorithms
 * Easy (75.95%)
 * Likes:    165
 * Dislikes: 47
 * Total Accepted:    19.3K
 * Total Submissions: 24.8K
 * Testcase Example:  '[17,18,5,4,6,1]'
 *
 * Given an array arr, replace every element in that array with the greatest
 * element among the elements to its right, and replace the last element with
 * -1.
 *
 * After doing so, return the array.
 *
 *
 * Example 1:
 * Input: arr = [17,18,5,4,6,1]
 * Output: [18,6,6,6,1,-1]
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 10^4
 * 1 <= arr[i] <= 10^5
 *
 */

// @lc code=start
func replaceElements(arr []int) []int {
	for i, _ := range arr {
		arr[i] = findMax(arr[i+1:])
	}

	arr[len(arr)-1] = -1

	return arr
}

func findMax(arr []int) int {
	res := 0
	for _, n := range arr {
		res = max(res, n)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

