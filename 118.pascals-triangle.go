/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (49.45%)
 * Likes:    1191
 * Dislikes: 98
 * Total Accepted:    352.9K
 * Total Submissions: 693.2K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 * 
 * 
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 * 
 * Example:
 * 
 * 
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 * 
 * 
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i:=1; i<numRows; i++ {
		res = append(res, genNext(res[i-1]))
	}
	return res
}

func genNext(p []int) []int {
	res := make([]int, 1, len(p)+1)
	res = append(res, p...)

	for i:=0; i<len(res)-1; i++ {
		res[i] += res[i+1]
	}
	return res
}
// @lc code=end

