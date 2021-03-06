/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.40%)
 * Likes:    560
 * Dislikes: 2215
 * Total Accepted:    349.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word (last word means the last
 * appearing word if we loop from left to right) in the string.
 * 
 * If the last word does not exist, return 0.
 * 
 * Note: A word is defined as a maximal substring consisting of non-space
 * characters only.
 * 
 * Example:
 * 
 * 
 * Input: "Hello World"
 * Output: 5
 * 
 * 
 * 
 * 
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	res := 0
	for i := size-1; i >= 0; i-- {
		if s[i] == ' ' {
				if res != 0 {
					return res
				}
				continue
		}
		res++
	}
	return res
}
// @lc code=end

