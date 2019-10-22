import "fmt"

/*
 * @lc app=leetcode id=830 lang=golang
 *
 * [830] Positions of Large Groups
 *
 * https://leetcode.com/problems/positions-of-large-groups/description/
 *
 * algorithms
 * Easy (48.19%)
 * Likes:    243
 * Dislikes: 66
 * Total Accepted:    32.4K
 * Total Submissions: 67.1K
 * Testcase Example:  '"abbxxxxzzy"'
 *
 * In a string S of lowercase letters, these letters form consecutive groups of
 * the same character.
 *
 * For example, a string like S = "abbxxxxzyy" has the groups "a", "bb",
 * "xxxx", "z" and "yy".
 *
 * Call a group large if it has 3 or more characters.  We would like the
 * starting and ending positions of every large group.
 *
 * The final answer should be in lexicographic order.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "abbxxxxzzy"
 * Output: [[3,6]]
 * Explanation: "xxxx" is the single large group with starting  3 and ending
 * positions 6.
 *
 *
 * Example 2:
 *
 *
 * Input: "abc"
 * Output: []
 * Explanation: We have "a","b" and "c" but no large group.
 *
 *
 * Example 3:
 *
 *
 * Input: "abcdddeeeeaabbbcd"
 * Output: [[3,5],[6,9],[12,14]]
 *
 *
 *
 * Note:  1 <= S.length <= 1000
 *
 */

// @lc code=start
func largeGroupPositions(S string) [][]int {

	res := make([][]int, 0, len(S)/3)
	l, r := 0, 1

	for ; r < len(S); r++ {
		if S[l] != S[r] {
			l = r
			continue
		}
		// Find a large group
		if r-l+1 == 3 {
			res = append(res, []int{l, r})
		} else if r-l+1 > 3 {
			// The last large group continues to grow bigger
			res[len(res)-1][1] = r
		}
	}

	fmt.Println(res)
	return res

}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// @lc code=end

