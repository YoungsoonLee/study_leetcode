/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.85%)
 * Likes:    4327
 * Dislikes: 202
 * Total Accepted:    890.5K
 * Total Submissions: 2.3M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

// @lc code=start
func isValid(s string) bool {

	/* fail. because It's need a stack.
	m := make(map[string]int)

	for _, c := range s {
		switch string(c) {
		case ")":
			if _, ok := m["("]; ok {
				m["("]--
			}
		case "]":
			if _, ok := m["["]; ok {
				m["["]--
			}
		case "}":
			if _, ok := m["{"]; ok {
				m["{"]--
			}
		default:
			m[string(c)]++
		}
	}

	fmt.Println(m)

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
	*/
}

// @lc code=end

