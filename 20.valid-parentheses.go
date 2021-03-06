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
	stack := make([]string, 0)
	opa := make(map[string]string)
	opa["("] = ")"
	opa["{"] = "}"
	opa["["] = "]"

	cpa := make(map[string]string)
	cpa[")"] = ")"
	cpa["}"] = "}"
	cpa["]"] = "]"

	for _, c := range s {
		ov, ook := opa[string(c)]
		if string(c) != ov && ook { //open
			stack = append(stack, ov)
		}

		cv, cok := cpa[string(c)]
		if cok { //close
			if len(stack) == 0 {
				return false
			}

			if cv != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

// @lc code=end

