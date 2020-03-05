import "fmt"

/*
 * @lc app=leetcode id=1260 lang=golang
 *
 * [1260] Shift 2D Grid
 *
 * https://leetcode.com/problems/shift-2d-grid/description/
 *
 * algorithms
 * Easy (59.79%)
 * Likes:    116
 * Dislikes: 64
 * Total Accepted:    11.6K
 * Total Submissions: 19.1K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]\n1'
 *
 * Given a 2D grid of size m x n and an integer k. You need to shift the grid k
 * times.
 *
 * In one shift operation:
 *
 *
 * Element at grid[i][j] moves to grid[i][j + 1].
 * Element at grid[i][n - 1] moves to grid[i + 1][0].
 * Element at grid[m - 1][n - 1] moves to grid[0][0].
 *
 *
 * Return the 2D grid after applying shift operation k times.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
 * Output: [[9,1,2],[3,4,5],[6,7,8]]
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
 * Output: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
 *
 *
 * Example 3:
 *
 *
 * Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
 * Output: [[1,2,3],[4,5,6],[7,8,9]]
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m <= 50
 * 1 <= n <= 50
 * -1000 <= grid[i][j] <= 1000
 * 0 <= k <= 100
 *
 *
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
	row := len(grid)
	col := len(grid[0])

	if k == row*col {
		return grid
	}

	tmp := []int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp = append(tmp, grid[i][j])
		}
	}

	fmt.Println("tmp1: ", tmp)
	k = k % (row * col) // !!! why 왜 곱하는지...당연히 메트릭스를 펼치니깐? ...
	tmp = append(tmp[len(tmp)-k:], tmp[:len(tmp)-k]...)
	fmt.Println("tmp2: ", tmp)

	// init matrix
	res := make([][]int, row)
	index := 0
	for i := range res {
		res[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res[i][j] = tmp[index]
			index++
		}
	}
	return res

}

// @lc code=end

