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

func findTilt(root *TreeNode) int {
	var tilt int
	helper(root, &tilt)
	return tilt
}

func helper(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}

	l := helper(root.Left, tilt)
	r := helper(root.Right, tilt)

	fmt.Println(l, r)

	if l > r {
		*tilt += l - r
	} else {
		*tilt += r - l
	}

	return l + r + root.Val
}

func main() {

}
