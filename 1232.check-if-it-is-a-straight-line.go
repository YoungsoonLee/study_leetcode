/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 *
 * https://leetcode.com/problems/check-if-it-is-a-straight-line/description/
 *
 * algorithms
 * Easy (47.01%)
 * Likes:    144
 * Dislikes: 17
 * Total Accepted:    16.5K
 * Total Submissions: 35.1K
 * Testcase Example:  '[[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]'
 *
 * You are given an array coordinates, coordinates[i] = [x, y], where [x, y]
 * represents the coordinate of a point. Check if these points make a straight
 * line in the XY plane.
 *
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= coordinates.length <= 1000
 * coordinates[i].length == 2
 * -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
 * coordinates contains no duplicate point.
 *
 *
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	dx0 := coordinates[1][0] - coordinates[0][0]
	dy0 := coordinates[1][1] - coordinates[0][1]

	for i := 1; i < len(coordinates)-1; i++ {
		dx := coordinates[i+1][0] - coordinates[i][0]
		dy := coordinates[i+1][1] - coordinates[i][1]

		if dy*dx0 != dy0*dx { // check cross product
			return false
		}
	}
	return true
}

// @lc code=end

