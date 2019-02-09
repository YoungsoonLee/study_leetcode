package main

import "fmt"

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := 1

	/**
	 * 从上往下，从左往右 扫描矩阵
	 * 利用 mat[i][0] = 0 表示，第 i 行中含有 0
	 * 利用 mat[0][j] = 0 表示，第 j 列中含有 0
	 * 特别地，
	 * mat[0][0] = 0 仅表示，第 0 行中含有 0
	 * Col0 = 0 表示，第 0 列中含有 0
	 * NOTICE: 循环的顺序很重要
	 * 需要保证 mat[i][0] 和 mat[0][j] 被标记后，不再做为别的标记的依据
	 * 要不然的话，标记有可能会被污染
	 */

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	fmt.Println(matrix)
	//fmt.Println(col0)

	//col0 = 0

	/**
	 * 从下往上，从右往左 扫描矩阵
	 * 并根据前面的标记修改 mat[i][j] 的值
	 * NOTICE: 循环的顺序很重要
	 * 需要保证 mat[i][0] 和 mat[0][j] 被修改后，不再做为别的修改的标记
	 * 要不然的话，标记有可能会被污染
	 */
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 == 0 {
			matrix[i][0] = 0
		}
	}
	fmt.Println(matrix)

	/*
		rl := len(matrix)
		cl := len(matrix[0])

		r := make([]int, 0, (rl+1)/2)
		c := make([]int, 0, (cl+1)/2)

		// 1. full scan
		for i := 0; i < rl; i++ {
			for j := 0; j < cl; j++ {
				// 2. get index of 0
				if matrix[i][j] == 0 {
					r = append(r, i)
					c = append(c, j)
				}
			}
		}

		fmt.Println("r: ", r)
		fmt.Println("c: ", c)

		// 3. change matrix with #2
		// row
		for _, ri := range r {
			for i := 0; i < cl; i++ {
				matrix[ri][i] = 0
			}
		}

		// column
		for _, ci := range c {
			for i := 0; i < rl; i++ {
				matrix[i][ci] = 0
			}
		}

		fmt.Println(matrix)
	*/
}

func main() {
	m := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(m)

	m = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(m)
}
