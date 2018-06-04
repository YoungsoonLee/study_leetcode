package main

import (
	"fmt"
	"math"
)

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depthLeft, depthRight int
	depthLeft = maxDepth(root.Left) + 1
	depthRight = maxDepth(root.Right) + 1
	return int(math.Max(float64(depthLeft), float64(depthRight)))
}

func main() {

	t2Left := &TreeNode{Val: 15, Left: nil, Right: nil}
	t2Right := &TreeNode{Val: 7, Left: nil, Right: nil}

	t1Left := &TreeNode{Val: 9, Left: nil, Right: nil}
	t1Right := &TreeNode{Val: 20, Left: t2Left, Right: t2Right}

	t1Root := &TreeNode{Val: 3, Left: t1Left, Right: t1Right}

	fmt.Println(maxDepth(t1Root))
}
