/*
 * @lc app=leetcode id=1128 lang=golang
 *
 * [1128] Number of Equivalent Domino Pairs
 *
 * https://leetcode.com/problems/number-of-equivalent-domino-pairs/description/
 *
 * algorithms
 * Easy (46.56%)
 * Likes:    122
 * Dislikes: 75
 * Total Accepted:    16.4K
 * Total Submissions: 35K
 * Testcase Example:  '[[1,2],[2,1],[3,4],[5,6]]'
 *
 * Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j]
 * = [c, d] if and only if either (a==c and b==d), or (a==d and b==c) - that
 * is, one domino can be rotated to be equal to another domino.
 *
 * Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length,
 * and dominoes[i] is equivalent to dominoes[j].
 *
 *
 * Example 1:
 * Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
 * Output: 1
 *
 *
 * Constraints:
 *
 *
 * 1 <= dominoes.length <= 40000
 * 1 <= dominoes[i][j] <= 9
 *
 */

// @lc code=start
func numEquivDominoPairs(dominoes [][]int) int {
	res := 0
	count := [100]int{} // why 100 ???, related on mapping function
	for _, domino := range dominoes {
		d := mapping(domino)
		res += count[d]
		count[d]++
	}

	return res

}

func mapping(A []int) int {
	a, b := A[0], A[1]
	if a < b {
		return a*10 + b
	}
	return b*10 + a
}

// @lc code=end

