package main

import (
	"fmt"
)

// inout
// [ [ 1, 2, 3 ],[ 4, 5, 6 ],[ 7, 8, 9 ]]
// output
// 1,2,3,6,9,8,7,4,5
// index
// (0,0) (0,1) (0,2) (1,2) (2,2) (2,1) (2,0) (1,0) (1,1)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])

	next := nextFunc(m, n)
	//fmt.Println(next)

	res := make([]int, m*n)
	for i := range res {
		x, y := next()
		fmt.Println(x, y)
		res[i] = matrix[x][y]
	}

	fmt.Println(res)
	return res
}

func nextFunc(m, n int) func() (int, int) {
	top, down := 0, m-1
	left, right := 0, n-1
	// top 0 , down 2 , left 0 , right 2

	x, y := 0, -1
	dx, dy := 0, 1
	return func() (int, int) {
		fmt.Println(x, dx, y, dy)
		//fmt.Println(top, down, left, right)
		x += dx
		y += dy
		switch { // 如果撞墙了，需要修改 dx, dy 和相应的边界值
		case y+dy > right:
			top++
			dx, dy = 1, 0
		case x+dx > down:
			right--
			dx, dy = 0, -1
		case y+dy < left:
			down--
			dx, dy = -1, 0
		case x+dx < top:
			left++
			dx, dy = 0, 1
		}
		return x, y
	}

}

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	spiralOrder(a)
}
