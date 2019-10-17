import "math"

/*
 * @lc app=leetcode id=754 lang=golang
 *
 * [754] Reach a Number
 *
 * https://leetcode.com/problems/reach-a-number/description/
 *
 * algorithms
 * Easy (33.27%)
 * Likes:    341
 * Dislikes: 287
 * Total Accepted:    13.1K
 * Total Submissions: 39.1K
 * Testcase Example:  '1'
 *
 *
 * You are standing at position 0 on an infinite number line.  There is a goal
 * at position target.
 *
 * On each move, you can either go left or right.  During the n-th move
 * (starting from 1), you take n steps.
 *
 * Return the minimum number of steps required to reach the destination.
 *
 *
 * Example 1:
 *
 * Input: target = 3
 * Output: 2
 * Explanation:
 * On the first move we step from 0 to 1.
 * On the second step we step from 1 to 3.
 *
 *
 *
 * Example 2:
 *
 * Input: target = 2
 * Output: 3
 * Explanation:
 * On the first move we step from 0 to 1.
 * On the second move we step  from 1 to -1.
 * On the third move we step from -1 to 2.
 *
 *
 *
 * Note:
 * target will be a non-zero integer in the range [-10^9, 10^9].
 *
 */

// @lc code=start
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	n := int(math.Ceil())
}

// @lc code=end

