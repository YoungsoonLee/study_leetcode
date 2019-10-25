import "fmt"

/*
 * @lc app=leetcode id=840 lang=golang
 *
 * [840] Magic Squares In Grid
 *
 * https://leetcode.com/problems/magic-squares-in-grid/description/
 *
 * algorithms
 * Easy (36.11%)
 * Likes:    91
 * Dislikes: 792
 * Total Accepted:    15.9K
 * Total Submissions: 43.9K
 * Testcase Example:  '[[4,3,8,4],[9,5,1,9],[2,7,6,2]]'
 *
 * A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to
 * 9 such that each row, column, and both diagonals all have the same sum.
 *
 * Given an grid of integers, how many 3 x 3 "magic square" subgrids are
 * there?  (Each subgrid is contiguous).
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[4,3,8,4],
 * ⁠       [9,5,1,9],
 * ⁠       [2,7,6,2]]
 * Output: 1
 * Explanation:
 * The following subgrid is a 3 x 3 magic square:
 * 438
 * 951
 * 276
 *
 * while this one is not:
 * 384
 * 519
 * 762
 *
 * In total, there is only one magic square inside the given grid.
 *
 *
 * Note:
 *
 *
 * 1 <= grid.length <= 10
 * 1 <= grid[0].length <= 10
 * 0 <= grid[i][j] <= 15
 *
 *
 */

// @lc code=start
func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}

	// loop := 0
	sum := 0
	totalCount := []int{}

	// row & column
	// TODO: set i
	for i := 0; i < 3; i++ {
		sum = 0
		for j := 0; j < 3; j++ {
			sum += grid[i][j]
		}
		totalCount = append(totalCount, sum)
	}
	fmt.Println(totalCount)

	// diagonals top-left to bottom-right
	j := 0
	sum = 0
	for i := 0; i < 3; i++ {
		sum += grid[i][j]
		j++
	}
	totalCount = append(totalCount, sum)

	// diagonals top-right to botton-left
	j = 2
	sum = 0
	for i := 0; i < 3; i++ {
		sum += grid[i][j]
		j--
	}
	totalCount = append(totalCount, sum)
	fmt.Println(totalCount)
	return 0

}

// @lc code=end

