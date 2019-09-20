/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
 *
 * https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (53.21%)
 * Likes:    990
 * Dislikes: 116
 * Total Accepted:    102.1K
 * Total Submissions: 191.5K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * Given a Binary Search Tree and a target number, return true if there exist
 * two elements in the BST such that their sum is equal to the given target.
 * 
 * Example 1:
 * 
 * 
 * Input: 
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 * 
 * Target = 9
 * 
 * Output: True
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 * 
 * Target = 28
 * 
 * Output: False
 * 
 * 
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	return helper(root, root, k)
}

func helper(root, searchRoot *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	return (root.Val * 2 != k &&findNode(searchRoot, k-root.Val)) ||
	helper(root.Left, searchRoot, k) ||
	helper(root.Right, searchRoot, k) 
}

func findNode(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	if root.Val < target {
		return findNode(root.Right, target)
	}

	return findNode(root.Left, target)
}
 
