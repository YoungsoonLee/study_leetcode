import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=819 lang=golang
 *
 * [819] Most Common Word
 *
 * https://leetcode.com/problems/most-common-word/description/
 *
 * algorithms
 * Easy (42.60%)
 * Likes:    383
 * Dislikes: 955
 * Total Accepted:    82.8K
 * Total Submissions: 193.8K
 * Testcase Example:  '"Bob hit a ball, the hit BALL flew far after it was hit."\n["hit"]'
 *
 * Given a paragraph and a list of banned words, return the most frequent word
 * that is not in the list of banned words.  It is guaranteed there is at least
 * one word that isn't banned, and that the answer is unique.
 *
 * Words in the list of banned words are given in lowercase, and free of
 * punctuation.  Words in the paragraph are not case sensitive.  The answer is
 * in lowercase.
 *
 *
 *
 * Example:
 *
 *
 * Input:
 * paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
 * banned = ["hit"]
 * Output: "ball"
 * Explanation:
 * "hit" occurs 3 times, but it is a banned word.
 * "ball" occurs twice (and no other word does), so it is the most frequent
 * non-banned word in the paragraph.
 * Note that words in the paragraph are not case sensitive,
 * that punctuation is ignored (even if adjacent to words, such as "ball,"),
 * and that "hit" isn't the answer even though it occurs more because it is
 * banned.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= paragraph.length <= 1000.
 * 0 <= banned.length <= 100.
 * 1 <= banned[i].length <= 10.
 * The answer is unique, and written in lowercase (even if its occurrences in
 * paragraph may have uppercase symbols, and even if it is a proper noun.)
 * paragraph only consists of letters, spaces, or the punctuation symbols
 * !?',;.
 * There are no hyphens or hyphenated words.
 * Words only consist of letters, never apostrophes or other punctuation
 * symbols.
 *
 *
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	// map
	m := make(map[string]int, len(paragraph))
	b := make(map[string]bool, len(banned))
	// init banned to map
	for _, bs := range banned {
		//if _, ok := b[bs]; !ok {
		b[bs] = true
		//}
	}

	// split & count if not banned word
	paragraph = strings.Replace(paragraph, ",", "", -1)
	paragraph = strings.Replace(paragraph, ".", "", -1)
	paragraph = strings.ToLower(paragraph)

	ps := parse(paragraph)
	fmt.Println(ps)

	for _, p := range ps {
		if _, ok := b[p]; !ok {
			m[p]++
		}
	}

	fmt.Println(m)

	// check max in map
	maxCount := 0
	maxWord := ""
	for k, v := range m {
		if maxCount < v {
			maxCount = v
			maxWord = k
		}
	}

	return maxWord
}

func parse(s string) []string {
	ss := make([]string, 0)

	ss = strings.Split(s, " ")
	return ss
}

// @lc code=end

