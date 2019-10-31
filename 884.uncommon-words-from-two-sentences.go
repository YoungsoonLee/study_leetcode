import (
	"fmt"
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 *
 * https://leetcode.com/problems/uncommon-words-from-two-sentences/description/
 *
 * algorithms
 * Easy (61.26%)
 * Likes:    278
 * Dislikes: 69
 * Total Accepted:    39.6K
 * Total Submissions: 64.5K
 * Testcase Example:  '"this apple is sweet"\n"this apple is sour"'
 *
 * We are given two sentences A and B.  (A sentence is a string of space
 * separated words.  Each word consists only of lowercase letters.)
 *
 * A word is uncommon if it appears exactly once in one of the sentences, and
 * does not appear in the other sentence.
 *
 * Return a list of all uncommon words.
 *
 * You may return the list in any order.
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
 * Input: A = "this apple is sweet", B = "this apple is sour"
 * Output: ["sweet","sour"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = "apple apple", B = "banana"
 * Output: ["banana"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= A.length <= 200
 * 0 <= B.length <= 200
 * A and B both contain only spaces and lowercase letters.
 *
 *
 *
 *
 */

// @lc code=start
func uncommonFromSentences(A string, B string) []string {

	tmp := strings.Split(A, " ")
	tmp = append(tmp, strings.Split(B, " ")...)
	tmp = append(tmp, "", "~")
	sort.Strings(tmp)

	// 정렬이 완료되면 ""가 시작 부분에 있고 "~"가 끝 부분에 있습니다
	// 관련이없는 두 단어를 tmp로 정렬하면 판단 논리를 단순화하는 데 도움이됩니다.

	fmt.Println(tmp)

	res := make([]string, 0, len(tmp))

	for i := 1; i < len(tmp)-1; i++ {
		if tmp[i-1] != tmp[i] && tmp[i] != tmp[i+1] {
			res = append(res, tmp[i])
		}
	}

	fmt.Println(res)

	return res

	/*
		sa := strings.Split(A, " ")
		sb := strings.Split(B, " ")

		res := make([]string, 0, len(sa))

		check(sa, sb, &res)
		check(sb, sa, &res)

		return res
	*/

}

func check(s []string, d []string, r *[]string) {
	m := make(map[string]bool, len(s))

	for _, ss := range s {
		m[ss] = true
	}

	for _, ss := range d {
		if _, ok := m[ss]; !ok {
			*r = append(*r, ss)
		}
	}
}

// @lc code=end

