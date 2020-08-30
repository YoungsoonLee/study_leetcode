/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 *
 * https://leetcode.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (48.11%)
 * Likes:    3253
 * Dislikes: 201
 * Total Accepted:    353.6K
 * Total Submissions: 730.2K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 
 * Given a binary tree, you need to compute the length of the diameter of the
 * tree. The diameter of a binary tree is the length of the longest path
 * between any two nodes in a tree. This path may or may not pass through the
 * root.
 * 
 * 
 * 
 * Example:
 * Given a binary tree 
 * 
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \     
 * ⁠     4   5    
 * 
 * 
 * 
 * Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
 * 
 * 
 * Note:
 * The length of path between two nodes is represented by the number of edges
 * between them.
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
func diameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)
	return res
}

func helper(root *TreeNode) (length, diameter int) {
	if root == nil {
		return
	}

	// DFS !!!
	leftLen, leftDia := helper(root.Left)
	rightLen, rightDia := helper(root.Right)

	length = max(leftLen, rightLen) + 1
	diameter = max(leftLen+rightLen, max(leftDia, rightDia))

	fmt.Println(root.Val, length, diameter, leftLen, rightLen, leftDia, rightDia)

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

