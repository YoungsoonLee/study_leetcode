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

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	search(root, &sum)
	return root
}

func search(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	search(root.Right, sum)

	*sum += root.Val
	root.Val = *sum

	search(root.Left, sum)

}

func main() {

}
