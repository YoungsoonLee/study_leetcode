/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 *
 * https://leetcode.com/problems/robot-return-to-origin/description/
 *
 * algorithms
 * Easy (72.10%)
 * Likes:    756
 * Dislikes: 584
 * Total Accepted:    186K
 * Total Submissions: 257.8K
 * Testcase Example:  '"UD"'
 *
 * There is a robot starting at position (0, 0), the origin, on a 2D plane.
 * Given a sequence of its moves, judge if this robot ends up at (0, 0) after
 * it completes its moves.
 * 
 * The move sequence is represented by a string, and the character moves[i]
 * represents its ith move. Valid moves are R (right), L (left), U (up), and D
 * (down). If the robot returns to the origin after it finishes all of its
 * moves, return true. Otherwise, return false.
 * 
 * Note: The way that the robot is "facing" is irrelevant. "R" will always make
 * the robot move to the right once, "L" will always make it move left, etc.
 * Also, assume that the magnitude of the robot's movement is the same for each
 * move.
 * 
 * Example 1:
 * 
 * 
 * Input: "UD"
 * Output: true 
 * Explanation: The robot moves up once, and then down once. All moves have the
 * same magnitude, so it ended up at the origin where it started. Therefore, we
 * return true.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "LL"
 * Output: false
 * Explanation: The robot moves left twice. It ends up two "moves" to the left
 * of the origin. We return false because it is not at the origin at the end of
 * its moves.
 * 
 * 
 */
func judgeCircle(moves string) bool {
	// 일단 홀수면 false
	n := len(moves) % 2
	fmt.Println(n)
	if n != 0 {
		return false
	}
	// U <-> D
	// R <-> L
	// string을 루프 돌면서 map에 저장.
	// map 저장 할때, 반대가 있으면 반대 값의 값을 --
	// 루프가 끝내고, map을 검색 해서 0이 아닌값이 있으면 
	m := make(map[byte]int, 4)
	m['U'] = 0
	m['D'] = 0
	m['R'] = 0
	m['L'] = 0

	for _, c := range moves {
		fmt.Println(c)

		switch c {
		case 'U':
			if m['D'] > 0 {
				m['D']--
			}else{
				m['U']++
			}
		case 'D':
			if m['U'] > 0 {
				m['U']--
			}else{
				m['D']++
			}
		case 'R':
			if m['L'] > 0 {
				m['L']--
			}else{
				m['R']++
			}
		case 'L':
			if m['R'] > 0 {
				m['R']--
			}else{
				m['L']++
			}
		}
	}

	//fmt.Println(m)

	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}

