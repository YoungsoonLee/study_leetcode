/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (46.46%)
 * Likes:    702
 * Dislikes: 194
 * Total Accepted:    265.8K
 * Total Submissions: 556.9K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 * 
 * Note that the row index starts from 0.
 * 
 * 
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 * 
 * Example:
 * 
 * 
 * Input: 3
 * Output: [1,3,3,1]
 * 
 * 
 * Follow up:
 * 
 * Could you optimize your algorithm to use only O(k) extra space?
 * 
 */

// @lc code=start
func getRow(rowIndex int) []int {
	res := make([]int, 1, rowIndex+1)
	res[0] = 1

	if rowIndex == 0 {
		return res
	}

	for i:=0; i<rowIndex; i++ {
		res = append(res, 1)
		for j:=len(res)-2; j>0; j-- {
			res[j] += res[j-1]
		}
	}

	fmt.Println(res)
	return res
}
// @lc code=end

