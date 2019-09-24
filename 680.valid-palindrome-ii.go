/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 *
 * https://leetcode.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (34.85%)
 * Likes:    897
 * Dislikes: 63
 * Total Accepted:    92.4K
 * Total Submissions: 264.8K
 * Testcase Example:  '"aba"'
 *
 *
 * Given a non-empty string s, you may delete at most one character.  Judge
 * whether you can make it a palindrome.
 *
 *
 * Example 1:
 *
 * Input: "aba"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "abca"
 * Output: True
 * Explanation: You could delete the character 'c'.
 *
 *
 *
 * Note:
 *
 * The string will only contain lowercase characters a-z.
 * The maximum length of the string is 50000.
 *
 *
 */
func validPalindrome(s string) bool {
	return helper([]byte(s), 0, len(s)-1, false)
}

func helper(bs []byte, lo, hi int, hasDeleted bool) bool {
	for lo < hi {
		if bs[lo] != bs[hi] {
			if hasDeleted {
				return false
			}

			return helper(bs, lo+1, hi, true) || helper(bs, lo, hi-1, true)
		}

		lo++
		hi--
	}

	return true
}

