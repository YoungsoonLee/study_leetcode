package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

	t2Left := &TreeNode{Val: 15, Left: nil, Right: nil}
	t2Right := &TreeNode{Val: 7, Left: nil, Right: nil}

	t1Left := &TreeNode{Val: 9, Left: nil, Right: nil}
	t1Right := &TreeNode{Val: 20, Left: t2Left, Right: t2Right}

	t1Root := &TreeNode{Val: 3, Left: t1Left, Right: t1Right}

	fmt.Println(minDepth(t1Root))
}
