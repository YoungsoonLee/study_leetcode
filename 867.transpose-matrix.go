import "fmt"

/*
 * @lc app=leetcode id=867 lang=golang
 *
 * [867] Transpose Matrix
 *
 * https://leetcode.com/problems/transpose-matrix/description/
 *
 * algorithms
 * Easy (63.93%)
 * Likes:    255
 * Dislikes: 248
 * Total Accepted:    57.4K
 * Total Submissions: 89.7K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix A, return the transpose of A.
 *
 * The transpose of a matrix is the matrix flipped over it's main diagonal,
 * switching the row and column indices of the matrix.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [[1,4,7],[2,5,8],[3,6,9]]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,2,3],[4,5,6]]
 * Output: [[1,4],[2,5],[3,6]]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 1000
 * 1 <= A[0].length <= 1000
 *
 *
 *
 */

// @lc code=start
func transpose(A [][]int) [][]int {
	m, n := len(A), len(A[0])

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	fmt.Println("res1: ", res)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = A[i][j]
		}
	}

	fmt.Println("res2: ", res)

	return res
}

// @lc code=end

