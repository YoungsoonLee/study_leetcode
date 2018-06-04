package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
type Res struct {
	Val []int
}
*/

/*
func bfs_levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var s [][]int
	bfs(&s, 0, root)
	fmt.Println("s: ", s)

	// 반을 기준으로 뒤집음.
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	fmt.Println("s: ", s)
	return s
}

func bfs(s *[][]int, level int, root *TreeNode) {
	if root == nil {
		return
	}

	if len(*s) == level {
		*s = append(*s, []int{})
	}

	(*s)[level] = append((*s)[level], root.Val)

	for _, v := range []*TreeNode{root.Left, root.Right} {
		bfs(s, level+1, v)
	}
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	fmt.Println("dfs res: ", res)
	return res
}
*/

// normallu DFS + stack, BFS + queue
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if level >= len(res) {
			res = append([][]int{[]int{}}, res...)
		}
		//fmt.Println(res)
		n := len(res)
		//fmt.Println(n)
		res[n-level-1] = append(res[n-level-1], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	fmt.Println("dfs bottom: ", res)
	return res
}

func main() {

	t2Left := &TreeNode{Val: 15, Left: nil, Right: nil}
	t2Right := &TreeNode{Val: 7, Left: nil, Right: nil}

	t1Left := &TreeNode{Val: 9, Left: nil, Right: nil}
	t1Right := &TreeNode{Val: 20, Left: t2Left, Right: t2Right}

	t1Root := &TreeNode{Val: 3, Left: t1Left, Right: t1Right}

	//bfs_levelOrderBottom(t1Root)

	//levelOrder(t1Root)

	levelOrderBottom(t1Root)
}
