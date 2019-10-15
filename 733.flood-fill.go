/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 *
 * https://leetcode.com/problems/flood-fill/description/
 *
 * algorithms
 * Easy (51.72%)
 * Likes:    585
 * Dislikes: 127
 * Total Accepted:    67.1K
 * Total Submissions: 129.3K
 * Testcase Example:  '[[1,1,1],[1,1,0],[1,0,1]]\n1\n1\n2'
 *
 *
 * An image is represented by a 2-D array of integers, each integer
 * representing the pixel value of the image (from 0 to 65535).
 *
 * Given a coordinate (sr, sc) representing the starting pixel (row and column)
 * of the flood fill, and a pixel value newColor, "flood fill" the image.
 *
 * To perform a "flood fill", consider the starting pixel, plus any pixels
 * connected 4-directionally to the starting pixel of the same color as the
 * starting pixel, plus any pixels connected 4-directionally to those pixels
 * (also with the same color as the starting pixel), and so on.  Replace the
 * color of all of the aforementioned pixels with the newColor.
 *
 * At the end, return the modified image.
 *
 * Example 1:
 *
 * Input:
 * image = [[1,1,1],[1,1,0],[1,0,1]]
 * sr = 1, sc = 1, newColor = 2
 * Output: [[2,2,2],[2,2,0],[2,0,1]]
 * Explanation:
 * From the center of the image (with position (sr, sc) = (1, 1)), all pixels
 * connected
 * by a path of the same color as the starting pixel are colored with the new
 * color.
 * Note the bottom corner is not colored 2, because it is not 4-directionally
 * connected
 * to the starting pixel.
 *
 *
 *
 * Note:
 * The length of image and image[0] will be in the range [1, 50].
 * The given starting pixel will satisfy 0  and 0 .
 * The value of each color in image[i][j] and newColor will be an integer in
 * [0, 65535].
 *
 */

// @lc code=start
var dx = []int{-1, 1, 0.0}
var dy = []int{0, 0, -1, 1}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	m, n := len(image), len(image[0])

	coordinates := make([][]int, 1, m*n)
	coordinates[0] = []int{sr, sc}

	for len(coordinates) > 0 {
		c := coordinates[0]
		coordinates = coordinates[1:]
		image[c[0]][c[1]] = newColor

		for i := 0; i < 4; i++ {
			x := c[0] + dx[i]
			y := c[1] + dy[i]
			if 0 <= x && x < m && 0 <= y && y < n && image[x][y] == oldColor {
				coordinates = append(coordinates, []int{x, y})
			}
		}
	}

	return image

}

// @lc code=end

