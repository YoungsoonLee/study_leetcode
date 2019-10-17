/*
 * @lc app=leetcode id=783 lang=golang
 *
 * [783] Minimum Distance Between BST Nodes
 *
 * https://leetcode.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (50.95%)
 * Likes:    428
 * Dislikes: 124
 * Total Accepted:    42K
 * Total Submissions: 82.4K
 * Testcase Example:  '[4,2,6,1,3,null,null]'
 *
 * Given a Binary Search Tree (BST) with the root node root, return the minimum
 * difference between the values of any two different nodes in the tree.
 *
 * Example :
 *
 *
 * Input: root = [4,2,6,1,3,null,null]
 * Output: 1
 * Explanation:
 * Note that root is a TreeNode object, not an array.
 *
 * The given tree [4,2,6,1,3,null,null] is represented by the following
 * diagram:
 *
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \
 * ⁠   1   3
 *
 * while the minimum difference in this tree is 1, it occurs between node 1 and
 * node 2, also between node 3 and node 2.
 *
 *
 * Note:
 *
 *
 * The size of the BST will be between 2 and 100.
 * The BST is always valid, each node's value is an integer, and each node's
 * value is different.
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
func minDiffInBST(root *TreeNode) int {
	res := 1<<63 - 1
	pre := 1 >> 63
	null := pre

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}

		// 얻어야 하는 조건.
		if pre != null {
			res = min(res, root.Val-pre)
		}

		//
		pre = root.Val

		//
		if root.Right != nil {
			dfs(root.Right)
		}

	}

	dfs(root)

	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

