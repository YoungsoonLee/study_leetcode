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

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

// recursion
// time: O(n), space: O(n)
func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isMirror(t1.Right, t2.Left) && isMirror(t1.Left, t2.Right)
}

// Iterative, BFS using queue
func isSymmetric2(root *TreeNode) bool {
	q := make([]*TreeNode, 0)
	//fmt.Println(len(q))

	q = append(q, root)
	q = append(q, root)

	for len(q) >= 0 {
		var t1 *TreeNode
		//t1, q = q[0], q[1:] // pop
		t1, q = q[len(q)-1], q[:len(q)-1]

		var t2 *TreeNode
		//t2, q = q[0], q[1:] // pop
		t2, q = q[len(q)-1], q[:len(q)-1]

		fmt.Println(q, t1, t2)

		if t1 == nil && t2 == nil {
			continue
		}

		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}

		q = append(q, t1.Left)
		q = append(q, t2.Right)
		q = append(q, t1.Right)
		q = append(q, t2.Left)

	}

	return false
}

func main() {
	t3Left := &TreeNode{Val: 4, Left: nil, Right: nil}
	t3Right := &TreeNode{Val: 3, Left: nil, Right: nil}

	t2Left := &TreeNode{Val: 3, Left: nil, Right: nil}
	t2Right := &TreeNode{Val: 4, Left: nil, Right: nil}

	t1Left := &TreeNode{Val: 2, Left: t2Left, Right: t3Left}
	t1Right := &TreeNode{Val: 2, Left: t3Right, Right: t2Right}

	t1Root := &TreeNode{Val: 1, Left: t1Left, Right: t1Right}

	fmt.Println(isSymmetric2(t1Root))

}
