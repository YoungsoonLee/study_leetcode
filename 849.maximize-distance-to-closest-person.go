/*
 * @lc app=leetcode id=849 lang=golang
 *
 * [849] Maximize Distance to Closest Person
 *
 * https://leetcode.com/problems/maximize-distance-to-closest-person/description/
 *
 * algorithms
 * Easy (41.50%)
 * Likes:    524
 * Dislikes: 84
 * Total Accepted:    43.5K
 * Total Submissions: 104.2K
 * Testcase Example:  '[1,0,0,0,1,0,1]'
 *
 * In a row of seats, 1 represents a person sitting in that seat, and 0
 * represents that the seat is empty.
 *
 * There is at least one empty seat, and at least one person sitting.
 *
 * Alex wants to sit in the seat such that the distance between him and the
 * closest person to him is maximized.
 *
 * Return that maximum distance to closest person.
 *
 *
 * Example 1:
 *
 *
 * Input: [1,0,0,0,1,0,1]
 * Output: 2
 * Explanation:
 * If Alex sits in the second open seat (seats[2]), then the closest person has
 * distance 2.
 * If Alex sits in any other open seat, the closest person has distance 1.
 * Thus, the maximum distance to the closest person is 2.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,0,0,0]
 * Output: 3
 * Explanation:
 * If Alex sits in the last seat, the closest person is 3 seats away.
 * This is the maximum distance possible, so the answer is 3.
 *
 *
 * Note:
 *
 *
 * 1 <= seats.length <= 20000
 * seats contains only 0s or 1s, at least one 0, and at least one 1.
 *
 *
 *
 *
 */

// @lc code=start
func maxDistToClosest(seats []int) int {
	size := len(seats)
	maxDis := 0
	// e는 연속 공석 수를 나타냅니다.
	// 연속 공석의 양쪽에 사람이있는 경우 maxDis = (e + e % 2) / 2
	// 측면에 아무도 없으면 maxDis = e
	e := 0
	for i := 0; i < size; i++ {
		if e == i {
			// 설명 좌석 [0 : i]은 모두 0입니다.
			maxDis = e
		} else {
			maxDis = max(maxDis, (e+e%2)/2) // 그냥 -1 하면 될 듯 한데...!!!
		}
		if seats[i] == 1 {
			e = 0
		} else {
			e++
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

