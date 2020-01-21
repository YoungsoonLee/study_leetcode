import "strings"

/*
 * @lc app=leetcode id=1002 lang=golang
 *
 * [1002] Find Common Characters
 *
 * https://leetcode.com/problems/find-common-characters/description/
 *
 * algorithms
 * Easy (66.46%)
 * Likes:    541
 * Dislikes: 68
 * Total Accepted:    49.9K
 * Total Submissions: 75K
 * Testcase Example:  '["bella","label","roller"]'
 *
 * Given an array A of strings made only from lowercase letters, return a list
 * of all characters that show up in all strings within the list (including
 * duplicates).  For example, if a character occurs 3 times in all strings but
 * not 4 times, you need to include that character three times in the final
 * answer.
 *
 * You may return the answer in any order.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["bella","label","roller"]
 * Output: ["e","l","l"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["cool","lock","cook"]
 * Output: ["c","o"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 100
 * 1 <= A[i].length <= 100
 * A[i][j] is a lowercase letter
 *
 *
 *
 */

// @lc code=start
func commonChars(A []string) []string {
	res := make([]string, 0, 128)
	for i := 'a'; i <= 'z'; i++ {
		sub := string(i)
		count := strings.Count(A[0], sub)
		for j := 1; j < len(A) && count > 0; j++ {
			count = min(count, strings.Count(A[j], sub)) // !!!
		}
		for count > 0 {
			res = append(res, sub)
			count--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

