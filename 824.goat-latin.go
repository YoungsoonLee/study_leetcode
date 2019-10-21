import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=824 lang=golang
 *
 * [824] Goat Latin
 *
 * https://leetcode.com/problems/goat-latin/description/
 *
 * algorithms
 * Easy (58.49%)
 * Likes:    184
 * Dislikes: 464
 * Total Accepted:    40.8K
 * Total Submissions: 69.2K
 * Testcase Example:  '"I speak Goat Latin"'
 *
 * A sentence S is given, composed of words separated by spaces. Each word
 * consists of lowercase and uppercase letters only.
 *
 * We would like to convert the sentence to "Goat Latin" (a made-up language
 * similar to Pig Latin.)
 *
 * The rules of Goat Latin are as follows:
 *
 *
 * If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of
 * the word.
 * For example, the word 'apple' becomes 'applema'.
 *
 * If a word begins with a consonant (i.e. not a vowel), remove the first
 * letter and append it to the end, then add "ma".
 * For example, the word "goat" becomes "oatgma".
 *
 * Add one letter 'a' to the end of each word per its word index in the
 * sentence, starting with 1.
 * For example, the first word gets "a" added to the end, the second word gets
 * "aa" added to the end and so on.
 *
 *
 * Return the final sentence representing the conversion from S to Goat
 * Latin.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "I speak Goat Latin"
 * Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
 *
 *
 * Example 2:
 *
 *
 * Input: "The quick brown fox jumped over the lazy dog"
 * Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa
 * hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
 *
 *
 *
 *
 * Notes:
 *
 *
 * S contains only uppercase, lowercase and spaces. Exactly one space between
 * each word.
 * 1 <= S.length <= 150.
 *
 *
 */

// @lc code=start
func toGoatLatin(S string) string {
	res := parse(S)

	for i, gl := range res {
		switch gl[0] {
		case 'a', 'e', 'i', 'o', 'u':
			res[i] = gl + "ma"
		default:
			res[i] = gl[1:] + "ma"
		}
	}

	for i, gl := range res {
		if gl[len(gl-1)] == 'a' {
			for j := 0; j < i+1; j++ {
				res[i] = gl + "a"
			}
		}
	}

	fmt.Println(res)
	return strings.Join(s, "")
}

func parse(s string) []string {
	return strings.Split(s, " ")
}

// @lc code=end

