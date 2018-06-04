package main

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

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSame(s, t) {
		return true
	}

	if s == nil {
		return false
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t) // !!!!!
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}

	// 이미 s는 nil이 아니므로...
	if t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}

func main() {}
