import (
	"strings"
)

/*
 * @lc app=leetcode id=748 lang=golang
 *
 * [748] Shortest Completing Word
 *
 * https://leetcode.com/problems/shortest-completing-word/description/
 *
 * algorithms
 * Easy (55.10%)
 * Likes:    139
 * Dislikes: 496
 * Total Accepted:    25K
 * Total Submissions: 45K
 * Testcase Example:  '"1s3 PSt"\n["step","steps","stripe","stepple"]'
 *
 *
 * Find the minimum length word from a given dictionary words, which has all
 * the letters from the string licensePlate.  Such a word is said to complete
 * the given string licensePlate
 *
 * Here, for letters we ignore case.  For example, "P" on the licensePlate
 * still matches "p" on the word.
 *
 * It is guaranteed an answer exists.  If there are multiple answers, return
 * the one that occurs first in the array.
 *
 * The license plate might have the same letter occurring multiple times.  For
 * example, given a licensePlate of "PP", the word "pair" does not complete the
 * licensePlate, but the word "supper" does.
 *
 *
 * Example 1:
 *
 * Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe",
 * "stepple"]
 * Output: "steps"
 * Explanation: The smallest length word that contains the letters "S", "P",
 * "S", and "T".
 * Note that the answer is not "step", because the letter "s" must occur in the
 * word twice.
 * Also note that we ignored case for the purposes of comparing whether a
 * letter exists in the word.
 *
 *
 *
 * Example 2:
 *
 * Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
 * Output: "pest"
 * Explanation: There are 3 smallest length words that contains the letters
 * "s".
 * We return the one that occurred first.
 *
 *
 *
 * Note:
 *
 * licensePlate will be a string with length in range [1, 7].
 * licensePlate will contain digits, spaces, or letters (uppercase or
 * lowercase).
 * words will have a length in the range [10, 1000].
 * Every words[i] will consist of lowercase letters, and have length in range
 * [1, 15].
 *
 *
 */

// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	keys := getKeys(licensePlate)
	minLen := 1<<63 - 1
	res := ""

	for _, w := range words {
		if len(w) >= minLen {
			continue
		}

		isCompleting := true
		for k, c := range keys {
			if strings.Count(w, k) < c {
				isCompleting = false
				break
			}
		}

		if isCompleting {
			res = w
			minLen = len(w)
		}
	}

	return res
}

func getKeys(licensePlate string) map[string]int {
	licensePlate = strings.ToLower(licensePlate)

	bs := []byte(licensePlate) // 숫자제거를 위해서

	res := make(map[string]int, len(licensePlate))

	for _, b := range bs {
		if 'a' <= b && b <= 'z' {
			res[string(b)]++
		}
	}

	return res
}

// @lc code=end

