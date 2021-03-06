import "fmt"

/*
 * @lc app=leetcode id=1160 lang=golang
 *
 * [1160] Find Words That Can Be Formed by Characters
 *
 * https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/description/
 *
 * algorithms
 * Easy (66.89%)
 * Likes:    197
 * Dislikes: 47
 * Total Accepted:    32.6K
 * Total Submissions: 48.6K
 * Testcase Example:  '["cat","bt","hat","tree"]\n"atach"'
 *
 * You are given an array of strings words and a string chars.
 *
 * A string is good if it can be formed by characters from chars (each
 * character can only be used once).
 *
 * Return the sum of lengths of all good strings in words.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: words = ["cat","bt","hat","tree"], chars = "atach"
 * Output: 6
 * Explanation:
 * The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 =
 * 6.
 *
 *
 * Example 2:
 *
 *
 * Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
 * Output: 10
 * Explanation:
 * The strings that can be formed are "hello" and "world" so the answer is 5 +
 * 5 = 10.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= words.length <= 1000
 * 1 <= words[i].length, chars.length <= 100
 * All strings contain lowercase English letters only.
 *
 */

// @lc code=start
func countCharacters(words []string, chars string) int {
	char := parse(chars)
	fmt.Println(char)

	res := 0
	for _, word := range words {
		res += count(char, word)
	}

	return res
}

func parse(s string) []int {
	res := make([]int, 26)
	for _, r := range s {
		res[r-'a']++
	}
	return res
}

func count(char []int, word string) int {
	res := 0
	w := make([]int, 26)
	for _, r := range word {
		b := r - 'a'
		w[b]++
		if w[b] > char[b] {
			return 0
		}
		res++
	}
	return res
}

// @lc code=end

