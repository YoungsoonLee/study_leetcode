package main

import "fmt"

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

/*
func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			res += 4
			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
					res--
				}
			}
		}
	}

	return res
}
*/

type data struct {
	n int
	s string
	//checks func() int
	b    bool
	doit func() float32
	m    map[string]string
}

func (d *data) checks() int {
	//
	return 0
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println(v1 == v2)
}

func sft(s []int) []int {
	s[0] = len(s)
	return s
}
