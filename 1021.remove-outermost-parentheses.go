import "strings"

/*
 * @lc app=leetcode id=1021 lang=golang
 *
 * [1021] Remove Outermost Parentheses
 *
 * https://leetcode.com/problems/remove-outermost-parentheses/description/
 *
 * algorithms
 * Easy (76.23%)
 * Likes:    306
 * Dislikes: 443
 * Total Accepted:    65.5K
 * Total Submissions: 85.7K
 * Testcase Example:  '"(()())(())"'
 *
 * A valid parentheses string is either empty (""), "(" + A + ")", or A + B,
 * where A and B are valid parentheses strings, and + represents string
 * concatenation.  For example, "", "()", "(())()", and "(()(()))" are all
 * valid parentheses strings.
 *
 * A valid parentheses string S is primitive if it is nonempty, and there does
 * not exist a way to split it into S = A+B, with A and B nonempty valid
 * parentheses strings.
 *
 * Given a valid parentheses string S, consider its primitive decomposition: S
 * = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.
 *
 * Return S after removing the outermost parentheses of every primitive string
 * in the primitive decomposition of S.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "(()())(())"
 * Output: "()()()"
 * Explanation:
 * The input string is "(()())(())", with primitive decomposition "(()())" +
 * "(())".
 * After removing outer parentheses of each part, this is "()()" + "()" =
 * "()()()".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "(()())(())(()(()))"
 * Output: "()()()()(())"
 * Explanation:
 * The input string is "(()())(())(()(()))", with primitive decomposition
 * "(()())" + "(())" + "(()(()))".
 * After removing outer parentheses of each part, this is "()()" + "()" +
 * "()(())" = "()()()()(())".
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "()()"
 * Output: ""
 * Explanation:
 * The input string is "()()", with primitive decomposition "()" + "()".
 * After removing outer parentheses of each part, this is "" + "" = "".
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * S.length <= 10000
 * S[i] is "(" or ")"
 * S is a valid parentheses string
 *
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	var sb strings.Builder

	i, count, size := 0, 0, len(S)
	for j := 0; j < size; j++ {
		if S[j] == '(' {
			count++
			continue
		}
		count--
		if count == 0 { // S[i] and S[j] are outer parentheses
			sb.WriteString(S[i+1 : j]) // !!!
			i = j + 1                  // !!1
		}
	}

	return sb.String()
}

// @lc code=end

