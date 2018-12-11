package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}

	res := make([][]int, n)
	fmt.Println("res: ", res)
	for i := range res {
		res[i] = make([]int, n)
	}

	fmt.Println("res: ", res)

	max := n * n
	next := nextFunc(n) //return function

	for i := 1; i <= max; i++ {
		x, y := next()
		fmt.Println("x,y,i: ", x, y, i)
		res[x][y] = i
	}

	fmt.Println("res: ", res)
	return res
}

func nextFunc(n int) func() (int, int) {
	top, down := 0, n-1
	left, right := 0, n-1
	x, y := 0, -1  // why?
	dx, dy := 0, 1 // why?

	return func() (int, int) {
		x += dx
		y += dy
		switch {

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

		//fmt.Println("x,y,dx,dy,top,down, left, right: ", x, y, dx, dy, top, down, left, right)
		return x, y
	}
}

func main() {
	generateMatrix(3)
}
