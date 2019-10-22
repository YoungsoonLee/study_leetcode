import "fmt"

/*
 * @lc app=leetcode id=821 lang=golang
 *
 * [821] Shortest Distance to a Character
 *
 * https://leetcode.com/problems/shortest-distance-to-a-character/description/
 *
 * algorithms
 * Easy (64.23%)
 * Likes:    712
 * Dislikes: 61
 * Total Accepted:    46.1K
 * Total Submissions: 71.6K
 * Testcase Example:  '"loveleetcode"\n"e"'
 *
 * Given a string S and a character C, return an array of integers representing
 * the shortest distance from the character C in the string.
 *
 * Example 1:
 *
 *
 *
 * Input: S = "loveleetcode", C = 'e'
 * Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
 *
 *
 *
 *
 * Note:
 *
 *
 * S string length is in [1, 10000].
 * C is a single character, and guaranteed to be in string S.
 * All letters in S and C are lowercase.
 *
 *
 */

// @lc code=start
func shortestToChar(S string, C byte) []int {
	n := len(S)

	res := make([]int, n)
	for i := range res {
		res[i] = n
	}

	//fmt.Println(res)

	left, right := -n, 2*n // -12, 24

	for i := 0; i < n; i++ {
		j := n - i - 1
		if S[i] == C {
			left = i
		}
		if S[j] == C {
			right = j
		}

		res[i] = min(res[i], dist(i, left))  // !!!
		res[i] = min(res[j], dist(i, right)) // !!!
	}
	fmt.Println(res)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

// @lc code=end

