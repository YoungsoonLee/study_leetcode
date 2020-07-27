/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 *
 * https://leetcode.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (42.69%)
 * Likes:    882
 * Dislikes: 193
 * Total Accepted:    349.1K
 * Total Submissions: 800.1K
 * Testcase Example:  '1'
 *
 * Given an integer, write a function to determine if it is a power of two.
 * 
 * Example 1:
 * 
 * 
 * Input: 1
 * Output: true 
 * Explanation: 2^0 = 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 16
 * Output: true
 * Explanation: 2^4 = 16
 * 
 * Example 3:
 * 
 * 
 * Input: 218
 * Output: false
 * 
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
    if n < 1{
		return false
	}

	for n > 1{
		if n%2 == 1 {
			return false
		}

		n /= 2
	}

	return true
}
// @lc code=end

