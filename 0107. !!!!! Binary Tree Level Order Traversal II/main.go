package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	var dfs func(*TreeNode, int, *[][]int)
	dfs = func(root *TreeNode, level int, res *[][]int) {
		if root == nil {
			return
		}

		if level >= len((*res)) {
			(*res) = append([][]int{[]int{}}, (*res)...) // !!! 앞에다가 결과 넣을 배열을 만드네~!!!!!!
		}

		n := len((*res))
		(*res)[n-level-1] = append((*res)[n-level-1], root.Val) // !!!

		dfs(root.Left, level+1, res)
		dfs(root.Right, level+1, res)
	}

	dfs(root, 0, &res)

	return res
}

func main() {
	res := [][]int{}
	res = append([][]int{[]int{}}, res...) // !!!
	fmt.Println(res, len(res))

	n := len(res)
	level := 0
	res[n-level-1] = append(res[n-level-1], 3) // !!!
	fmt.Println(res, len(res))

	// level 1
	res = append([][]int{[]int{}}, res...) // !!!
	fmt.Println(res, len(res))
	n = len(res)
	level = 1
	res[n-level-1] = append(res[n-level-1], 9) // !!!
	fmt.Println(res, len(res))
}
