package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func bfs(s *[][]int, level int, root *TreeNode) {
	if root == nil {
		return
	}

	if len(*s) == level {
		*s = append(*s, []int{})
	}

	(*s)[level] = append((*s)[level], root.Val)

	for _, v := range []*TreeNode{root.Left, root.Right} { // !!!
		bfs(s, level+1, v)
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var s [][]int
	bfs(&s, 0, root)
	return s
}

// dfs
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)

	fmt.Println(res)
	fmt.Println(len(res))

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 새로운 레벨이 나타났습니다.
		if level == len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}

func main() {

}
