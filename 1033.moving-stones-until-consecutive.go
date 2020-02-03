/*
 * @lc app=leetcode id=1033 lang=golang
 *
 * [1033] Moving Stones Until Consecutive
 *
 * https://leetcode.com/problems/moving-stones-until-consecutive/description/
 *
 * algorithms
 * Easy (39.36%)
 * Likes:    47
 * Dislikes: 351
 * Total Accepted:    9.9K
 * Total Submissions: 24.9K
 * Testcase Example:  '1\n2\n5'
 *
 * Three stones are on a number line at positions a, b, and c.
 * 
 * Each turn, you pick up a stone at an endpoint (ie., either the lowest or
 * highest position stone), and move it to an unoccupied position between those
 * endpoints.  Formally, let's say the stones are currently at positions x, y,
 * z with x < y < z.  You pick up the stone at either position x or position z,
 * and move that stone to an integer position k, with x < k < z and k != y.
 * 
 * The game ends when you cannot make any more moves, ie. the stones are in
 * consecutive positions.
 * 
 * When the game ends, what is the minimum and maximum number of moves that you
 * could have made?  Return the answer as an length 2 array: answer =
 * [minimum_moves, maximum_moves]
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: a = 1, b = 2, c = 5
 * Output: [1,2]
 * Explanation: Move the stone from 5 to 3, or move the stone from 5 to 4 to
 * 3.
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: a = 4, b = 3, c = 2
 * Output: [0,0]
 * Explanation: We cannot make any moves.
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: a = 3, b = 5, c = 1
 * Output: [1,2]
 * Explanation: Move the stone from 1 to 4; or move the stone from 1 to 2 to
 * 4.
 * 
 * 
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= a <= 100
 * 1 <= b <= 100
 * 1 <= c <= 100
 * a != b, b != c, c != a
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
var zerozero = []int{0, 0}

func numMovesStones(a int, b int, c int) []int {
	a,b,c = sort(a,b,c)

	if c-a == 2 {
		return zerozero
	}

	minM, maxM := 2, c-a-2 // !!!
	
	if min(b-a, c-b) <2 {
		minM = 1
	}

	return []int{minM, maxM}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sort(a,b,c int) (int, int, int) {
	if a > b {
		a, b = b, a
	}

	if a > c {
		a, c = c, a
	}

	if b > c {
		b,c = c, b
	}
	return a,b,c
}
// @lc code=end

