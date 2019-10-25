import "fmt"

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
	s1 := make([]byte, 0, len(S))
	t1 := make([]byte, 0, len(T))

	for i := 0; i < len(S); i++ {
		if S[i] == '#' {
			s1 = s1[:i-1]
		} else {
			s1 = append(s1, S[i])
		}
	}

	fmt.Println(s1)

	for i := 0; i < len(T); i++ {
		if T[i] == '#' {
			t1 = t1[:i-1]
		} else {
			t1 = append(t1, T[i])
		}
	}

	fmt.Println(t1)

	return s1 == t1

}

// @lc code=end

