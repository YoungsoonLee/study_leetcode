package main

import "fmt"

func combine(n int, k int) [][]int {
	c := make([]int, k)
	res := [][]int{}

	fmt.Println("c: ", c)
	fmt.Println("res: ", res)

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == k {
			temp := make([]int, k)
			copy(temp, c)
			res = append(res, temp)
			return
		}

		for i := begin; i <= n+1-k+idx; i++ {
			c[idx] = i
			dfs(idx+1, i+1)
		}
	}

	dfs(0, 1)

	fmt.Println(res)
	return res
}

func main() {
	n := 6
	k := 3
	combine(n, k)
}
