import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 *
 * https://leetcode.com/problems/rotate-string/description/
 *
 * algorithms
 * Easy (49.27%)
 * Likes:    467
 * Dislikes: 40
 * Total Accepted:    49K
 * Total Submissions: 99.2K
 * Testcase Example:  '"abcde"\n"cdeab"'
 *
 * We are given two strings, A and B.
 *
 * A shift on A consists of taking string A and moving the leftmost character
 * to the rightmost position. For example, if A = 'abcde', then it will be
 * 'bcdea' after one shift on A. Return True if and only if A can become B
 * after some number of shifts on A.
 *
 *
 * Example 1:
 * Input: A = 'abcde', B = 'cdeab'
 * Output: true
 *
 * Example 2:
 * Input: A = 'abcde', B = 'abced'
 * Output: false
 *
 *
 * Note:
 *
 *
 * A and B will have length at most 100.
 *
 *
 */

// @lc code=start
func rotateString(A string, B string) bool {
	/*
		if len(A) != len(B) {
			return false
		}

		if len(A) == 0 && len(B) == 0 {
			return true
		}

		for i := 0; i < len(A); i++ {
			A = shift(A)
			if A == B {
				return true
			}
		}

		return false
	*/
	return len(A) == len(B) && strings.Contains(A+A, B)
}

func shift(s string) string {
	fmt.Println(s[1:] + string(s[0]))
	return s[1:] + string(s[0])
}

// @lc code=end

