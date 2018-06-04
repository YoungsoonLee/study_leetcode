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

func minDiffInBST(root *TreeNode) int {
	res := 1<<63 - 1
	pre := 1 >> 63
	null := pre

	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root.Left != nil {
			helper(root.Left)
		}

		if pre != null {
			res = min(res, root.Val-pre)
		}

		pre = root.Val

		if root.Right != nil {
			helper(root.Right)
		}
	}

	helper(root)

	fmt.Println(res)
	return res
}

func minDiffInBST_my(root *TreeNode) int {
	res := 1<<63 - 1
	fmt.Println("res: ", res)
	dfs(root, &res)
	fmt.Println("result res: ", res)
	return res
}

func dfs(root *TreeNode, res *int) {
	pre := 1 >> 63
	null := pre

	if root.Left != nil {
		dfs(root.Left, res)
	}

	if pre != null {
		*res = min(*res, root.Val-pre)
	}

	pre = root.Val

	if root.Right != nil {
		dfs(root.Right, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {}
