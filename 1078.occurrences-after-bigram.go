import "strings"

/*
 * @lc app=leetcode id=1078 lang=golang
 *
 * [1078] Occurrences After Bigram
 *
 * https://leetcode.com/problems/occurrences-after-bigram/description/
 *
 * algorithms
 * Easy (64.41%)
 * Likes:    84
 * Dislikes: 124
 * Total Accepted:    20.8K
 * Total Submissions: 32.2K
 * Testcase Example:  '"alice is a good girl she is a good student"\n"a"\n"good"'
 *
 * Given words first and second, consider occurrences in some text of the form
 * "first second third", where second comes immediately after first, and third
 * comes immediately after second.
 *
 * For each such occurrence, add "third" to the answer, and return the
 * answer.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: text = "alice is a good girl she is a good student", first = "a",
 * second = "good"
 * Output: ["girl","student"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: text = "we will we will rock you", first = "we", second = "will"
 * Output: ["we","rock"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= text.length <= 1000
 * text consists of space separated words, where each word consists of
 * lowercase English letters.
 * 1 <= first.length, second.length <= 10
 * first and second consist of lowercase English letters.
 *
 *
 *
 */

// @lc code=start
func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	n := len(words)
	res := make([]string, 0, n)
	for i := 0; i+2 < n; i++ {
		if words[i] == first && words[i+1] == second {
			res = append(res, words[i+2])
		}
	}

	return res

	/*
		sp := strings.Split(text, " ")
		tp := sp
		//fmt.Println(sp)

		res := make([]string, 0)

		for i, s := range sp {
			//fmt.Println(tp)
			if s == first {
				//fmt.Println(i)
				tp = tp[i+2:]
				//fmt.Println(tp, i)
				res = append(res, tp[0])
				continue
			}

			if s == second {
				//fmt.Println(i)
				tp = tp[i+2:]
				//fmt.Println(tp, i)
				res = append(res, tp[0])
				continue
			}

		}

		fmt.Println(res)

		return res
	*/
}

// @lc code=end

