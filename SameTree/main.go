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

// brute-force

// recursive
// DFS = stack , BFS = queue
// this can be solved with one queue
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false
}

func main() {
	t1Left := &TreeNode{Val: 2, Left: nil, Right: nil}
	t1Right := &TreeNode{Val: 1, Left: nil, Right: nil}
	t1Root := &TreeNode{Val: 1, Left: t1Left, Right: t1Right}

	t2Left := &TreeNode{Val: 1, Left: nil, Right: nil}
	t2Right := &TreeNode{Val: 2, Left: nil, Right: nil}
	t2Root := &TreeNode{Val: 1, Left: t2Left, Right: t2Right}

	fmt.Println(isSameTree(t1Root, t2Root))

}
