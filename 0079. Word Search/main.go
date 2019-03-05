package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	if n == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}

	var dfs func(int, int, int) bool
	dfs = func(r, c, index int) bool {
		if len(word) == index {
			return true
		}

		if r < 0 || m <= r || c < 0 || n <= c || board[r][c] != word[index] {
			return false
		}

		// !!!
		temp := board[r][c]
		// !!!
		board[r][c] = 0

		if dfs(r-1, c, index+1) || dfs(r+1, c, index+1) || dfs(r, c-1, index+1) || dfs(r, c+1, index+1) {
			return true
		}
		// !!!
		board[r][c] = temp

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// me ...
/*
func exist(board [][]byte, word string) bool {

	m := make(map[byte]int, 0)

	for _, r := range board {
		for _, c := range r {
			if _, ok := m[c]; ok {
				m[c]++
			} else {
				m[c] = 1
			}
		}
	}
	fmt.Println("map:  ", m)

	for _, b := range word {
		if _, ok := m[byte(b)]; ok {
			if m[byte(b)] == 0 {
				return false
			} else {
				m[byte(b)]--
			}
		} else {
			return false
		}
	}

	return true
}
*/

func main() {
	/*
		[
		  ['A','B','C','E'],
		  ['S','F','C','S'],
		  ['A','D','E','E']
		]
	*/
	/*
		b := [][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'C', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		}
	*/
	b := [][]byte{
		[]byte{'a', 'b'},
		[]byte{'c', 'd'},
	}
	w := "abcd"
	fmt.Println(exist(b, w))
}
