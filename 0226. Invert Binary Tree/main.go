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

// my solution - failed
/*
func invertTree(root *TreeNode) *TreeNode {
	// BFS ?
	subleft := root.Left
	subright := root.Right

	for subleft != nil && subright != nil {
		subleft, subright = changeTress(root.Left, root.Right)
	}

	return root
}

func changeTress(left *TreeNode, right *TreeNode) (subleft, subright *TreeNode) {
	left.Val, right.Val = right.Val, left.Val
	return left, right
}
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func main() {
	fmt.Println("")
}
