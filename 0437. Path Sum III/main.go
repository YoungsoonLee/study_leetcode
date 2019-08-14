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

func pathSum(root *TreeNode, sum int) int {
	c := 0
	s := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		s += root.Val
		if s == sum {
			c++
		}

		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)
	return c
}

/*
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	cnt := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum -= node.Val
		if sum == 0 {
			cnt++
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, sum)

	fmt.Println(cnt)
	return cnt + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
*/

func main() {

}
