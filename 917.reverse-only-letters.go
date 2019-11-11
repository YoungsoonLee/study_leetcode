/*
 * @lc app=leetcode id=917 lang=golang
 *
 * [917] Reverse Only Letters
 *
 * https://leetcode.com/problems/reverse-only-letters/description/
 *
 * algorithms
 * Easy (56.04%)
 * Likes:    327
 * Dislikes: 30
 * Total Accepted:    38.2K
 * Total Submissions: 67.8K
 * Testcase Example:  '"ab-cd"'
 *
 * Given a string S, return the "reversed" string where all characters that are
 * not a letter stay in the same place, and all letters reverse their
 * positions.
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "ab-cd"
 * Output: "dc-ba"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "a-bC-dEf-ghIj"
 * Output: "j-Ih-gfE-dCba"
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "Test1ng-Leet=code-Q!"
 * Output: "Qedo1ct-eeLg=ntse-T!"
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * S.length <= 100
 * 33 <= S[i].ASCIIcode <= 122
 * S doesn't contain \ or "
 *
 *
 *
 *
 *
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	bs := []byte(S)

	left, right := 0, len(bs)-1
	for left < right {
		for left < right && !isLetter(bs[left]) {
			left++
		}
		for left < right && !isLetter(bs[right]) {
			right--
		}
		bs[left], bs[right] = bs[right], bs[left]
		left++
		right--
	}

	return string(bs)
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' ||
		'A' <= b && b <= 'Z'
}

// @lc code=end

