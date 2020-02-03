/*
 * @lc app=leetcode id=1037 lang=golang
 *
 * [1037] Valid Boomerang
 *
 * https://leetcode.com/problems/valid-boomerang/description/
 *
 * algorithms
 * Easy (37.54%)
 * Likes:    69
 * Dislikes: 188
 * Total Accepted:    13.1K
 * Total Submissions: 34.9K
 * Testcase Example:  '[[1,1],[2,3],[3,2]]'
 *
 * A boomerang is a set of 3 points that are all distinct and not in a straight
 * line.
 * 
 * Given a list of three points in the plane, return whether these points are a
 * boomerang.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: [[1,1],[2,3],[3,2]]
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [[1,1],[2,2],[3,3]]
 * Output: false
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * points.length == 3
 * points[i].length == 2
 * 0 <= points[i][j] <= 100
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
func isBoomerang(p [][]int) bool {
	x0, y0 := p[0][0], p[0][1]
	x1, y1 := p[1][0], p[1][1]
	x2, y2 := p[2][0], p[2][1]

	return (x0-x2)*(y1-y2) != (y0-y2)*(x1-x2) // !!!
}
// @lc code=end

