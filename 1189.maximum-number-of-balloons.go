import "fmt"

/*
 * @lc app=leetcode id=1189 lang=golang
 *
 * [1189] Maximum Number of Balloons
 *
 * https://leetcode.com/problems/maximum-number-of-balloons/description/
 *
 * algorithms
 * Easy (61.07%)
 * Likes:    166
 * Dislikes: 24
 * Total Accepted:    24.4K
 * Total Submissions: 40.1K
 * Testcase Example:  '"nlaebolko"'
 *
 * Given a string text, you want to use the characters of text to form as many
 * instances of the word "balloon" as possible.
 *
 * You can use each character in text at most once. Return the maximum number
 * of instances that can be formed.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: text = "nlaebolko"
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: text = "loonbalxballpoon"
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: text = "leetcode"
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= text.length <= 10^4
 * text consists of lower case English letters only.
 *
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	//text = "loonbalxballpoon"

	b, a, l, o, n := 0, 0, 0, 0, 0
	for _, r := range text {
		switch r {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}

	fmt.Println(b, a, l, o, n)

	res := min(b, min(a, n))
	res = min(res, min(l/2, o/2))
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

