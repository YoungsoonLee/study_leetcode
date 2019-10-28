/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 *
 * https://leetcode.com/problems/backspace-string-compare/description/
 *
 * algorithms
 * Easy (46.59%)
 * Likes:    837
 * Dislikes: 51
 * Total Accepted:    80.5K
 * Total Submissions: 172.4K
 * Testcase Example:  '"ab#c"\n"ad#c"'
 *
 * Given two strings S and T, return if they are equal when both are typed into
 * empty text editors. # means a backspace character.
 *
 *
 * Example 1:
 *
 *
 * Input: S = "ab#c", T = "ad#c"
 * Output: true
 * Explanation: Both S and T become "ac".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: S = "ab##", T = "c#d#"
 * Output: true
 * Explanation: Both S and T become "".
 *
 *
 *
 * Example 3:
 *
 *
 * Input: S = "a##c", T = "#a#c"
 * Output: true
 * Explanation: Both S and T become "c".
 *
 *
 *
 * Example 4:
 *
 *
 * Input: S = "a#c", T = "b"
 * Output: false
 * Explanation: S becomes "c" while T becomes "b".
 *
 *
 * Note:
 *
 *
 * 1 <= S.length <= 200
 * 1 <= T.length <= 200
 * S and T only contain lowercase letters and '#' characters.
 *
 *
 * Follow up:
 *
 *
 * Can you solve it in O(N) time and O(1) space?
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func backspaceCompare(S string, T string) bool {
	i := len(S)
	j := len(T)

	for i >= 0 || j >= 0 {
		i = nextIndex(&S, i)
		j = nextIndex(&T, j)

		if i >= 0 && j >= 0 && S[i] != T[j] {
			return false
		}

	}

	return i == j
}

// 返回 s[:i] 中，不是 '#' 的字符的最大的索引号
func nextIndex(s *string, i int) int {
	i--
	count := 0
	for i >= 0 && ((*s)[i] == '#' || count > 0) {
		if (*s)[i] == '#' {
			count++
		} else {
			count--
		}
		i--
	}
	return i
}

// @lc code=end

