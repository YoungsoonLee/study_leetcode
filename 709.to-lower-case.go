import "strings"

/*
 * @lc app=leetcode id=709 lang=golang
 *
 * [709] To Lower Case
 *
 * https://leetcode.com/problems/to-lower-case/description/
 *
 * algorithms
 * Easy (77.52%)
 * Likes:    337
 * Dislikes: 1172
 * Total Accepted:    149.9K
 * Total Submissions: 192.9K
 * Testcase Example:  '"Hello"'
 *
 * Implement function ToLowerCase() that has a string parameter str, and
 * returns the same string in lowercase.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "Hello"
 * Output: "hello"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "here"
 * Output: "here"
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "LOVELY"
 * Output: "lovely"
 *
 *
 *
 *
 *
 */

// @lc code=start
func toLowerCase(str string) string {
	return strings.ToLower(str)
}

// @lc code=end

