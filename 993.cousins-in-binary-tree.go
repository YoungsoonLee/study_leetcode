/*
 * @lc app=leetcode id=993 lang=golang
 *
 * [993] Cousins in Binary Tree
 *
 * https://leetcode.com/problems/cousins-in-binary-tree/description/
 *
 * algorithms
 * Easy (51.88%)
 * Likes:    364
 * Dislikes: 24
 * Total Accepted:    35.6K
 * Total Submissions: 68.4K
 * Testcase Example:  '[1,2,3,4]\n4\n3'
 *
 * In a binary tree, the root node is at depth 0, and children of each depth k
 * node are at depth k+1.
 *
 * Two nodes of a binary tree are cousins if they have the same depth, but have
 * different parents.
 *
 * We are given the root of a binary tree with unique values, and the values x
 * and y of two different nodes in the tree.
 *
 * Return true if and only if the nodes corresponding to the values x and y are
 * cousins.
 *
 *
 *
 * Example 1:
 *
 *
 *
 * Input: root = [1,2,3,4], x = 4, y = 3
 * Output: false
 *
 *
 *
 * Example 2:
 *
 *
 *
 * Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: root = [1,2,3,null,4], x = 2, y = 3
 * Output: false
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of nodes in the tree will be between 2 and 100.
 * Each node has a unique integer value from 1 to 100.
 *
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {

	root = &TreeNode{Left: root}
	px, dx := dfs(root, x)
	py, dy := dfs(root, y)

	return px != py && dx == dy
}

func dfs(root *TreeNode, x int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	if (root.Left != nil && root.Left.Val == x) || (root.Right != nil && root.Right.Val == x) {
		return root, 1
	}

	if parent, depth := dfs(root.Left, x); depth > 0 {
		return parent, depth + 1
	}

	if parent, depth := dfs(root.Right, x); depth > 0 {
		return parent, depth + 1
	}

	return nil, 0
}

// @lc code=end

