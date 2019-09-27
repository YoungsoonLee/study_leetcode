import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=686 lang=golang
 *
 * [686] Repeated String Match
 *
 * https://leetcode.com/problems/repeated-string-match/description/
 *
 * algorithms
 * Easy (31.68%)
 * Likes:    561
 * Dislikes: 565
 * Total Accepted:    75K
 * Total Submissions: 236.4K
 * Testcase Example:  '"abcd"\n"cdabcdab"'
 *
 * Given two strings A and B, find the minimum number of times A has to be
 * repeated such that B is a substring of it. If no such solution, return -1.
 *
 * For example, with A = "abcd" and B = "cdabcdab".
 *
 * Return 3, because by repeating A three times (“abcdabcdabcd”), B is a
 * substring of it; and B is not a substring of A repeated two times
 * ("abcdabcd").
 *
 * Note:
 * The length of A and B will be between 1 and 10000.
 *
 */
func repeatedStringMatch(a string, b string) int {
	times := (len(b)-1)/len(a) + 1

	fmt.Println(times)
	fmt.Println(strings.Repeat(a, times))

	if strings.Contains(strings.Repeat(a, times), b) {
		return times
	}

	if strings.Contains(strings.Repeat(a, times+1), b) { // 왜 딱 1번만 더할가...
		return times + 1
	}

	return -1
}

