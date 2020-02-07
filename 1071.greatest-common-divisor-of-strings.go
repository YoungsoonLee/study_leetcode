import "strings"

/*
 * @lc app=leetcode id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 *
 * https://leetcode.com/problems/greatest-common-divisor-of-strings/description/
 *
 * algorithms
 * Easy (53.88%)
 * Likes:    277
 * Dislikes: 83
 * Total Accepted:    17.6K
 * Total Submissions: 32.9K
 * Testcase Example:  '"ABCABC"\n"ABC"'
 *
 * For strings S and T, we say "T divides S" if and only if S = T + ... + T  (T
 * concatenated with itself 1 or more times)
 *
 * Return the largest string X such that X divides str1 and X divides str2.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: str1 = "ABCABC", str2 = "ABC"
 * Output: "ABC"
 *
 *
 * Example 2:
 *
 *
 * Input: str1 = "ABABAB", str2 = "ABAB"
 * Output: "AB"
 *
 *
 * Example 3:
 *
 *
 * Input: str1 = "LEET", str2 = "CODE"
 * Output: ""
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= str1.length <= 1000
 * 1 <= str2.length <= 1000
 * str1[i] and str2[i] are English uppercase letters.
 *
 *
 */

// @lc code=start
func gcdOfStrings(s1 string, s2 string) string {
	l1, l2 := len(s1), len(s2)
	d := gcd(max(l1, l2), min(l1, l2))
	p := s2[:d]
	if s1 == strings.Repeat(p, l1/d) && s2 == strings.Repeat(p, l2/d) {
		return p
	}

	return ""
}

// a >= b
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

