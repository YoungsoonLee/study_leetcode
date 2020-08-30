/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 *
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (52.10%)
 * Likes:    892
 * Dislikes: 68
 * Total Accepted:    91.9K
 * Total Submissions: 170.8K
 * Testcase Example:  '[1,null,3,2]'
 *
 * Given a binary search tree with non-negative values, find the minimum
 * absolute difference between values of any two nodes.
 * 
 * Example:
 * 
 * 
 * Input:
 * 
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 * 
 * Output:
 * 1
 * 
 * Explanation:
 * The minimum absolute difference is 1, which is the difference between 2 and
 * 1 (or between 2 and 3).
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * There are at least two nodes in this BST.
 * This question is the same as 783:
 * https://leetcode.com/problems/minimum-distance-between-bst-nodes/
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
type state struct {
	minDiff, previous int
}

func getMinimumDifference(root *TreeNode) int {
	st := state{1024, 1024}
	search(root, &st)
	return st.minDiff
}

// BFS
func search(root *TreeNode, st *state) {
	if root == nil {
		return
	}

	search(root.Left, st)

	newDiff := diff(st.previous, root.Val)
	if st.minDiff > newDiff {
		st.minDiff = newDiff
	}

	search(root.Right, st)
}

func diff(a, b int) int {
	if a > b {
		return a -b
	}
	return b-a
}

// @lc code=end

