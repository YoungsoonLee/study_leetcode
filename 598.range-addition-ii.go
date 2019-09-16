/*
 * @lc app=leetcode id=598 lang=golang
 *
 * [598] Range Addition II
 */
func maxCount(m int, n int, ops [][]int) int {
    for _, v := range ops {
		m = min(m, v[0])
		n = min(n, v[1])
	}

	return m*n
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

