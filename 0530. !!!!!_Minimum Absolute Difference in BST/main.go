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

type state struct {
	minDiff, previous int
}

func getMinimumDifference(root *TreeNode) int {
	st := state{1024, 1024}
	search(root, &st)
	return st.minDiff
}

// 주의 사항 : BST의 재귀 순회 방법
// 아주 좋아요.
func search(root *TreeNode, st *state) {
	if root == nil {
		return
	}

	search(root.Left, st)

	newDiff := diff(st.previous, root.Val)
	if st.minDiff > newDiff {
		st.minDiff = newDiff
	}

	st.previous = root.Val

	search(root.Right, st)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func main() {}
