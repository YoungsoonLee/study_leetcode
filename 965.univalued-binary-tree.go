/*
 * @lc app=leetcode id=965 lang=golang
 *
 * [965] Univalued Binary Tree
 *
 * https://leetcode.com/problems/univalued-binary-tree/description/
 *
 * algorithms
 * Easy (66.77%)
 * Likes:    277
 * Dislikes: 37
 * Total Accepted:    58.6K
 * Total Submissions: 87.5K
 * Testcase Example:  '[1,1,1,1,1,null,1]'
 *
 * A binary tree is univalued if every node in the tree has the same value.
 *
 * Return true if and only if the given tree is univalued.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,1,1,1,1,null,1]
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [2,2,2,5,2]
 * Output: false
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of nodes in the given tree will be in the range [1, 100].
 * Each node's value will be an integer in the range [0, 99].
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (root.Left != nil && root.Val != root.Left.Val) || (root.Right != nil && root.Val != root.Right.Val) {
		return false
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

// @lc code=end

