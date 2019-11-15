import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode id=937 lang=golang
 *
 * [937] Reorder Data in Log Files
 *
 * https://leetcode.com/problems/reorder-data-in-log-files/description/
 *
 * algorithms
 * Easy (55.42%)
 * Likes:    318
 * Dislikes: 964
 * Total Accepted:    60.8K
 * Total Submissions: 112.5K
 * Testcase Example:  '["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]'
 *
 * You have an array of logs.  Each log is a space delimited string of words.
 *
 * For each log, the first word in each log is an alphanumeric identifier.
 * Then, either:
 *
 *
 * Each word after the identifier will consist only of lowercase letters,
 * or;
 * Each word after the identifier will consist only of digits.
 *
 *
 * We will call these two varieties of logs letter-logs and digit-logs.  It is
 * guaranteed that each log has at least one word after its identifier.
 *
 * Reorder the logs so that all of the letter-logs come before any digit-log.
 * The letter-logs are ordered lexicographically ignoring identifier, with the
 * identifier used in case of ties.  The digit-logs should be put in their
 * original order.
 *
 * Return the final order of the logs.
 *
 *
 * Example 1:
 * Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit
 * dig","let3 art zero"]
 * Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5
 * 1","dig2 3 6"]
 *
 *
 * Constraints:
 *
 *
 * 0 <= logs.length <= 100
 * 3 <= logs[i].length <= 100
 * logs[i] is guaranteed to have an identifier, and a word after the
 * identifier.
 *
 *
 */

// @lc code=start
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		s1 := strings.SplitN(logs[i], " ", 2)
		s2 := strings.SplitN(logs[j], " ", 2)

		//fmt.Println(s1, s2)

		f1, f2 := "0"+s1[1], "0"+s2[1]

		if unicode.IsNumber(rune(f1[1])) {
			f1 = "1"
		}
		if unicode.IsNumber(rune(f2[1])) {
			f2 = "1"
		}
		return f1 < f2
	})
	fmt.Println(logs)
	return logs
}

// @lc code=end

